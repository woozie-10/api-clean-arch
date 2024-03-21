package rest

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/woozie-10/api-clean-arch/domain"
)

type ResponseError struct {
	Message string `json:"message"`
}

type StudentService interface {
	Get(ctx context.Context) ([]*domain.Student, error)
	GetByUsername(ctx context.Context, username string) (*domain.Student, error)
	Add(ctx context.Context, student *domain.Student) error
	Delete(ctx context.Context, username string) error
	GetByGroup(ctx context.Context, group string) ([]*domain.Student, error)
	GetByCourse(ctx context.Context, course string) ([]*domain.Student, error)
	Update(ctx context.Context, username string, newStudent *domain.Student) error
}

type StudentHandler struct {
	Service StudentService
}

func NewStudentHandler(g *gin.Engine, svc StudentService) {
	handler := &StudentHandler{
		Service: svc,
	}

}
func (s *StudentHandler) Get(c *gin.Context) {
	ctx := c.Request.Context()
	students, err := s.Service.Get(ctx)
	if err != nil {
		c.IndentedJSON(getStatusCode(err), ResponseError{Message: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, students)
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	switch err {
	case domain.ErrInternalServerError:
		return http.StatusInternalServerError
	case domain.ErrNotFound:
		return http.StatusNotFound
	case domain.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
