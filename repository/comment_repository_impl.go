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

func (repository *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	query := "INSERT into " + repository.Comments_TB + "(email, comment) VALUES (?,?)"
	_, err := repository.DB.ExecContext(ctx, query, comment.Email, comment.Comment)

	if err != nil {
		return comment, err
	} else {
		return comment, nil
	}
}

func (repository *commentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	query := "SELECT id, email, comment FROM " + repository.Comments_TB + "WHERE id = ?"
	rows, err := repository.DB.QueryContext(ctx, query, id)
	defer rows.Close()

	comment := entity.Comment{}

	if err != nil {
		return comment, err
	}

	if rows.Next() {
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
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
