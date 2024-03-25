package student

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

type Service struct {
	studentRepo StudentRepository
}

func NewService(sr StudentRepository) *Service {
	return &Service{
		studentRepo: sr,
	}
}

func (s *Service) Get(ctx context.Context) ([]*domain.Student, error) {
	return s.studentRepo.Get(ctx)
}
func (s *Service) GetByUsername(ctx context.Context, username string) (*domain.Student, error) {
	return s.studentRepo.GetByUsername(ctx, username)
}
func (s *Service) Add(ctx context.Context, student *domain.Student) error {
	return s.studentRepo.Add(ctx, student)
}
func (s *Service) Delete(ctx context.Context, username string) error {
	return s.studentRepo.Delete(ctx, username)
}
func (s *Service) GetByGroup(ctx context.Context, group string) ([]*domain.Student, error) {
	return s.studentRepo.GetByGroup(ctx, group)
}
func (s *Service) GetByCourse(ctx context.Context, course string) ([]*domain.Student, error) {
	return s.studentRepo.GetByCourse(ctx, course)
}
func (s *Service) Update(ctx context.Context, username string, newStudent *domain.Student) error {
	return s.studentRepo.Update(ctx, username, newStudent)
}
