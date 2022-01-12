package user

type UploadAvatarRequest struct {
	Base64   string `json:"img_b64" form:"img_b64" binding:"required,base64"`
	FileName string `json:"file_name" form:"file_name" binding:"required"`
}

type UploadAvatarResponse struct {
	FilePath string `json:"file_path"`
	AbsPath  string `json:"abs_path"`
}
