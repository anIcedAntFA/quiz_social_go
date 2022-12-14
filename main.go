package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"social_quiz/components/appctx"
	"social_quiz/middleware"
)

func main() {
	dsn := os.Getenv("MYSQL_CONNECTION")
	secretKey := os.Getenv("SYSTEM_SECRET")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Cannot connect to MySQL:", err)
	}

	db = db.Debug()

	appContext := appctx.NewAppContext(db, secretKey)

	router := gin.Default()

	// - No origin allowed by default
	// - GET,POST, PUT, HEAD methods
	// - Credentials share disabled
	// - Preflight requests cached for 12 hours
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://127.0.0.1:5173"}
	// config.AllowOrigins = []string{"http://google.com", "http://facebook.com"}
	// config.AllowAllOrigins = true

	router.Use(cors.New(config))

	router.Use(middleware.Recover(appContext))
	v1 := router.Group("/v1")

	setupMainRoute(appContext, v1)

	setupAdminRoute(appContext, v1)

	router.GET("/qwerty", aQwerty)

	router.Run()
}

type pagingParams struct {
	PageIndex int    `json:"page_index" form:"page_index"`
	PageSize  int    `json:"page_size" form:"page_size"`
	Sort      string `json:"sort" form:"sort"`
	Search    string `json:"search" form:"search"`
}

func aQwerty(c *gin.Context) {
	var param pagingParams

	if c.ShouldBindQuery(&param) == nil {
		log.Println(param)
	}

	c.JSON(http.StatusOK, param)
}
