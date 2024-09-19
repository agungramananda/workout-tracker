package app

import (
	"net/http"
	"strconv"
	"workout-tracker/m/v0.0.0/src/internal/exceptions"
	"workout-tracker/m/v0.0.0/src/pkg/utils"

	"github.com/gin-gonic/gin"
)

type ExerciseHandler struct {
	ExerciseService *ExerciseService
}

func NewExerciseHandler(es *ExerciseService) *ExerciseHandler {
	return &ExerciseHandler{
		ExerciseService: es,
	}
}

func (eh *ExerciseHandler) GetExercisesHandler(c *gin.Context) {
	exercises, err := eh.ExerciseService.ExerciseRepository.GetExercises()
	if err != nil {
		e := exceptions.NewInvalidRequestError("Failed to get exercises", err)
		errResponse := utils.NewErrorResponse(e.ClientError.Code, e.ClientError.Kind, e.ClientError.Message)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	exercisesInterface := make([]interface{}, len(exercises))
	for i, exercise := range exercises {
		exercisesInterface[i] = exercise
	}
	h := utils.NewResponse(http.StatusOK, "Exercises retrieved successfully", exercisesInterface)
	c.JSON(http.StatusOK, h)
}

func (eh *ExerciseHandler) GetExerciseByIDHandler(c *gin.Context) {
	id := c.Params.ByName("id")
	exerciseID, err := strconv.Atoi(id)
	if err != nil {
		e := exceptions.NewInvalidRequestError("Invalid exercise ID", err)
		errResponse := utils.NewErrorResponse(e.ClientError.Code, e.ClientError.Kind, e.ClientError.Message)
		c.JSON(e.ClientError.Code, errResponse)
		return
	}
	exercise, err := eh.ExerciseService.ExerciseRepository.GetExerciseByID(exerciseID)
	if err != nil {
		e := exceptions.NewInvalidRequestError("Failed to get exercise", err)
		errResponse := utils.NewErrorResponse(e.ClientError.Code, e.ClientError.Kind, e.ClientError.Message)
		c.JSON(e.ClientError.Code, errResponse)
		return
	}
	exerciseInterface := []interface{}{exercise}
	h := utils.NewResponse(http.StatusOK, "Exercise retrieved successfully", exerciseInterface)
	c.JSON(http.StatusOK, h)
}