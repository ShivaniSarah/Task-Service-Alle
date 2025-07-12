// config/config.go
package config

import "github.com/gin-gonic/gin"

func Init() {
	gin.SetMode(gin.ReleaseMode)
}
