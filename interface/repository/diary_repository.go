package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/WomenWhoGoTokyo/shoten10-diary-application/domain/model"
)

type diaryRepository struct {
	db *sql.DB
}

func NewDiaryRepository(db *sql.DB) *diaryRepository {
	return &diaryRepository{db}
}

func (dr *diaryRepository) Store(ctx context.Context, diary *model.Diary) (*model.Diary, error) {
	s := `INSERT INTO diaries (title, description) VALUES ($1, $2) RETURNING id`
	if err := dr.db.QueryRowContext(ctx, s, diary.Title, diary.Description).Scan(&diary.ID); err != nil {
		return nil, err
	}
	return diary, nil
}

func (dr *diaryRepository) Update(ctx context.Context, diary *model.Diary) (*model.Diary, error) {
	s := `UPDATE diaries SET title = $1, description = $2 WHERE id = $3`
	result, err := dr.db.ExecContext(ctx, s,
		diary.Title,
		diary.Description,
		diary.ID,
	)
	if err != nil {
		return nil, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rows != 1 {
		return nil, fmt.Errorf("expected to affect 1 row, affected %d", rows)
	}
	return diary, nil
}

func (dr *diaryRepository) Delete(ctx context.Context, id int) error {
	s := `DELETE FROM diaries WHERE id = $1`
	result, err := dr.db.ExecContext(ctx, s, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return fmt.Errorf("expected to affect 1 row, affected %d", rows)
	}
	return nil
}

func (dr *diaryRepository) FindAll(ctx context.Context) ([]*model.Diary, error) {
	s := `SELECT id, title, description FROM diaries`
	rows, err := dr.db.QueryContext(ctx, s)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	diaries := make([]*model.Diary, 0)

	for rows.Next() {
		diary := &model.Diary{}
		if err := rows.Scan(&diary.ID, &diary.Title, &diary.Description); err != nil {
			return nil, err
		}
		diaries = append(diaries, diary)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return diaries, nil
}

func (dr *diaryRepository) FindByID(ctx context.Context, id int) (*model.Diary, error) {
	diary := &model.Diary{}
	s := `SELECT id, title, description FROM diaries WHERE id = $1`
	if err := dr.db.QueryRowContext(ctx, s, id).Scan(&diary.ID, &diary.Title, &diary.Description); err != nil {
		return nil, err
	}
	return diary, nil
}
