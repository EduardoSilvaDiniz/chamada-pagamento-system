// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: query.sql

package db

import (
	"context"
)

const createAssoc = `-- name: CreateAssoc :exec
INSERT INTO associated
(cpf, name, date_birth, marital_status)
VALUES ($1, $2, $3, $4)
`

type CreateAssocParams struct {
	Cpf           int64
	Name          string
	DateBirth     string
	MaritalStatus string
}

func (q *Queries) CreateAssoc(ctx context.Context, arg CreateAssocParams) error {
	_, err := q.db.Exec(ctx, createAssoc,
		arg.Cpf,
		arg.Name,
		arg.DateBirth,
		arg.MaritalStatus,
	)
	return err
}

const getAssoc = `-- name: GetAssoc :exec
SELECT cpf, name, date_birth, marital_status FROM associated
`

func (q *Queries) GetAssoc(ctx context.Context) error {
	_, err := q.db.Exec(ctx, getAssoc)
	return err
}
