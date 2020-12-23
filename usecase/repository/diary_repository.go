package repository

import (
	"context"

	"github.com/WomenWhoGoTokyo/shoten10-diary-application/domain/model"
)

type DiaryRepository interface {
	Store(ctx context.Context, diary *model.Diary) (*model.Diary, error)
	Update(ctx context.Context, diary *model.Diary) (*model.Diary, error)
	Delete(ctx context.Context, id int) error
	FindAll(ctx context.Context) ([]*model.Diary, error)
	FindByID(ctx context.Context, id int) (*model.Diary, error)
}
