package main

import(
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	db := connectDB()
	defer db.Close()

	db.Set("gorm:table_options", "ENGINE = InnoDB CHARSET=utf8mb4",).AutoMigrate(&channel{})

	go trendingTicker()

	router := gin.Default()
	router.StaticFS("/static", http.Dir("static"))
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/home/:date", getTrending)
	router.GET("/home", transition)

	router.Run(":8000")
}