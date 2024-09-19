package app

import (
	"net/http"
	"strconv"

	"time"
	"workout-tracker/m/v0.0.0/src/internal/domain"
	"workout-tracker/m/v0.0.0/src/internal/exceptions"
	"workout-tracker/m/v0.0.0/src/pkg/utils"

	"github.com/gin-gonic/gin"
)

type WorkoutHandler struct {
	WorkoutService *WorkoutService
}

func NewWorkoutHandler(ws *WorkoutService) *WorkoutHandler {
	return &WorkoutHandler{
		WorkoutService: ws,
	}
}

type WorkoutResponse struct {
	ID                   int `json:"id"`
	UserID               int `json:"user_id"`
	Name                 string `json:"name"`
	Description          string `json:"description"`
	Date                 time.Time `json:"date"`
	Time                 time.Time `json:"time"`
	RestBetweenExercises int `json:"rest_between_exercises"`
	IsCompleted          bool `json:"is_completed"`
	ExercisesPlan        []struct {
		ExerciseID  int `json:"exercise_id"`
		ExerciseName string `json:"exercise_name"`
		Sets        int `json:"sets"`
		Reps        int `json:"reps"`
		Weight      int `json:"weight"` // in kg
		RestTime    int `json:"rest_time"` // in seconds
		Order       int `json:"order"`
		IsCompleted bool `json:"is_completed"`
	} `json:"exercises_plan"`
	Comments []struct {
		UserID  int `json:"user_id"`
		Comment string `json:"comment"`
	} `json:"comments"`
}

func (wh *WorkoutHandler) CreateWorkoutHandler(c *gin.Context) {
	var requestBody struct {
			Name                 string `json:"name"`
			Description          string `json:"description"`
			Date                 string `json:"date"` // Use string if you want to parse it later
			Time                 string `json:"time"` // Use string if you want to parse it later
			RestBetweenExercises int `json:"rest_between_exercises"`
			IsCompleted          bool `json:"is_completed"`
			ExercisesPlan        []struct {
					ExerciseID  int `json:"exercise_id"`
					Sets        int `json:"sets"`
					Reps        int `json:"reps"`
					Weight      int `json:"weight"` // in kg
					RestTime    int `json:"rest_time"` // in seconds
					Order       int `json:"order"`
					IsCompleted bool `json:"is_completed"`
			} `json:"exercises_plan"`
			Comments []struct {
					UserID  int `json:"user_id"`
					Comment string `json:"comment"`
			} `json:"comments"`
	}

	userID, err := utils.ExtractTokenID(c)
	if err != nil {
		e := exceptions.NewInvalidRequestError("Failed to extract token ID", err)
		errResponse := utils.NewErrorResponse(e.ClientError.Code, e.ClientError.Kind, e.ClientError.Message)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
			e := exceptions.NewInvalidRequestError("Invalid workout data", err)
			errResponse := utils.NewErrorResponse(e.ClientError.Code, e.ClientError.Kind, e.ClientError.Message)
			c.JSON(http.StatusBadRequest, errResponse)
			return
	}

	// Parse date and time
	date, err := time.Parse("2006-01-02", requestBody.Date)
	if err != nil {
			e := exceptions.NewInvalidRequestError("Invalid date format", err)
			errResponse := utils.NewErrorResponse(e.ClientError.Code, e.ClientError.Kind, e.ClientError.Message)
			c.JSON(http.StatusBadRequest, errResponse)
			return
	}

	time, err := time.Parse("15:04", requestBody.Time)
	if err != nil {
			e := exceptions.NewInvalidRequestError("Invalid time format", err)
			errResponse := utils.NewErrorResponse(e.ClientError.Code, e.ClientError.Kind, e.ClientError.Message)
			c.JSON(http.StatusBadRequest, errResponse)
			return
	}

	// Convert requestBody to domain model
	workout := domain.Workout{
			UserID:               int(userID),
			Name:                 requestBody.Name,
			Description:          requestBody.Description,
			Date:                 date,
			Time:                 time,
			RestBetweenExercises: requestBody.RestBetweenExercises,
			IsCompleted:          requestBody.IsCompleted,
	}

	for _, ep := range requestBody.ExercisesPlan {
			workout.ExercisesPlan = append(workout.ExercisesPlan, domain.ExercisePlan{
					ExerciseID:      ep.ExerciseID,
					Sets:            ep.Sets,
					Reps:            ep.Reps,
					Weight:          &ep.Weight,
					RestBetweenSets: ep.RestTime,
					Order:           ep.Order,
					IsCompleted:     ep.IsCompleted,
			})
	}

	for _, cm := range requestBody.Comments {
			workout.Comments = append(workout.Comments, domain.Comment{
					UserID:  cm.UserID,
					Comment: cm.Comment,
			})
	}

	if err := wh.WorkoutService.WorkoutRepository.CreateWorkout(&workout); err != nil {
			e := exceptions.NewInvalidRequestError(err.Error(), err)
			errResponse := utils.NewErrorResponse(e.ClientError.Code, e.ClientError.Kind, e.ClientError.Message)
			c.JSON(http.StatusBadRequest, errResponse)
			return
	}
	c.String(http.StatusCreated, "Workout created successfully")
}

