package comment

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

type Service struct {
	Db *gorm.DB
}

// Comment - struct
type Comment struct {
	gorm.Model
	Slug    string
	Body    string
	Author  string
	Created time.Time
}

// CommentService - implement an interface with all method attached to comment
type CommentService interface {
	GetComment(ID uint) (Comment, error)
	PostComment(comment Comment) (Comment, error)
	UpdateComment(ID uint, newComment Comment) (Comment, error)
	DeleteComment(ID uint) error
	GetAllComments() ([]Comment, error)
	GetCommentBySlug(slug string) ([]Comment, error)
}

func Newcomment(db *gorm.DB) *Service {
	return &Service{
		Db: db,
	}
}

// GetComment - return a comment
func (s *Service) GetComment(ID uint) (Comment, error) {
	var comment Comment
	if result := s.Db.First(&comment, ID); result.Error != nil {
		return Comment{}, result.Error
	}
	return comment, nil
}

// PostComment - creates a new comment
func (s *Service) PostComment(comment Comment) (Comment, error) {
	if result := s.Db.Save(&comment); result.Error != nil {
		return Comment{}, result.Error
	}
	return comment, nil
}

// UpdateComment - update an existing comment
func (s *Service) UpdateComment(ID uint, newComment Comment) (Comment, error) {
	comment, err := s.GetComment(ID)
	if err != nil {
		fmt.Printf("Comment with ID %d does not exist", ID)
	}
	if result := s.Db.Model(&comment).Updates(newComment); err != nil {
		return Comment{}, result.Error
	}
	return comment, nil
}

// GetAllComments - return all comments
func (s *Service) GetAllComments() ([]Comment, error) {
	var comment []Comment
	if result := s.Db.Find(&comment); result.Error != nil {
		return nil, result.Error
	}
	return comment, nil
}

// DeleteComment -  delete a comment by ID
func (s *Service) DeleteComment(ID uint) error {
	if result := s.Db.Delete(&Comment{}, ID); result.Error != nil {
		return result.Error
	}
	return nil
}

// GetCommentBySlug - return a list of comment
func (s *Service) GetCommentBySlug(slug string) ([]Comment, error) {
	var comment []Comment
	if result := s.Db.Find(&comment).Where("slug = ?", slug); result.Error != nil {
		return []Comment{}, result.Error
	}
	return comment, nil
}
