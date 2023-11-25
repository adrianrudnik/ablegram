package webservice

import (
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"os"
)

var Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()

func LoggerMiddleware() gin.HandlerFunc {
	return logger.SetLogger(logger.WithLogger(func(context *gin.Context, z zerolog.Logger) zerolog.Logger {
		return Logger
	}))
}
