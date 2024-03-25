package repository

type GroupRepository struct {
	DB *sql.DB
}

func (g *GroupRepository) Create() error {
	tx, err := g.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	err = tx.QueryRowContext(ctx, "INSERT INTO groups (id, name, created_at, owner_username) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}

	err = tx.ExecContext(ctx, "INSERT INTO follow_groups (username, id_group) VALUES (?, ?)")
	if err != nil {
		return err
	}


}