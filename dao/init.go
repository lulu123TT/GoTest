package dao

import (
	"blog/model"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Manager interface {
	Register(user *model.User)
	Login(username string) model.User
	// AddPost
	// 博客操作
	//AddPost(post *model.Post)
	//GetAllPost() []model.Post
	//GetPost(pid int) model.Post
	CountEmail(email string) int64
}

var Mgr Manager

type manager struct {
	db *gorm.DB
}

//init执行一次
func init() {
	dns := "root:wsxdx147369@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to init db", err)
	}
	Mgr = &manager{db: db}
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Post{})
}

var RDB = InitRedisDB()

func InitRedisDB() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
