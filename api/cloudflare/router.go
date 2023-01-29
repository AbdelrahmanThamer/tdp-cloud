package cloudflare

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/module/midware"
)

func Router(api *gin.RouterGroup) {

	rg := api.Group("/cloudflare")

	rg.Use(midware.AuthGuard())

	{
		rg.GET("/:id", getProxy)
		rg.POST("/:id", postProxy)
	}

}