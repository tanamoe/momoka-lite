package models

import (
	"errors"
	"os"

	"github.com/labstack/echo/v5"
	"go.uber.org/zap"
)

type AppContext struct {
	Logger       *zap.Logger
	ImagorSecret string
}

type EchoContext struct {
	echo.Context
	AppContext *AppContext
}

func NewContext() (*AppContext, error) {
	logger, err := logger()
	if err != nil {
		return nil, err
	}

	imagorSecret, err := imagorSecret()
	if err != nil {
		return nil, err
	}

	return &AppContext{
		Logger:       logger,
		ImagorSecret: imagorSecret,
	}, nil
}

func logger() (*zap.Logger, error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}
	defer logger.Sync()
	return logger, nil
}

func (context *EchoContext) Logger() *zap.Logger {
	return context.AppContext.Logger
}

func imagorSecret() (string, error) {
	value, ok := os.LookupEnv("IMAGOR_SECRET")

	if !ok {
		return "", errors.New("IMAGOR_SECRET not set")
	}

	return value, nil
}

func (context *EchoContext) ImagorSecret() string {
	return context.AppContext.ImagorSecret
}
