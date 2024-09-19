import { useState, useEffect } from "react";
import {
  Box,
  Heading,
  Text,
  VStack,
  Button,
  FormControl,
  FormLabel,
  Input,
} from "@chakra-ui/react";
import { useParams, useNavigate } from "react-router-dom";
import moment from "moment";

const StartWorkout = () => {
  const { workoutId } = useParams();
  const [workout, setWorkout] = useState(null);
  const [duration, setDuration] = useState("");
  const navigate = useNavigate();

  useEffect(() => {
    const token = localStorage.getItem("token");
    fetch(`http://localhost:8080/workouts/${workoutId}`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    })
      .then((response) => response.json())
      .then((data) => {
        setWorkout(data.data);
      })
      .catch((error) => {
        console.error("Error fetching workout:", error);
      });
  }, [workoutId]);

  const handleCompleteWorkout = () => {
    const token = localStorage.getItem("token");
    fetch(`http://localhost:8080/workouts/${workoutId}`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify({
        workout,
      }),
    })
      .then((response) => response.json())
      .then(() => {
        navigate("/history");
      })
      .catch((error) => {
        console.error("Error completing workout:", error);
      });
  };

  if (!workout) {
    return <Text>Loading...</Text>;
  }

  return (
    <Box minH="100vh" p={4} bg="gray.100">
      <Box
        maxW="xl"
        mx="auto"
        p={6}
        bg="white"
        boxShadow="lg"
        borderRadius="lg"
      >
        <VStack spacing={4}>
          <Heading as="h1" size="xl" color="teal.500" textAlign="center">
            Do Workout
          </Heading>
          <Box p={4} bg="gray.50" borderRadius="md" boxShadow="md" width="100%">
            <Text fontSize="lg" fontWeight="bold" color="teal.500">
              {workout.name}
            </Text>
            <Text>Date: {moment(workout.date).format("MMMM Do YYYY")}</Text>
            <Text>Time: {moment(workout.time, "HH:mm").format("h:mm A")}</Text>
          </Box>
          <FormControl mt={4}>
            <FormLabel>Duration (minutes)</FormLabel>
            <Input
              type="number"
              placeholder="Enter workout duration"
              value={duration}
              onChange={(e) => setDuration(e.target.value)}
              isRequired
            />
          </FormControl>
          <Button
            colorScheme="teal"
            size="lg"
            width="full"
            onClick={handleCompleteWorkout}
          >
            Complete Workout
          </Button>
        </VStack>
      </Box>
    </Box>
  );
};

export default StartWorkout;
