// ไฟล์ main.go
package main

import (
	"fmt"
	"myproject/calc"
	"myproject/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	result := calc.Add(5, 3)
	fmt.Println("5 + 3 =", result)

	thaiNumber := utils.NumberToThai(result)
	fmt.Println("ผลลัพธ์เป็นภาษาไทย:", thaiNumber)

	// สร้าง Gin engine
	r := gin.Default()

	// กำหนด route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// กำหนด route ที่รับพารามิเตอร์
	r.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// รัน server ที่ port 8080
	r.Run(":8080")
}
