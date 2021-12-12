package users

import (
	"github.com/CodeHanHan/ferry-backend/utils/idutil"
)

const UsersTableName = "users_table"

type UsersTable struct {
	ID       string `gorm:"column:id;primary_key"`
	UserName string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	Role     string `gorm:"column:role"`
	Email    string `gorm:"column:email"`
}

func NewUsersTable(username string, password string, role string, email string) *UsersTable {
	return &UsersTable{
		ID:       idutil.NewHexId(),
		UserName: username,
		Password: password,
		Role:     role,
		Email:    email,
	}
}
