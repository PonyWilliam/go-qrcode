package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ponywilliam/go-qrcode/handler"
)

func showImage(){

}
func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/",handler.Cors(),handler.Show)
	_ = router.Run(":80")
}