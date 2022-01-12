package user

import (
	"context"
	"encoding/base64"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/CodeHanHan/ferry-backend/pkg/app"
	"github.com/CodeHanHan/ferry-backend/pkg/constants"
	formUser "github.com/CodeHanHan/ferry-backend/pkg/form/user"
	"github.com/CodeHanHan/ferry-backend/pkg/logger"
	"github.com/CodeHanHan/ferry-backend/pkg/pi"
	"github.com/CodeHanHan/ferry-backend/pkg/sender"
	"github.com/CodeHanHan/ferry-backend/utils/fileutil"
	"github.com/CodeHanHan/ferry-backend/utils/idutil"
	"github.com/CodeHanHan/ferry-backend/utils/stringutil"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// @Summary 上传头像
// @Description 上传头像图片，只支持png, jpeg, jpg
// @Tags user
// @ID user-avatar
// @Param avatar body formUser.UploadAvatarRequest true "用户头像"
// @Success 200 {object} formUser.UploadAvatarResponse
// @Failure 500 {object} app.ErrResponse
// @Failure 400 {object} app.ErrResponse
// @Accept application/json
// @Produce  json
// @Router /user/upload [post]
// @Security BearerAuth
func UploadAvatar(c *gin.Context) {
	var req formUser.UploadAvatarRequest
	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		logger.ErrorParams(c, err)
		app.ErrorParams(c, err)
		return
	}

	b64Str := req.Base64
	filename := req.FileName

	b64, err := base64.StdEncoding.DecodeString(b64Str)
	if err != nil {
		logger.Error(c, err.Error())
		app.InternalServerError(c)
		return
	}

	filetype := fileutil.GetFileTypeByFileBytes(b64)
	if filetype != fileutil.PNG && filetype != fileutil.JPEG && filetype != fileutil.JPG {
		app.Errorf(c, app.Err_Invalid_Argument, "invalid file type: %s, only PNG, JPEG and JPG are supported", filetype)
		return
	}

	sender, _, err := sender.GetSender(c)
	if err != nil {
		app.Errorf(c, app.Err_Unauthenticated, "Unable to get sender")
		return
	}

	filepath, err := saveAvatar(c, sender, filename, b64)
	if err != nil {
		logger.Error(c, err.Error())
		app.InternalServerError(c)
		return
	}

	resp := formUser.UploadAvatarResponse{
		FilePath: filepath,
		AbsPath:  pi.Global.Cfg.Application.StaticPrefix + filepath,
	}

	app.OK(c, resp)
}

// ===================================
func saveAvatar(ctx context.Context, sender, filename string, data []byte) (filepath string, err error) {
	hashDir := fileutil.MakeHashDir(sender)
	split := strings.SplitN(filename, ".", 2)

	var storeName string
	if len(split) > 1 {
		storeName = idutil.GetId(sender) + "." + split[len(split)-1]
	} else {
		storeName = idutil.GetId(sender)
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	prefix := stringutil.Join(home, "/", constants.PROJECT_DATABASE, constants.AVATAR_PATH)

	storeDir := path.Join(prefix, hashDir)
	if err := fileutil.MakeDirAll(storeDir, os.ModePerm); err != nil {
		return "", err
	}

	if err := ioutil.WriteFile(stringutil.Join(storeDir, "/", storeName), data, 0666); err != nil {
		return "", err
	}

	return stringutil.Join("avatar/", hashDir, "/", storeName), nil
}
