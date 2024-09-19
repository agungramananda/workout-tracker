import { useState, useEffect } from "react";
import {
  Box,
  Heading,
  Text,
  VStack,
  Button,
  List,
  ListItem,
  Divider,
  HStack,
  IconButton,
} from "@chakra-ui/react";
import { EditIcon, DeleteIcon, AddIcon, CheckIcon } from "@chakra-ui/icons";
import moment from "moment";
import "react-big-calendar/lib/css/react-big-calendar.css";
import { useNavigate } from "react-router-dom";

const Home = () => {
  const [workouts, setWorkouts] = useState([]);
  const navigate = useNavigate();

  useEffect(() => {
    const token = localStorage.getItem("token");
    fetch("http://localhost:8080/workouts", {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    })
      .then((response) => response.json())
      .then((data) => {
        const formattedWorkouts = data.data.map((workout) => ({
          id: workout.id,
          title: workout.name,
          date: workout.date,
          time: workout.time,
          isCompleted: workout.isCompleted,
        }));
        setWorkouts(formattedWorkouts);
      })
      .catch((error) => {
        console.error("Error fetching workouts:", error);
      });
  }, []);

  const handleEditWorkout = (workout) => {
    navigate(`/edit/${workout.id}`);
  };

  const handleDeleteWorkout = (workoutId) => {
    const token = localStorage.getItem("token");
    fetch(`http://localhost:8080/workouts/${workoutId}`, {
      method: "DELETE",
      headers: {
        Authorization: `Bearer ${token}`,
      },
    })
      .then((response) => {
        if (response.ok) {
          setWorkouts(workouts.filter((workout) => workout.id !== workoutId));
        } else {
          console.error("Error deleting workout");
        }
      })
      .catch((error) => {
        console.error("Error deleting workout:", error);
      });
  };

  const handleCompleteWorkout = (workoutId) => {
    const token = localStorage.getItem("token");
    fetch(`http://localhost:8080/workouts/${workoutId}/complete`, {
      method: "PATCH",
      headers: {
        Authorization: `Bearer ${token}`,
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ isCompleted: true }),
    })
      .then((response) => {
        if (response.ok) {
          setWorkouts(
            workouts.map((workout) =>
              workout.id === workoutId
                ? { ...workout, isCompleted: true }
                : workout
            )
          );
        } else {
          console.error("Error completing workout");
        }
      })
      .catch((error) => {
        console.error("Error completing workout:", error);
      });
  };

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
        <VStack spacing={4} align="stretch">
          <Heading as="h1" size="xl" textAlign="center" color="teal.500">
            Welcome to Workout Tracker
          </Heading>
          <Text fontSize="lg" color="gray.600" textAlign="center">
            Track your workouts and stay fit!
          </Text>
          <Box p={4} bg="gray.50" borderRadius="md" boxShadow="md">
            <Heading as="h2" size="md" mb={4} color="teal.500">
              Your Workout Summary
            </Heading>
            <List spacing={2}>
              <ListItem>Total Workouts: {workouts.length}</ListItem>
            </List>
          </Box>
          <Box p={4} bg="gray.50" borderRadius="md" boxShadow="md">
            <Heading as="h2" size="md" mb={4} color="teal.500">
              Upcoming Workouts
            </Heading>
            <List spacing={2}>
              {workouts.map((workout) => (
                <ListItem key={workout.id}>
                  <HStack justifyContent="space-between">
                    <Text>
                      {workout.title} -{" "}
                      {moment(workout.date).format("MMMM Do YYYY, h:mm a")}
                    </Text>
                    <HStack>
                      <IconButton
                        size="sm"
                        colorScheme="blue"
                        icon={<EditIcon />}
                        onClick={() => handleEditWorkout(workout)}
                      />
                      <IconButton
                        size="sm"
                        colorScheme="red"
                        icon={<DeleteIcon />}
                        onClick={() => handleDeleteWorkout(workout.id)}
                      />
                      {!workout.isCompleted && (
                        <IconButton
                          size="sm"
                          colorScheme="green"
                          icon={<CheckIcon />}
                          onClick={() => handleCompleteWorkout(workout.id)}
                        />
                      )}
                    </HStack>
                  </HStack>
                </ListItem>
              ))}
            </List>
          </Box>
          <Divider />
          <Box>
            <Button
              colorScheme="teal"
              size="lg"
              width="full"
              leftIcon={<AddIcon />}
              onClick={() => navigate("/new")}
            >
              Add New Workout
            </Button>
            <Button
              mt={4}
              colorScheme="teal"
              size="lg"
              width="full"
              onClick={() => navigate("/history")}
            >
              View Workout History
            </Button>
          </Box>
          <Divider />
        </VStack>
      </Box>
    </Box>
  );
};

export default Home;
