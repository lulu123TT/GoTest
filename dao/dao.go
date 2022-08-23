package dao

import (
	"blog/model"
)

func (mgr *manager) Register(user *model.User) {
	mgr.db.Create(user)
}

func (mgr *manager) Login(username string) model.User {
	var user model.User
	mgr.db.Where("username=?", username).First(&user)
	return user
}

func (mgr *manager) CountEmail(email string) int64 {
	var cnt int64
	mgr.db.Where("email=?", email).Count(&cnt)
	return cnt
}

////AddPost
////博客操作
//func (mgr *manager) AddPost(post *model.Post) {
//	mgr.db.Create(post)
//}
//
//func (mgr *manager) GetAllPost() []model.Post {
//	var posts = make([]model.Post, 10)
//	mgr.db.Find(&posts)
//	return posts
//}
//
//func (mgr *manager) GetPost(pid int) model.Post {
//	var post model.Post
//	mgr.db.First(&post, pid)
//	return post
//}
