package handler

import (
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/gin-gonic/gin"
	"image/png"
	"net/http"
	"os"
)
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, token, Token")
		c.Header("Access-Control-Allow-Methods", "*")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
func Show(c *gin.Context){
	url := c.Query("url")
	if url == ""{
		url = "123"
	}
	fileName := url + ".png"
	qrCode, _ := qr.Encode(url, qr.M, qr.Auto)

	qrCode, _ = barcode.Scale(qrCode, 256, 256)
	file, _ := os.Create(fileName)
	defer file.Close()
	_ = png.Encode(file, qrCode)
	c.Header("Content-Type", "image/jpg")
	c.File(fileName)
}