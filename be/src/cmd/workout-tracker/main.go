package main

import (
	"log"

	"workout-tracker/m/v0.0.0/src/configs"
	"workout-tracker/m/v0.0.0/src/internal/app"
	"workout-tracker/m/v0.0.0/src/internal/domain"
	"workout-tracker/m/v0.0.0/src/internal/routes"
	"workout-tracker/m/v0.0.0/src/scripts"

	"github.com/gin-contrib/cors"
)

func main() {
    models := []interface{}{
        &domain.User{},
        &domain.MuscleGroup{},
        &domain.Exercise{},
        &domain.Workout{},
        &domain.Comment{},
        &domain.ExercisePlan{},
    }

    dbConn, err := configs.ConnectDB()
    if err != nil {
        log.Fatal(err)
    }

    if err = dbConn.Migrator().DropTable("exercise_plans", "exercise_muscle_groups", "workout_reports", "comments", "workouts", "exercises", "muscle_groups", "users"); err != nil {
        log.Fatal("failed to drop table" + err.Error())
    }

    for _, model := range models {
        if err = dbConn.AutoMigrate(model); err != nil {
            log.Fatal("failed migrating table: " + err.Error())
        }
    }

    scripts.Seeder()

    exerciseRepo := domain.NewGormExerciseRepository(dbConn)
    exerciseService := app.NewExerciseService(exerciseRepo)
    exerciseHandler := app.NewExerciseHandler(exerciseService)

    userRepo := domain.NewGormUserRepository(dbConn)
    userService := app.NewUserService(userRepo)
    userHandler := app.NewUserHandler(userService)

    workoutRepo := domain.NewGormWorkoutRepository(dbConn)
    workoutService := app.NewWorkoutService(workoutRepo)
    workoutHandler := app.NewWorkoutHandler(workoutService)

    handler := map[string]interface{}{
        "exerciseHandler": exerciseHandler,
        "userHandler":     userHandler,
        "workoutHandler":  workoutHandler,
    }

    r := configs.SetupRouter()

    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
    }))

    routes.Routes(r, handler)
    r.Run(configs.GetServerPort())
}