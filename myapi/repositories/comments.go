package repositories

import (
	"database/sql"

	"github.com/Danchitomoo/go_api_learning/models"
)

func InsertComment(db *sql.DB, comment models.Comment) (models.Comment, error) {
	const sqlStr = `
		insert into comments (article_id, message, created_at) values
		?, ?, now();
	`
	var newComment models.Comment
	result, err := db.Exec(sqlStr, comment.ArticleID, comment.Message)
	if err != nil {
		return models.Comment{}, nil
	}
	newCommentID, err := result.LastInsertId()
	if err != nil {
		return models.Comment{}, nil
	}
	newComment.CommentID, newComment.ArticleID, newComment.Message = int(newCommentID), comment.ArticleID, comment.Message

	return newComment, nil
}

func SelectCommentList(db *sql.DB, articleID int) ([]models.Comment, error) {
	const sqlStr = `
		select * from comments
		where article_id = ?;
	`
	rows, err := db.Query(sqlStr, articleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	commentArray := make([]models.Comment, 0)
	for rows.Next() {
		var comment models.Comment
		var createdAt sql.NullTime
		rows.Scan(
			&comment.CommentID,
			&comment.ArticleID,
			&comment.Message,
			createdAt,
		)
		if createdAt.Valid {
			comment.CreatedAt = createdAt.Time
		}
		commentArray = append(commentArray, comment)
	}

	return commentArray, nil
}
