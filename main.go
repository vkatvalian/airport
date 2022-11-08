package main

import (
    "fmt"
    "log"
    "os"
    "time"
    "net/http"
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
    "github.com/joho/godotenv"
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
)

var db *gorm.DB

type Users struct {
    ID        uint      `gorm:"primaryKey"`
    Username  string	`gorm:"not null"`
    Email     string	`gorm:"not null"`
    Password  string	`gorm: "not null"`
    CreatedAt time.Time
    UpdatedAt time.Time
}

func feed(c *gin.Context){
    username := c.PostForm("username")
    email := c.PostForm("email")
    password := []byte(c.PostForm("password"))

    hashed_password, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
    if err != nil {
	panic(err)
    }

    err = bcrypt.CompareHashAndPassword(hashed_password, password)
    if err != nil {
        log.Fatal(err)
    }

    db.AutoMigrate(&Users{})
    user := &Users{
        Username: username,
	Email: email,
	Password: string(hashed_password),
    }

    db.Create(&user)
    c.JSON(http.StatusOK, gin.H{
	"status":  "success",
        "message": username,
    })
}

func signup(c *gin.Context){
    c.HTML(http.StatusOK, "signup.tmpl", gin.H{})
}

func signin_post(c *gin.Context){
     _name := c.PostForm("username")
    password := c.PostForm("password")

    var user Users
    db.Table("users").Where("username = ?", _name).Select("username, email, password").Find(&user)
    ok := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    if ok != nil {
        log.Println(ok)
        c.JSON(http.StatusUnauthorized, gin.H{
      	  "status": "fail",
        })
    } else {
	c.JSON(http.StatusOK, gin.H{
          "status": "success",
	  "username": user.Username,
	  "email": user.Email,
        })
    }
}

func signin(c *gin.Context){
    c.HTML(http.StatusOK, "signin.tmpl", gin.H{})
}

func index(c *gin.Context){
    c.HTML(http.StatusOK, "index.tmpl", gin.H{})
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

    db.AutoMigrate(&Users{})

    router := gin.Default()
    router.LoadHTMLGlob("tmpl/*")
    router.GET("/signup", signup)
    router.GET("/signin", signin)
    router.POST("/signin", signin_post)
    router.POST("/feed", feed)
    router.GET("/", index)
    router.Run()
}
