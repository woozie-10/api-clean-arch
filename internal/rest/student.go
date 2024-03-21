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
	g.GET("", handler.Get)
	g.GET("/:username", handler.GetByUsername)
	g.GET("/group/:group", handler.GetByGroup)
	g.GET("course/:course", handler.GetByCourse)
	g.POST("", handler.Add)
	g.PUT("/:username", handler.Update)
	g.DELETE("/:username", handler.Delete)
}
func (s *StudentHandler) Get(c *gin.Context) {
	ctx := c.Request.Context()
	students, err := s.Service.Get(ctx)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, students)
}

func (s *StudentHandler) GetByUsername(c *gin.Context) {
	username := c.Param("username")
	ctx := c.Request.Context()
	student, err := s.Service.GetByUsername(ctx, username)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, student)
}

func (s *StudentHandler) GetByGroup(c *gin.Context) {
	group := c.Param("group")
	ctx := c.Request.Context()
	students, err := s.Service.GetByGroup(ctx, group)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, students)
}

func (s *StudentHandler) GetByCourse(c *gin.Context) {
	course := c.Param("course")
	ctx := c.Request.Context()
	students, err := s.Service.GetByCourse(ctx, course)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, students)
}

func (s *StudentHandler) Add(c *gin.Context) {
	var student *domain.Student
	err := c.Bind(&student)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, ResponseError{Message: err.Error()})
	}
	ctx := c.Request.Context()
	err = s.Service.Add(ctx, student)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, student)
}

func (s *StudentHandler) Update(c *gin.Context) {
	var newStudent *domain.Student
	username := c.Param("username")
	err := c.Bind(&newStudent)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, ResponseError{Message: err.Error()})
	}
	ctx := c.Request.Context()
	err = s.Service.Update(ctx, username, newStudent)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, newStudent)
}

func (s *StudentHandler) Delete(c *gin.Context) {
	username := c.Param("username")
	ctx := c.Request.Context()
	err := s.Service.Delete(ctx, username)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "the student was successfully removed"})
}