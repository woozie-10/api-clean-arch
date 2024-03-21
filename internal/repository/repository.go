package repository

import (
	"context"

	"github.com/woozie-10/api-clean-arch/domain"
)

type StudentRepository interface {
	Get(ctx context.Context) ([]*domain.Student, error)
	GetByUsername(ctx context.Context, username string) (*domain.Student, error)
	Add(ctx context.Context, student *domain.Student) error
	Delete(ctx context.Context, username string) error
	GetByGroup(ctx context.Context, group string) ([]*domain.Student, error)
	GetByCourse(ctx context.Context, course string) ([]*domain.Student, error)
	Update(ctx context.Context, username string, newStudent *domain.Student) error
}
