package idutil

import (
	"strings"

	"github.com/google/uuid"
)

func NewHexId() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
}
