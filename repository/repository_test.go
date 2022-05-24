package repository

import (
	"context"
	"fmt"
	"godatabase"
	"godatabase/entity"
	"testing"
)

func TestJobInsert(t *testing.T) {
	jobRepository := NewJobRepository(godatabase.GetConnections())

	ctx := context.Background()
	job := entity.Job{
		Title:        "title test",
		Organization: 1,
		Sequence:     1,
		Winner:       1,
	}

	result, err := jobRepository.Insert(ctx, job)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestJobFindById(t *testing.T) {
	jobRepository := NewJobRepository(godatabase.GetConnections())

	ctx := context.Background()
	job, err := jobRepository.FindById(ctx, 37)

	if err != nil {
		panic(err)
	}

	fmt.Println(job)
}

func TestJobFindaLL(t *testing.T) {
	jobRepository := NewJobRepository(godatabase.GetConnections())
	ctx := context.Background()
	jobs, err := jobRepository.FindAll(ctx)

	if err != nil {
		panic(err)
	}

	for _, job := range jobs {
		fmt.Println(job)
	}
}
