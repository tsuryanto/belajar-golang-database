package repository

import (
	"belajar-golang-database/entity"
	"context"
	"database/sql"
	"errors"
	"strconv"
)

type commentRepositoryImpl struct {
	DB          *sql.DB
	Comments_TB string
}

func NewCommentRepository(db *sql.DB, comments_tb string) commentRepository {
	return &commentRepositoryImpl{DB: db, Comments_TB: comments_tb}
}

func (repository *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	query := "INSERT into " + repository.Comments_TB + "(email, comment) VALUES (?,?)"
	result, err := repository.DB.ExecContext(ctx, query, comment.Email, comment.Comment)

	if err != nil {
		return comment, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}

	comment.Id = int32(id)
	return comment, nil
}

func (repository *commentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	// Error
	query := "SELECT id, email, comment FROM " + repository.Comments_TB + " WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, query, id)

	// Error too
	// table := "comments"
	// query := "SELECT id, email, comment FROM" + table + "WHERE id = ? LIMIT 1"
	// rows, err := repository.DB.QueryContext(ctx, query, id)

	// No Error
	// query := "SELECT id, email, comment FROM comments WHERE id = ? LIMIT 1"
	// rows, err := repository.DB.QueryContext(ctx, query, id)

	comment := entity.Comment{}

	if err != nil {
		return comment, err
	}
	defer rows.Close()

	if rows.Next() {
		// ada
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		// tidak ada
		return comment, errors.New("Id " + strconv.Itoa(int(id)) + " Not Found")
	}

}

func (repository *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	query := "SELECT id, email, comment FROM " + repository.Comments_TB
	rows, err := repository.DB.QueryContext(ctx, query)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	comments := []entity.Comment{}

	for rows.Next() {
		comment := entity.Comment{}
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}

	return comments, nil
}
