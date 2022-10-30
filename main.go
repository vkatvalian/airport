package main

import (
    "fmt"
    "log"
    "os"
    "net/http"
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
    "github.com/joho/godotenv"
    "github.com/gin-gonic/gin"
)

var db *gorm.DB

func feed(c *gin.Context){
    username := c.PostForm("username")
    email := c.PostForm("email")
    password := c.PostForm("password")
    fmt.Println(username, email, password)

    c.JSON(http.StatusOK, gin.H{
	"status":  "success",
        "message": username,
    })
}

func signup(c *gin.Context){
    c.HTML(http.StatusOK, "signup.tmpl", gin.H{})
}

func index(c *gin.Context){
    c.HTML(http.StatusOK, "signup.tmpl", gin.H{})
}

func main(){
    var err error

    err = godotenv.Load()
    if err != nil {
      log.Fatal("Error loading .env file")
    }

    dsn := os.Getenv("DSN")
    db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
	fmt.Println(err)
        panic("failed to connect database")
    }

    router := gin.Default()
    router.LoadHTMLGlob("tmpl/*")
    router.GET("/signup", signup)
    router.POST("/feed", feed)
    router.POST("/", index)
    router.Run()
}
