package user

import (
	"time"

	"github.com/CodeHanHan/ferry-backend/utils/idutil"
)

const UserTableName = "user"

type UserTable struct {
	ID         string     `gorm:"column:id;primary_key" json:"id"`
	Username   string     `gorm:"column:username;unique" json:"username"`
	Nickname   string     `gorm:"column:nickname" json:"nickname"`
	Password   string     `gorm:"column:password" json:"password"`
	Email      string     `gorm:"column:email" json:"email"`
	Role       string     `gorm:"column:role" json:"role"`
	CreateTime *time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL" json:"create_time"`
}

func NewUserTable(username string, password string, role string, email string) *UserTable {
	return &UserTable{
		ID:       idutil.GetId("user"),
		Username: username,
		Password: password,
		Role:     role,
		Email:    email,
	}
}
