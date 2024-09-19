import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import {
  Box,
  Button,
  FormControl,
  FormLabel,
  Input,
  NumberInput,
  NumberInputField,
  VStack,
  Textarea,
  useToast,
  Select,
  IconButton,
} from "@chakra-ui/react";
import { AddIcon, DeleteIcon } from "@chakra-ui/icons";

const NewWorkout = () => {
  const [workoutName, setWorkoutName] = useState("");
  const [description, setDescription] = useState("");
  const [date, setDate] = useState("");
  const [time, setTime] = useState("");
  const [restBetweenExercises, setRestBetweenExercises] = useState(60);
  const [isCompleted, setIsCompleted] = useState(false);
  const [comments, setComments] = useState([{ user_id: 1, comment: "" }]);
  const [exercises, setExercises] = useState([
    { name: "", sets: "", reps: "" },
  ]);
  const [exerciseOptions, setExerciseOptions] = useState([]);

  const navigate = useNavigate();
  const toast = useToast();

  useEffect(() => {
    const token = localStorage.getItem("token");
    fetch("http://localhost:8080/exercises", {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    })
      .then((response) => response.json())
      .then((data) => {
        setExerciseOptions(data.data);
      })
      .catch((error) => {
        console.error("Error fetching exercises:", error);
      });
  }, []);

  const handleWorkoutNameChange = (e) => {
    setWorkoutName(e.target.value);
  };

  const handleDescriptionChange = (e) => {
    setDescription(e.target.value);
  };

  const handleDateChange = (e) => {
    setDate(e.target.value);
  };

  const handleTimeChange = (e) => {
    setTime(e.target.value);
  };

  const handleRestBetweenExercisesChange = (e) => {
    setRestBetweenExercises(e.target.value);
  };

  const handleExerciseChange = (index, e) => {
    const { name, value } = e.target;
    const newExercises = [...exercises];
    newExercises[index][name] = value;
    setExercises(newExercises);
  };

  const addExercise = () => {
    setExercises([...exercises, { name: "", sets: "", reps: "" }]);
  };

  const deleteExercise = (index) => {
    const newExercises = exercises.filter((_, i) => i !== index);
    setExercises(newExercises);
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    const token = localStorage.getItem("token");

    fetch("http://localhost:8080/workouts", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify({
        name: workoutName,
        description,
        date,
        time,
        rest_between_exercises: parseInt(restBetweenExercises, 10),
        is_completed: isCompleted,
        exercises_plan: exercises.map((exercise, index) => ({
          exercise_id: exerciseOptions.find(
            (option) => option.name === exercise.name
          )?.id,
          sets: parseInt(exercise.sets, 10),
          reps: parseInt(exercise.reps, 10),
          weight: 0,
          rest_time: 60,
          order: index + 1,
          is_completed: false,
        })),
        comments: comments.map((comment) => ({
          user_id: comment.user_id,
          comment: comment.comment,
        })),
      }),
    }).then((response) => {
      if (!response.ok) {
        toast({
          title: "Error saving workout.",
          description: "There was an error saving your workout.",
          status: "error",
          duration: 5000,
          isClosable: true,
        });
        return;
      }
      toast({
        title: "Workout saved.",
        description: "Your workout has been successfully saved.",
        status: "success",
        duration: 5000,
        isClosable: true,
      });
      navigate(-1); // Go back after success
      return response.json();
    });
  };

  return (
    <Box p={4}>
      <VStack spacing={4} align="stretch">
        <FormControl>
          <FormLabel>Workout Name</FormLabel>
          <Input value={workoutName} onChange={handleWorkoutNameChange} />
        </FormControl>
        <FormControl>
          <FormLabel>Description</FormLabel>
          <Textarea value={description} onChange={handleDescriptionChange} />
        </FormControl>
        <FormControl>
          <FormLabel>Date</FormLabel>
          <Input type="date" value={date} onChange={handleDateChange} />
        </FormControl>
        <FormControl>
          <FormLabel>Time</FormLabel>
          <Input type="time" value={time} onChange={handleTimeChange} />
        </FormControl>
        <FormControl>
          <FormLabel>Rest Between Exercises (seconds)</FormLabel>
          <NumberInput>
            <NumberInputField
              value={restBetweenExercises}
              onChange={handleRestBetweenExercisesChange}
            />
          </NumberInput>
        </FormControl>
        {exercises.map((exercise, index) => (
          <Box key={index} borderWidth="1px" borderRadius="lg" p={4}>
            <FormControl>
              <FormLabel>Exercise Name</FormLabel>
              <Select
                name="name"
                value={exercise.name}
                onChange={(e) => handleExerciseChange(index, e)}
              >
                <option value="">-</option>
                {exerciseOptions.map((option, idx) => (
                  <option key={idx} value={option.name}>
                    {option.name}
                  </option>
                ))}
              </Select>
            </FormControl>
            <FormControl>
              <FormLabel>Sets</FormLabel>
              <NumberInput>
                <NumberInputField
                  name="sets"
                  value={exercise.sets}
                  onChange={(e) => handleExerciseChange(index, e)}
                />
              </NumberInput>
            </FormControl>
            <FormControl>
              <FormLabel>Reps</FormLabel>
              <NumberInput>
                <NumberInputField
                  name="reps"
                  value={exercise.reps}
                  onChange={(e) => handleExerciseChange(index, e)}
                />
              </NumberInput>
            </FormControl>
            <IconButton
              my="4"
              colorScheme="red"
              icon={<DeleteIcon />}
              onClick={() => deleteExercise(index)}
            />
          </Box>
        ))}
        <Button leftIcon={<AddIcon />} colorScheme="teal" onClick={addExercise}>
          Add Exercise
        </Button>
        <Button colorScheme="teal" onClick={handleSubmit}>
          Save Workout
        </Button>
        <Button onClick={() => navigate(-1)}>Back</Button>
      </VStack>
    </Box>
  );
};

export default NewWorkout;
