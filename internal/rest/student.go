package rest

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/woozie-10/api-clean-arch/domain"
	swaggerFiles "github.com/swaggo/files"
    "github.com/swaggo/gin-swagger"
)

type ResponseError struct {
	Message string `json:"message"`
}
type Response struct {
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
	g.GET("/students", handler.Get)
	g.GET("/students/:username", handler.GetByUsername)
	g.GET("/students/group/:group", handler.GetByGroup)
	g.GET("/students/course/:course", handler.GetByCourse)
	g.POST("/students", handler.Add)
	g.PUT("/students/:username", handler.Update)
	g.DELETE("/students/:username", handler.Delete)

	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
// Get godoc
// @Summary Retrieve a list of students
// @Description Returns a list of all students from the database
// @Tags students
// @Accept json
// @Produce json
// @Success 200 {array} domain.Student "List of students"
// @Failure 500 {object} ResponseError "Internal server error"
// @Router /students [get]
func (s *StudentHandler) Get(c *gin.Context) {
	ctx := c.Request.Context()
	students, err := s.Service.Get(ctx)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, students)
}
// GetByUsername godoc
// @Summary Get student by username
// @Description Retrieve a student by their username from the database
// @Tags students
// @Accept json
// @Produce json
// @Param username path string true "Username of the student to retrieve"
// @Success 200 {object} domain.Student "Student details"
// @Failure 400 {object} ResponseError "Bad request"
// @Failure 500 {object} ResponseError "Internal server error"
// @Router /students/{username} [get]
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
// GetByGroup godoc
// @Summary Get students by group
// @Description Retrieve a list of students by their group from the database
// @Tags students
// @Accept json
// @Produce json
// @Param group path string true "Group of the students to retrieve"
// @Success 200 {array} domain.Student "List of students"
// @Failure 400 {object} ResponseError "Bad request"
// @Failure 500 {object} ResponseError "Internal server error"
// @Router /students/group/{group} [get]
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
// GetByCourse godoc
// @Summary Get students by course
// @Description Retrieve a list of students by their course from the database
// @Tags students
// @Accept json
// @Produce json
// @Param course path string true "Course of the students to retrieve"
// @Success 200 {array} domain.Student "List of students"
// @Failure 400 {object} ResponseError "Bad request"
// @Failure 500 {object} ResponseError "Internal server error"
// @Router /students/course/{course} [get]
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
// Add godoc
// @Summary Add a new student
// @Description Add a new student to the database
// @Tags students
// @Accept json
// @Produce json
// @Param student body domain.Student true "Student object to add"
// @Success 201 {object} domain.Student "Added student"
// @Failure 422 {object} ResponseError "Unprocessable entity"
// @Failure 500 {object} ResponseError "Internal server error"
// @Router /students [post]
func (s *StudentHandler) Add(c *gin.Context) {
	var student *domain.Student

    if err := c.ShouldBindJSON(&student); err != nil {
        c.JSON(http.StatusUnprocessableEntity, ResponseError{err.Error()})
        return
    }

    ctx := c.Request.Context()
    if err := s.Service.Add(ctx, student); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, student)
}
// Update godoc
// @Summary Update a student
// @Description Update an existing student in the database
// @Tags students
// @Accept json
// @Produce json
// @Param username path string true "Username of the student to update"
// @Param newStudent body domain.Student true "New student data"
// @Success 200 {object} domain.Student "Updated student"
// @Failure 422 {object} ResponseError "Unprocessable entity"
// @Failure 500 {object} ResponseError "Internal server error"
// @Router /students/{username} [put]
func (s *StudentHandler) Update(c *gin.Context) {
	var newStudent *domain.Student
	username := c.Param("username")
	if err := c.ShouldBindJSON(&newStudent); err != nil {
        c.JSON(http.StatusUnprocessableEntity, ResponseError{err.Error()})
        return
    }
	ctx := c.Request.Context()
	err := s.Service.Update(ctx, username, newStudent)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, newStudent)
}
// Delete godoc
// @Summary Delete a student
// @Description Delete an existing student from the database
// @Tags students
// @Param username path string true "Username of the student to delete"
// @Success 200 {object} Response "Message indicating successful deletion"
// @Failure 500 {object} ResponseError "Internal server error"
// @Router /students/{username} [delete]
func (s *StudentHandler) Delete(c *gin.Context) {
	username := c.Param("username")
	ctx := c.Request.Context()
	err := s.Service.Delete(ctx, username)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, ResponseError{Message: err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, Response{Message: "the student was successfully removed"})
}
