package users

import "github.com/CodeHanHan/ferry-backend/utils/idutil"

const UsersTableName = "users_table"

type UsersTable struct {
	ID       string `gorm:"column:id;primary_key"`
	UserName string `gorm:"column:user_name"`
	Password string `gorm:"column:password"`
	Email    string `gorm:"column:email"`
}

func NewUsersTable(message string, reply string) *UsersTable {
	return &PingRecord{
		PingID:  idutil.NewHexId(),
		Message: message,
		Reply:   reply,
	}
}
