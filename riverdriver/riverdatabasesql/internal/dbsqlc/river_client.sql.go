// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: river_client.sql

package dbsqlc

import (
	"context"
	"time"
)

const clientCreateOrSetUpdatedAt = `-- name: ClientCreateOrSetUpdatedAt :one
INSERT INTO river_client (
    id,
    metadata,
    paused_at,
    updated_at
) VALUES (
    $1,
    coalesce($2::jsonb, '{}'::jsonb),
    coalesce($3::timestamptz, NULL),
    coalesce($4::timestamptz, now())
) ON CONFLICT (name) DO UPDATE
SET
    updated_at = coalesce($4::timestamptz, now())
RETURNING id, created_at, metadata, paused_at, updated_at
`

type ClientCreateOrSetUpdatedAtParams struct {
	ID        string
	Metadata  string
	PausedAt  *time.Time
	UpdatedAt *time.Time
}

func (q *Queries) ClientCreateOrSetUpdatedAt(ctx context.Context, db DBTX, arg *ClientCreateOrSetUpdatedAtParams) (*RiverClient, error) {
	row := db.QueryRowContext(ctx, clientCreateOrSetUpdatedAt,
		arg.ID,
		arg.Metadata,
		arg.PausedAt,
		arg.UpdatedAt,
	)
	var i RiverClient
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.Metadata,
		&i.PausedAt,
		&i.UpdatedAt,
	)
	return &i, err
}
