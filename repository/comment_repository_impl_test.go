package repository

import (
	"belajar-golang-database/entity"
	"belajar-golang-database/helper"
	"context"
	"fmt"
	"testing"
)

func TestCommentInsert(t *testing.T) {
	commentRepo := NewCommentRepository(helper.GetConnection(), helper.CommentsTable)
	ctx := context.Background()

	comment := entity.Comment{
		Email:   "repo@test.com",
		Comment: "Test Repository",
	}

	result, err := commentRepo.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepo := NewCommentRepository(helper.GetConnection(), helper.CommentsTable)
	ctx := context.Background()

	comment, err := commentRepo.FindById(ctx, 22)
	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
	// fmt.Println(runtime.NumGoroutine())
}

func TestFindAll(t *testing.T) {
	commentRepo := NewCommentRepository(helper.GetConnection(), helper.CommentsTable)

	ctx := context.Background()

	comments, err := commentRepo.FindAll(ctx)

	if err != nil {
		panic(err)
	}

	// pakai _ supaya data nya muncul
	for _, comment := range comments {
		fmt.Println(comment)
	}
}
