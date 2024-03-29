package repository

import (
	"context"
	"database/sql"
	// "log"
)

type GroupRepository struct {
	DB *sql.DB
}

// type DBTX interface {

// }

func NewGroupRepository(db *sql.DB) *GroupRepository {
	return &GroupRepository{
		DB: db,
	}
}

func (g *GroupRepository) Create(ctx context.Context, cl interface{}, roomID string) error {
	tx, err := g.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()
	
	_, err = tx.ExecContext(ctx, "INSERT INTO groups (id, name, created_at, owner_username) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, "INSERT INTO follow_groups (username, id_group) VALUES (?, ?)", cl.Username, roomID)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (g *GroupRepository) Join(ctx context.Context, username string, roomID string) error {
	tx, err := g.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, "INSERT INTO follow_groups (username, id_group) VALUES (?, ?)", username, roomID)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (g *GroupRepository) AlreadyJoin(ctx context.Context, username string, roomID string) (bool, error) {
	tx, err := g.DB.BeginTx(ctx, nil)
	if err != nil {
		return false, err
	}

	row := tx.QueryRowContext(ctx, "SELECT * FROM follow_groups WHERE username = ? AND id_group = ?", username, roomID)

	if err = row.Err(); err != nil {
		return false, row.Err()
	}

	if err = tx.Commit(); err != nil {
		return false, err
	}

	return true, nil
}