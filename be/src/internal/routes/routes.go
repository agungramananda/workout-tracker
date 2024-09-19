package routes

import (
	"workout-tracker/m/v0.0.0/src/internal/app"
	"workout-tracker/m/v0.0.0/src/internal/exceptions"
	"workout-tracker/m/v0.0.0/src/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, handlers map[string]interface{}) {
	r.NoRoute(func(c *gin.Context) {
		err := exceptions.NewNotFoundError("Route not found", nil)
		c.JSON(err.ClientError.Code, gin.H{
			"error":   err.ClientError.Kind,
			"message": err.ClientError.Message,
		})
	})

	userHandler, ok := handlers["userHandler"].(*app.UserHandler)
	if ok {
		r.POST("/login", userHandler.LoginHandler)
		r.POST("/register", userHandler.RegisterHandler)
		userGroup := r.Group("/users")
		{
			userGroup.Use(middleware.JwtAuthMiddleware())
			{
				userGroup.GET("/profile", userHandler.GetUserByIDHandler)
				userGroup.PUT("/profile", userHandler.UpdateUserHandler)
			}
		}
	}

	exerciseHandler, ok := handlers["exerciseHandler"].(*app.ExerciseHandler)
	if ok {
		exerciseGroup := r.Group("/exercises")
		{
			exerciseGroup.Use(middleware.JwtAuthMiddleware())
			{
				exerciseGroup.GET("", exerciseHandler.GetExercisesHandler)
				exerciseGroup.GET("/:id", exerciseHandler.GetExerciseByIDHandler)
			}
		}
	}

	workoutHandler, ok := handlers["workoutHandler"].(*app.WorkoutHandler)
	if ok {
		workoutGroup := r.Group("/workouts")
		{
			workoutGroup.Use(middleware.JwtAuthMiddleware())
			{
				workoutGroup.POST("", workoutHandler.CreateWorkoutHandler)
				workoutGroup.GET("", workoutHandler.GetWorkoutsHandler)
				workoutGroup.GET("/:id", workoutHandler.GetWorkoutByIDHandler)
				workoutGroup.PUT("/:id", workoutHandler.UpdateWorkoutHandler)
				workoutGroup.DELETE("/:id", workoutHandler.DeleteWorkoutHandler)
			}
		}
	}
}
