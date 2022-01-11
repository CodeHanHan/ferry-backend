package idutil

import (
	"strings"

	"github.com/google/uuid"
)

func NewHexId() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}

func GetId(prefix string) string {
	list := strings.Split(uuid.New().String(), "-")
	return prefix + "-" + list[len(list)-1]
}
