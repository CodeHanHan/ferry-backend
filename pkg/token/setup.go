package token

import (
	"context"

	"github.com/CodeHanHan/ferry-backend/pkg/logger"
)

func SetUp(register func(maker Maker), secretKey string) error {
	tokenMaker, err := NewJWTMaker(secretKey)
	if err != nil {
		logger.Critical(context.Background(), "failed to create token maker: %v", err)
		return err
	}
	register(tokenMaker)
	return nil
}
