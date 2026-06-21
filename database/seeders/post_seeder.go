package seeders

import (
	"log"

	"github.com/alirfanyasin/golang-starterkit/packages/post"
	"gorm.io/gorm"
)

// SeedPosts seeds default post data if empty
func SeedPosts(db *gorm.DB) error {
	var count int64
	db.Model(&post.Post{}).Count(&count)

	if count == 0 {
		posts := []post.Post{
			{
				Title:   "Selamat Datang di Golang Starterkit",
				Slug:    "selamat-datang-di-golang-starterkit",
				Content: "Ini adalah artikel pertama yang dibuat secara otomatis oleh seeder database. Starterkit ini didesain agar mudah dikembangkan.",
				Status:  "published",
			},
			{
				Title:   "Tips Belajar REST API dengan Gin Gonic",
				Slug:    "tips-belajar-rest-api-dengan-gin-gonic",
				Content: "Belajar membuat REST API di Go sangat menyenangkan. Framework Gin memberikan performa tinggi dan routing yang mudah dipahami.",
				Status:  "published",
			},
			{
				Title:   "Draft Artikel Baru",
				Slug:    "draft-artikel-baru",
				Content: "Artikel ini masih berupa draf kasar dan tidak ditampilkan ke publik secara langsung.",
				Status:  "draft",
			},
		}

		for _, p := range posts {
			if err := db.Create(&p).Error; err != nil {
				return err
			}
		}
		log.Println("Post seeder: 3 posts created successfully.")
	} else {
		log.Println("Post seeder: skipped (table is not empty).")
	}

	return nil
}
