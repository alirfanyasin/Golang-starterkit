package post

import (
	"regexp"
	"strings"
)

type CreatePostInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Status  string `json:"status"` // draft, published
}

type UpdatePostInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Status  string `json:"status"`
}

type Service interface {
	GetAll() ([]Post, error)
	GetByID(id uint) (Post, error)
	Create(input CreatePostInput) (Post, error)
	Update(id uint, input UpdatePostInput) (Post, error)
	Delete(id uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetAll() ([]Post, error) {
	return s.repo.FindAll()
}

func (s *service) GetByID(id uint) (Post, error) {
	return s.repo.FindByID(id)
}

func (s *service) Create(input CreatePostInput) (Post, error) {
	slug := generateSlug(input.Title)

	// Set default status if empty
	status := input.Status
	if status == "" {
		status = "draft"
	}

	post := Post{
		Title:   input.Title,
		Slug:    slug,
		Content: input.Content,
		Status:  status,
	}

	err := s.repo.Create(&post)
	return post, err
}

func (s *service) Update(id uint, input UpdatePostInput) (Post, error) {
	post, err := s.repo.FindByID(id)
	if err != nil {
		return post, err
	}

	if input.Title != "" {
		post.Title = input.Title
		post.Slug = generateSlug(input.Title)
	}

	if input.Content != "" {
		post.Content = input.Content
	}

	if input.Status != "" {
		post.Status = input.Status
	}

	err = s.repo.Update(&post)
	return post, err
}

func (s *service) Delete(id uint) error {
	_, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

// generateSlug simplifies a string into a URL-friendly slug
func generateSlug(title string) string {
	slug := strings.ToLower(title)
	// Remove non-alphanumeric chars
	reg := regexp.MustCompile("[^a-z0-9]+")
	slug = reg.ReplaceAllString(slug, "-")
	// Trim dashes
	slug = strings.Trim(slug, "-")
	return slug
}
