package main

import (
	"github.com/gin-gonic/gin"

	"github.com/dgkg/project/handler"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/artists", handler.GetAllArtist)
	r.GET("/artists/:id", handler.GetArtist)
	r.Run()
}
