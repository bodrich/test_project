package web

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"project/logger"
)

// для логгирования ошибок и отправки их пользователю по HTTP
func checkError(c *gin.Context, err error) error {
	if err != nil {
		err := errors.Errorf(err.Error())
		logger.WriteLog("%+v", err)
		c.JSON(500, gin.H{"error": err.Error()})
	}
	return err
}
