package models

import (
	"github.com/labstack/echo/v5"
	"go.uber.org/zap"
)

type AppContext struct {
	Logger *zap.Logger
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

	return &AppContext{
		Logger: logger,
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
