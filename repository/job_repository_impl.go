package repository

import (
	"context"
	"database/sql"
	"errors"
	"godatabase/entity"
	"strconv"
)

type jobRepositoryImpl struct {
	DB *sql.DB
}

func NewJobRepository(db *sql.DB) JobRepository {
	return &jobRepositoryImpl{DB: db}
}

func (repository *jobRepositoryImpl) Insert(ctx context.Context, job entity.Job) (entity.Job, error) {
	script := "INSERT INTO job (title, organization, sequence, winner) VALUES (?, ?, ?, ?)"
	result, err := repository.DB.ExecContext(ctx, script, job.Title, job.Organization, job.Sequence, job.Winner)
	if err != nil {
		panic(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return job, err
	}

	job.Id = int32(id)
	return job, nil
}

func (repository *jobRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Job, error) {
	script := "SELECT id_job as id, title, organization, sequence, winner FROM job WHERE id_job = ? limit 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	job := entity.Job{}
	if err != nil {
		return job, err
	}

	defer rows.Close()
	if rows.Next() {
		rows.Scan(&job.Id, &job.Title, &job.Organization, &job.Sequence, &job.Winner)
		return job, nil
	} else {
		return job, errors.New("Id " + strconv.Itoa(int(id)) + " Not Found")
	}
}

func (repository *jobRepositoryImpl) FindAll(ctx context.Context) ([]entity.Job, error) {
	script := "SELECT id_job as id, title, organization, sequence, winner FROM job"
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var jobs []entity.Job

	for rows.Next() {
		job := entity.Job{}
		rows.Scan(&job.Id, &job.Title, &job.Organization, &job.Sequence, &job.Winner)
		jobs = append(jobs, job)
	}

	return jobs, nil
}
