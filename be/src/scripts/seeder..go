package scripts

import (
	"log"
	"workout-tracker/m/v0.0.0/src/configs"
	"workout-tracker/m/v0.0.0/src/internal/domain"
	"workout-tracker/m/v0.0.0/src/pkg/utils"
)

func Seeder() {
	dbConn, err := configs.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	hashedPassword, err := utils.HashPasword("asep123")
	if err != nil {
		log.Fatal(err)
	}
	var user = &domain.User{
		Username: "asep",
		Password: hashedPassword,
		Fullname: "Asep Saepudin",
		Email:    "asep@gmail.com",
	}

	if err := dbConn.Create(user).Error; err != nil {
		log.Fatal(err)
	}

	muscleGroup := []*domain.MuscleGroup{
		{MuscleGroup: "Chest"},
		{MuscleGroup: "Shoulders"},
		{MuscleGroup: "Biceps"},
		{MuscleGroup: "Triceps"},
		{MuscleGroup: "Quadriceps"},
		{MuscleGroup: "Hamstrings"},
		{MuscleGroup: "Calves"},
		{MuscleGroup: "Abs"},
		{MuscleGroup: "Glutes"},
		{MuscleGroup: "Forearms"},
		{MuscleGroup: "Traps"},
		{MuscleGroup: "Lats"},
		{MuscleGroup: "Obliques"},
		{MuscleGroup: "Neck"},
		{MuscleGroup: "Trunk"},
		{MuscleGroup: "Lower Back"},
		{MuscleGroup: "Upper Back"},
		{MuscleGroup: "Core"},
		{MuscleGroup: "Legs"},
	}

	for _, muscleGroup := range muscleGroup {
		if err := dbConn.Create(&muscleGroup).Error; err != nil {
			log.Fatal(err)
		}
	}

	exercises := []*domain.Exercise{
		{
			Name:        "Push-up",
			Description: "A push-up is a common calisthenics exercise beginning from the prone position.",
			MuscleGroup: []*domain.MuscleGroup{muscleGroup[0], muscleGroup[1], muscleGroup[2], muscleGroup[3]},
		},
		{
			Name:        "Pull-up",
			Description: "A pull-up is an upper-body strength exercise.",
			MuscleGroup: []*domain.MuscleGroup{muscleGroup[1], muscleGroup[2], muscleGroup[3], muscleGroup[11]},
		},
		{
			Name:        "Squat",
			Description: "The squat is a lower body exercise.",
			MuscleGroup: []*domain.MuscleGroup{muscleGroup[4], muscleGroup[5], muscleGroup[6], muscleGroup[16]},
		},
		{
			Name:        "Deadlift",
			Description: "The deadlift is a weight training exercise in which a loaded barbell or bar is lifted off the ground to the level of the hips, then lowered to the ground.",
			MuscleGroup: []*domain.MuscleGroup{muscleGroup[4], muscleGroup[5], muscleGroup[6], muscleGroup[16]},
		},
		{
			Name:        "Bench Press",
			Description: "The bench press is an upper-body strength-training exercise that consists of pressing a weight upwards from a supine position.",
			MuscleGroup: []*domain.MuscleGroup{muscleGroup[0], muscleGroup[1], muscleGroup[2], muscleGroup[3]},
		},
		{
			Name:        "Dumbbell Curl",
			Description: "The dumbbell curl is a weight training exercise performed by lifting a dumbbell with one arm.",
			MuscleGroup: []*domain.MuscleGroup{muscleGroup[2]},
		},
		{
			Name:        "Tricep Extension",
			Description: "The tricep extension is a strength training exercise used to work the triceps muscles.",
			MuscleGroup: []*domain.MuscleGroup{muscleGroup[3]},
		},
		{
			Name:        "Leg Press",
			Description: "The leg press is a weight training exercise in which the individual pushes a weight or resistance away from them using their legs.",
			MuscleGroup: []*domain.MuscleGroup{muscleGroup[4], muscleGroup[5], muscleGroup[6], muscleGroup[16]},
		},
		{
			Name:        "Leg Curl",
			Description: "The leg curl is a resistance weight training exercise that targets the hamstrings.",
			MuscleGroup: []*domain.MuscleGroup{muscleGroup[5]},
		},
		{
			Name:        "Leg Extension",
			Description: "The leg extension is a resistance weight training exercise that targets the quadriceps.",
			MuscleGroup: []*domain.MuscleGroup{muscleGroup[4]},
		},
		{
			Name:        "Calf Raise",
			Description: "The calf raise is a strength training exercise which targets the calf muscles.",
			MuscleGroup: []*domain.MuscleGroup{muscleGroup[6]},
		},
		{
			Name:        "Crunch",
			Description: "The crunch is a strength training exercise for the abdominal muscles.",
			MuscleGroup: []*domain.MuscleGroup{muscleGroup[7]},
		},
		{
			Name:        "Russian Twist",
			Description: "The Russian twist is a type of exercise that is used to work the abdominal muscles by performing a twisting motion on the abdomen.",
			MuscleGroup: []*domain.MuscleGroup{muscleGroup[7], muscleGroup[12]},
		},
		{
			Name:        "Plank",
			Description: "The plank is an isometric core strength exercise that involves maintaining a position similar to a push-up for the maximum possible time.",
			MuscleGroup: []*domain.MuscleGroup{muscleGroup[7], muscleGroup[16]},
		},
		{
			Name:        "Glute Bridge",
			Description: "The glute bridge is a lower body exercise that targets the glutes.",
			MuscleGroup: []*domain.MuscleGroup{muscleGroup[8]},
		},
		{
			Name:        "Forearm Curl",
			Description: "The forearm curl is a weight training exercise that targets the forearm muscles.",
			MuscleGroup: []*domain.MuscleGroup{muscleGroup[9]},
		},
		{
			Name:        "Shrug",
			Description: "The shrug is an exercise that is used to strengthen the muscles in the shoulder region.",
			MuscleGroup: []*domain.MuscleGroup{muscleGroup[10]},
		},
		{
			Name:        "Lat Pulldown",
			Description: "The lat pulldown is a strength training exercise that targets the latissimus dorsi muscle.",
			MuscleGroup: []*domain.MuscleGroup{muscleGroup[11]},
		},
		{
			Name:        "Oblique Crunch",
			Description: "The oblique crunch is a strength training exercise that targets the oblique muscles.",
			MuscleGroup: []*domain.MuscleGroup{muscleGroup[12]},
		},
		{
			Name:        "Neck Curl",
			Description: "The neck curl is a strength training exercise that targets the neck muscles.",
			MuscleGroup: []*domain.MuscleGroup{muscleGroup[13]},
		},
		{
			Name:        "Trunk Rotation",
			Description: "The trunk rotation is a strength training exercise that targets the oblique muscles.",
			MuscleGroup: []*domain.MuscleGroup{muscleGroup[14]},
		},
		{
			Name:        "Lower Back Extension",
			Description: "The lower back extension is a strength training exercise that targets the lower back muscles.",
			MuscleGroup: []*domain.MuscleGroup{muscleGroup[15]},
		},
	}

	for _, exercise := range exercises {
		if err := dbConn.Create(&exercise).Error; err != nil {
			log.Fatal(err)
		}
	}
}