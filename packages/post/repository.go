package post

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Post, error)
	FindByID(id uint) (Post, error)
	FindBySlug(slug string) (Post, error)
	Create(post *Post) error
	Update(post *Post) error
	Delete(id uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) FindAll() ([]Post, error) {
	var posts []Post
	err := r.db.Find(&posts).Error
	return posts, err
}

func (r *repository) FindByID(id uint) (Post, error) {
	var post Post
	err := r.db.First(&post, id).Error
	return post, err
}

func (r *repository) FindBySlug(slug string) (Post, error) {
	var post Post
	err := r.db.Where("slug = ?", slug).First(&post).Error
	return post, err
}

func (r *repository) Create(post *Post) error {
	return r.db.Create(post).Error
}

func (r *repository) Update(post *Post) error {
	return r.db.Save(post).Error
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&Post{}, id).Error
}
