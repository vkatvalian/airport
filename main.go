package main

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)

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
    router := gin.Default()
    router.LoadHTMLGlob("tmpl/*")
    router.GET("/signup", signup)
    router.POST("/feed", feed)
    router.POST("/", index)
    router.Run()
}
