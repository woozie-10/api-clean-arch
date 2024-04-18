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

type AssessmentRepository interface {
	Get(ctx context.Context, username string) (*domain.Assessment, error)
	Add(ctx context.Context, assessment *domain.Assessment) error
}

type StudentService struct {
	studentRepo StudentRepository
}

func NewStudentService(studentRepo StudentRepository) *StudentService {
	return &StudentService{studentRepo: studentRepo}
}

type AssessmentService struct {
	assessmentRepo AssessmentRepository
}

func NewAssessmentService(assessmentRepo AssessmentRepository) *AssessmentService {
	return &AssessmentService{assessmentRepo: assessmentRepo}
}

func (s *StudentService) Get(ctx context.Context) ([]*domain.Student, error) {
	return s.studentRepo.Get(ctx)
}
func (s *StudentService) GetByUsername(ctx context.Context, username string) (*domain.Student, error) {
	return s.studentRepo.GetByUsername(ctx, username)
}
func (s *StudentService) Add(ctx context.Context, student *domain.Student) error {
	return s.studentRepo.Add(ctx, student)
}
func (s *StudentService) Delete(ctx context.Context, username string) error {
	return s.studentRepo.Delete(ctx, username)
}
func (s *StudentService) GetByGroup(ctx context.Context, group string) ([]*domain.Student, error) {
	return s.studentRepo.GetByGroup(ctx, group)
}
func (s *StudentService) GetByCourse(ctx context.Context, course string) ([]*domain.Student, error) {
	return s.studentRepo.GetByCourse(ctx, course)
}
func (s *StudentService) Update(ctx context.Context, username string, newStudent *domain.Student) error {
	return s.studentRepo.Update(ctx, username, newStudent)
}

func (s *AssessmentService) Get(ctx context.Context, username string) (*domain.Assessment, error) {
	return s.assessmentRepo.Get(ctx, username)
}

func (s *AssessmentService) Add(ctx context.Context, assessment *domain.Assessment) error {
	return s.assessmentRepo.Add(ctx, assessment)
}
