package frontend

import (
	"embed"
	"github.com/gin-gonic/gin"
	"net/http"
)

//go:embed *

var Static embed.FS

func Register(gin *gin.Engine, webPath string) {
	gin.StaticFS("/timer/", http.FS(Static))
}
