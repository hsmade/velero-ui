package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hsmade/velero-ui/api"
)

func main() {
	r := gin.Default()
	err := api.New(r.Group("/api/v1"))
	if err != nil {
		panic(err)
	}
	r.Static("/assets", "./web/dist/assets")
	r.StaticFile("/", "./web/dist/index.html")
	r.Run(":8080")
}
