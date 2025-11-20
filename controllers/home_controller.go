package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "🏠 Trang chủ")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ℹ️ Giới thiệu về chúng tôi")
}

func HomeHandler(c *gin.Context) {
	c.String(http.StatusOK, "🏠 Trang chủ!")
}

func AboutHandler(c *gin.Context) {
	c.String(http.StatusOK, "ℹ️ Giới thiệu về chúng tôi!")
}

func Http() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)

	fmt.Println("🌐 Server đang chạy tại http://localhost:8084")
	err := http.ListenAndServe(":8084", nil)
	if err != nil {
		fmt.Println("❌ Lỗi khi chạy server:", err)
	}

	// router := gin.Default()
	// router.GET("/", HomeHandler)
	// router.GET("/about", AboutHandler)
	// router.Run(":8084")
}

func HomeIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello from Go + Gin",
	})
}
