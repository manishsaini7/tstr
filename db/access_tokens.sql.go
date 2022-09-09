// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: access_tokens.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const authAccessToken = `-- name: AuthAccessToken :one
SELECT namespace_selectors, scopes::text[], expires_at
FROM access_tokens
WHERE
  token_hash = $1 AND
  (expires_at IS NULL OR expires_at > CURRENT_TIMESTAMP)
`

type AuthAccessTokenRow struct {
	NamespaceSelectors []string
	Scopes             []string
	ExpiresAt          sql.NullTime
}

// TODO re: ::text[] https://github.com/kyleconroy/sqlc/issues/1256
func (q *Queries) AuthAccessToken(ctx context.Context, db DBTX, tokenHash string) (AuthAccessTokenRow, error) {
	row := db.QueryRow(ctx, authAccessToken, tokenHash)
	var i AuthAccessTokenRow
	err := row.Scan(&i.NamespaceSelectors, &i.Scopes, &i.ExpiresAt)
	return i, err
}

const getAccessToken = `-- name: GetAccessToken :one
SELECT id, name, namespace_selectors, scopes::text[], issued_at, expires_at, revoked_at
FROM access_tokens
WHERE id = $1
`

type GetAccessTokenRow struct {
	ID                 uuid.UUID
	Name               string
	NamespaceSelectors []string
	Scopes             []string
	IssuedAt           sql.NullTime
	ExpiresAt          sql.NullTime
	RevokedAt          sql.NullTime
}

// TODO re: ::text[] https://github.com/kyleconroy/sqlc/issues/1256
func (q *Queries) GetAccessToken(ctx context.Context, db DBTX, id uuid.UUID) (GetAccessTokenRow, error) {
	row := db.QueryRow(ctx, getAccessToken, id)
	var i GetAccessTokenRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.NamespaceSelectors,
		&i.Scopes,
		&i.IssuedAt,
		&i.ExpiresAt,
		&i.RevokedAt,
	)
	return i, err
}

const issueAccessToken = `-- name: IssueAccessToken :one
INSERT INTO access_tokens (name, token_hash, namespace_selectors, scopes, expires_at)
VALUES (
  $1, 
  $2, 
  $3,
  $4::text[]::access_token_scope[], 
  $5
)
RETURNING id, name, namespace_selectors, scopes::text[], issued_at, expires_at
`

type IssueAccessTokenParams struct {
	Name               string
	TokenHash          string
	NamespaceSelectors []string
	Scopes             []string
	ExpiresAt          sql.NullTime
}

type IssueAccessTokenRow struct {
	ID                 uuid.UUID
	Name               string
	NamespaceSelectors []string
	Scopes             []string
	IssuedAt           sql.NullTime
	ExpiresAt          sql.NullTime
}

// TODO re: ::text[] https://github.com/kyleconroy/sqlc/issues/1256
func (q *Queries) IssueAccessToken(ctx context.Context, db DBTX, arg IssueAccessTokenParams) (IssueAccessTokenRow, error) {
	row := db.QueryRow(ctx, issueAccessToken,
		arg.Name,
		arg.TokenHash,
		arg.NamespaceSelectors,
		arg.Scopes,
		arg.ExpiresAt,
	)
	var i IssueAccessTokenRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.NamespaceSelectors,
		&i.Scopes,
		&i.IssuedAt,
		&i.ExpiresAt,
	)
	return i, err
}

const listAccessTokens = `-- name: ListAccessTokens :many
SELECT id, name, namespace_selectors, scopes::text[], issued_at, expires_at, revoked_at
FROM access_tokens
WHERE
  CASE WHEN $1::bool
   THEN TRUE
   ELSE expires_at IS NULL OR expires_at > CURRENT_TIMESTAMP
  END AND
  CASE WHEN $2::bool
   THEN TRUE
   ELSE revoked_at IS NULL OR revoked_at > CURRENT_TIMESTAMP
  END
`

type ListAccessTokensParams struct {
	IncludeExpired bool
	IncludeRevoked bool
}

type ListAccessTokensRow struct {
	ID                 uuid.UUID
	Name               string
	NamespaceSelectors []string
	Scopes             []string
	IssuedAt           sql.NullTime
	ExpiresAt          sql.NullTime
	RevokedAt          sql.NullTime
}

// TODO re: ::text[] https://github.com/kyleconroy/sqlc/issues/1256
func (q *Queries) ListAccessTokens(ctx context.Context, db DBTX, arg ListAccessTokensParams) ([]ListAccessTokensRow, error) {
	rows, err := db.Query(ctx, listAccessTokens, arg.IncludeExpired, arg.IncludeRevoked)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListAccessTokensRow
	for rows.Next() {
		var i ListAccessTokensRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.NamespaceSelectors,
			&i.Scopes,
			&i.IssuedAt,
			&i.ExpiresAt,
			&i.RevokedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const revokeAccessToken = `-- name: RevokeAccessToken :exec
UPDATE access_tokens
SET revoked_at = CURRENT_TIMESTAMP
WHERE id = $1
`

func (q *Queries) RevokeAccessToken(ctx context.Context, db DBTX, id uuid.UUID) error {
	_, err := db.Exec(ctx, revokeAccessToken, id)
	return err
}
