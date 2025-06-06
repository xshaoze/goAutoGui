package appError

import (
	"errors"
)

type AppError struct {
	code    string
	message string
}

var errorMessages = map[string]map[string]string{
	"zh": {
		"unknownError":      "未知错误",
		"systemNotSupport":  "当前系统暂时不支持",
		"getCursorPosError": "获取鼠标当前位置失败",
	},
	"en": {
		"unknownError":      "Unknown error!",
		"systemNotSupport":  "The current system does not currently support it",
		"getCursorPosError": "Failed to retrieve the current position of the mouse",
	},
}

func Error(errorType, lang string) error {
	if message := errorMessages[lang][errorType]; message != "" {
		return errors.New(errorMessages[lang][errorType])
	}
	return errors.New(errorMessages[lang]["unknownError"])
}
