package assessment

import (
	"context"

	"github.com/woozie-10/api-clean-arch/domain"
)

type AssessmentRepository interface {
	Get(ctx context.Context, username string) (*domain.Assessment, error)
	Add(ctx context.Context, assessment *domain.Assessment) error
}

type AssessmentService struct {
	assessmentRepo AssessmentRepository
}

func NewAssessmentService(assessmentRepo AssessmentRepository) *AssessmentService {
	return &AssessmentService{assessmentRepo: assessmentRepo}
}


func (s *AssessmentService) Get(ctx context.Context, username string) (*domain.Assessment, error) {
	return s.assessmentRepo.Get(ctx, username)
}

func (s *AssessmentService) Add(ctx context.Context, assessment *domain.Assessment) error {
	return s.assessmentRepo.Add(ctx, assessment)
}