package services

import (
	"github.com/Danchitomoo/go_api_learning/models"
	"github.com/Danchitomoo/go_api_learning/repositories"
)

func (s MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {
	newComment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		return models.Comment{}, err
	}

	return newComment, nil
}