func (wh *WorkoutHandler) GetWorkoutsHandler(c *gin.Context) {
	userId, err := utils.ExtractTokenID(c)
	if err != nil {
		e := exceptions.NewInvalidRequestError("Failed to extract token ID", err)
		errResponse := utils.NewErrorResponse(e.ClientError.Code, e.ClientError.Kind, e.ClientError.Message)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	workouts, err := wh.WorkoutService.WorkoutRepository.GetWorkouts(int(userId))
	if err != nil {
		e := exceptions.NewInvalidRequestError("Failed to get workouts", err)
		errResponse := utils.NewErrorResponse(e.ClientError.Code, e.ClientError.Kind, e.ClientError.Message)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	if len(workouts) == 0 {
		h := utils.NewResponse(http.StatusOK, "No workouts found", []interface{}{})
		c.JSON(http.StatusOK, h)
		return
	}

	var workoutsResponse []WorkoutResponse
	for _, workout := range workouts {
		workoutResponse := WorkoutResponse{
			ID:                   int(workout.ID),
			UserID:               workout.UserID,
			Name:                 workout.Name,
			Description:          workout.Description,
			Date:                 workout.Date,
			Time:                 workout.Time,
			RestBetweenExercises: workout.RestBetweenExercises,
			IsCompleted:          workout.IsCompleted,
		}

		for _, ep := range workout.ExercisesPlan {
			workoutResponse.ExercisesPlan = append(workoutResponse.ExercisesPlan, struct {
				ExerciseID  int `json:"exercise_id"`
				ExerciseName string `json:"exercise_name"`
				Sets        int `json:"sets"`
				Reps        int `json:"reps"`
				Weight      int `json:"weight"` // in kg
				RestTime    int `json:"rest_time"` // in seconds
				Order       int `json:"order"`
				IsCompleted bool `json:"is_completed"`
			}{
				ExerciseID:  ep.ExerciseID,
				ExerciseName: ep.Exercise.Name,
				Sets:        ep.Sets,
				Reps:        ep.Reps,
				Weight:      *ep.Weight,
				RestTime:    ep.RestBetweenSets,
				Order:       ep.Order,
				IsCompleted: ep.IsCompleted,
			})
		}

		for _, cm := range workout.Comments {
			workoutResponse.Comments = append(workoutResponse.Comments, struct {
				UserID  int `json:"user_id"`
				Comment string `json:"comment"`
			}{
				UserID:  cm.UserID,
				Comment: cm.Comment,
			})
		}

		workoutsResponse = append(workoutsResponse, workoutResponse)
	}

	var response []interface{}
	for _, wr := range workoutsResponse {
		response = append(response, wr)
	}
	h := utils.NewResponse(http.StatusOK, "Workouts retrieved successfully", response)
	c.JSON(http.StatusOK, h)
}

func (wh *WorkoutHandler) GetWorkoutByIDHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		e := exceptions.NewInvalidRequestError("Invalid workout ID", err)
		errResponse := utils.NewErrorResponse(e.ClientError.Code, e.ClientError.Kind, e.ClientError.Message)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	userId, err := utils.ExtractTokenID(c)
	if err != nil {
		e := exceptions.NewInvalidRequestError("Failed to extract token ID", err)
		errResponse := utils.NewErrorResponse(e.ClientError.Code, e.ClientError.Kind, e.ClientError.Message)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	workout, err := wh.WorkoutService.WorkoutRepository.GetWorkoutByID(id, int(userId))
	if err != nil {
		e := exceptions.NewInvalidRequestError("Failed to get workout", err)
		errResponse := utils.NewErrorResponse(e.ClientError.Code, e.ClientError.Kind, e.ClientError.Message)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}
	workoutResponse := WorkoutResponse{
		ID:                   int(workout.ID),
		UserID:               workout.UserID,
		Name:                 workout.Name,
		Description:          workout.Description,
		Date:                 workout.Date,
		Time:                 workout.Time,
		RestBetweenExercises: workout.RestBetweenExercises,
		IsCompleted:          workout.IsCompleted,
	}

	for _, ep := range workout.ExercisesPlan {
		workoutResponse.ExercisesPlan = append(workoutResponse.ExercisesPlan, struct {
			ExerciseID  int `json:"exercise_id"`
			ExerciseName string `json:"exercise_name"`
			Sets        int `json:"sets"`
			Reps        int `json:"reps"`
			Weight      int `json:"weight"` // in kg
			RestTime    int `json:"rest_time"` // in seconds
			Order       int `json:"order"`
			IsCompleted bool `json:"is_completed"`
		}{
			ExerciseID:  ep.ExerciseID,
			ExerciseName: ep.Exercise.Name,
			Sets:        ep.Sets,
			Reps:        ep.Reps,
			Weight:      *ep.Weight,
			RestTime:    ep.RestBetweenSets,
			Order:       ep.Order,
			IsCompleted: ep.IsCompleted,
		})
	}

	for _, cm := range workout.Comments {
		workoutResponse.Comments = append(workoutResponse.Comments, struct {
			UserID  int `json:"user_id"`
			Comment string `json:"comment"`
		}{
			UserID:  cm.UserID,
			Comment: cm.Comment,
		})
	}

	h := utils.NewResponse(http.StatusOK, "Workout retrieved successfully", []interface{}{workoutResponse})
	c.JSON(http.StatusOK, h)
}

func (wh *WorkoutHandler) UpdateWorkoutHandler(c *gin.Context) {
	var workout domain.Workout
	if err := c.ShouldBindJSON(&workout); err != nil {
		e := exceptions.NewInvalidRequestError("Invalid workout data", err)
		errResponse := utils.NewErrorResponse(e.ClientError.Code, e.ClientError.Kind, e.ClientError.Message)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	if err := wh.WorkoutService.WorkoutRepository.UpdateWorkout(&workout); err != nil {
		e := exceptions.NewInvalidRequestError("Failed to update workout", err)
		errResponse := utils.NewErrorResponse(e.ClientError.Code, e.ClientError.Kind, e.ClientError.Message)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	c.String(http.StatusOK, "Workout updated successfully")
}

func (wh *WorkoutHandler) DeleteWorkoutHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		e := exceptions.NewInvalidRequestError("Invalid workout ID", err)
		errResponse := utils.NewErrorResponse(e.ClientError.Code, e.ClientError.Kind, e.ClientError.Message)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	if err := wh.WorkoutService.WorkoutRepository.DeleteWorkout(id); err != nil {
		e := exceptions.NewInvalidRequestError("Failed to delete workout", err)
		errResponse := utils.NewErrorResponse(e.ClientError.Code, e.ClientError.Kind, e.ClientError.Message)
		c.JSON(http.StatusBadRequest, errResponse)
		return
	}

	c.String(http.StatusOK, "Workout deleted successfully")
}
