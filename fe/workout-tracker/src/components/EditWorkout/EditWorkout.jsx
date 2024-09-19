import { useState, useEffect } from "react";
import { useNavigate, useParams } from "react-router-dom";
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
  Select,
  useToast,
  IconButton,
} from "@chakra-ui/react";
import { AddIcon, DeleteIcon } from "@chakra-ui/icons";

const EditWorkout = () => {
  const { id } = useParams();
  const [workout, setWorkout] = useState({});
  const [exerciseOptions, setExerciseOptions] = useState([]);
  const navigate = useNavigate();
  const toast = useToast();

  useEffect(() => {
    const token = localStorage.getItem("token");

    fetch(`http://localhost:8080/workouts/${id}`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    })
      .then((response) => response.json())
      .then((data) => {
        const workoutData = data.data[0];

        const workoutTime = workoutData.time
          .split(":")
          .slice(0, 2)
          .join(":")
          .split("T")[1];
        const workoutDate = workoutData.date.split("T")[0];

        setWorkout({
          ...workoutData,
          time: workoutTime,
          date: workoutDate,
        });
        console.log(workoutData);
      })
      .catch((error) => {
        console.error("Error fetching workout:", error);
      });

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
  }, [id]);

  const handleChange = (e) => {
    setWorkout({
      ...workout,
      [e.target.name]: e.target.value,
    });
  };

  const handleExerciseChange = (index, e) => {
    const { name, value } = e.target;
    const newExercises = [...workout.exercises_plan];
    newExercises[index][name] = value;
    setWorkout({ ...workout, exercises_plan: newExercises });
  };

  const handleAddExercise = () => {
    setWorkout({
      ...workout,
      exercises_plan: [
        ...workout.exercises_plan,
        { name: "", sets: "", reps: "" },
      ],
    });
  };

  const handleDeleteExercise = (index) => {
    const newExercises = workout.exercises_plan.filter((_, i) => i !== index);
    setWorkout({ ...workout, exercises_plan: newExercises });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    const token = localStorage.getItem("token");

    fetch(`http://localhost:8080/workouts/${id}`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify(workout),
    })
      .then((response) => {
        if (response.ok) {
          toast({
            title: "Workout updated.",
            status: "success",
            duration: 3000,
            isClosable: true,
          });
          navigate("/");
        } else {
          toast({
            title: "Failed to update workout.",
            status: "error",
            duration: 3000,
            isClosable: true,
          });
        }
      })
      .catch((error) => {
        console.error("Error updating workout:", error);
        toast({
          title: "Failed to update workout.",
          status: "error",
          duration: 3000,
          isClosable: true,
        });
      });
  };

  const handleCancel = () => {
    navigate("/home");
  };

  return (
    <Box p={4}>
      <VStack spacing={4} align="stretch">
        <FormControl>
          <FormLabel>Workout Name</FormLabel>
          <Input name="name" value={workout.name} onChange={handleChange} />
        </FormControl>
        <FormControl>
          <FormLabel>Description</FormLabel>
          <Textarea
            name="description"
            value={workout.description}
            onChange={handleChange}
          />
        </FormControl>
        <FormControl>
          <FormLabel>Date</FormLabel>
          <Input
            type="date"
            name="date"
            value={workout.date}
            onChange={handleChange}
          />
        </FormControl>
        <FormControl>
          <FormLabel>Time</FormLabel>
          <Input
            type="time"
            name="time"
            value={workout.time}
            onChange={handleChange}
          />
        </FormControl>
        {workout.exercises_plan?.map((exercise, index) => (
          <Box key={index} borderWidth="1px" borderRadius="lg" p={4}>
            <FormControl>
              <FormLabel>Exercise Name</FormLabel>
              <Select
                name="name"
                value={exercise.name}
                onChange={(e) => handleExerciseChange(index, e)}
              >
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
              aria-label="Delete exercise"
              icon={<DeleteIcon />}
              onClick={() => handleDeleteExercise(index)}
              colorScheme="red"
              mt={2}
            />
          </Box>
        ))}
        <Button
          leftIcon={<AddIcon />}
          colorScheme="teal"
          onClick={handleAddExercise}
        >
          Add Exercise
        </Button>
        <Button colorScheme="teal" onClick={handleSubmit}>
          Save Changes
        </Button>
        <Button colorScheme="red" onClick={handleCancel}>
          Cancel
        </Button>
      </VStack>
    </Box>
  );
};

export default EditWorkout;
