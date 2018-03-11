package models

import (
	"time"

	"github.com/ezradiniz/meeting-server/database"
	"github.com/ezradiniz/meeting-server/serializers"
)

// PostModel model
type PostModel struct {
	ID        uint      `gorm:"primary_key"`
	Text      string    `gorm:"text"`
	UserID    uint      `gorm:"user_id"`
	CreatedAt time.Time `gorm:"created_at"`
	User      UserModel
}

// PostAutoMigrate ...
func PostAutoMigrate() {
	db := database.GetDB()
	db.AutoMigrate(&PostModel{})
}

// AddPost insert post into database
func AddPost(post interface{}) error {
	db := database.GetDB()
	return db.Save(post).Error
}

// FindAllPostsByUser find all posts from user
func FindAllPostsByUser(userID uint) ([]PostModel, error) {
	db := database.GetDB()

	user := UserModel{}
	posts := []PostModel{}

	db.First(&user, userID)
	return posts, db.Model(&user).Related(&posts, "Posts").Error
}

// FindAllPosts find all posts
func FindAllPosts() ([]serializers.FeedResponse, error) {
	db := database.GetDB()

	res := []serializers.FeedResponse{}

	// TODO: Refactor later
	err := db.Table("post_models").Select("post_models.text as text, post_models.id as id, post_models.created_at as created_at, user_models.name as name, user_models.username as username").Joins("inner join user_models ON post_models.user_id = user_models.id").Order("created_at desc").Scan(&res).Error
	return res, err
}
