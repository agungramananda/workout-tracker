import { useState, useEffect } from "react";
import {
  Box,
  Heading,
  Text,
  VStack,
  List,
  ListItem,
  Divider,
} from "@chakra-ui/react";
import moment from "moment";

const History = () => {
  const [workouts, setWorkouts] = useState([]);

  useEffect(() => {
    const token = localStorage.getItem("token");
    fetch("http://localhost:8080/workouts", {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    })
      .then((response) => response.json())
      .then((data) => {
        const completedWorkouts = data.data.filter(
          (workout) => workout.isCompleted
        );
        setWorkouts(completedWorkouts);
      })
      .catch((error) => {
        console.error("Error fetching workout history:", error);
      });
  }, []);

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
            Workout History
          </Heading>
          <Text fontSize="lg" color="gray.600" textAlign="center">
            Review your past workouts
          </Text>
          <Divider />
          <Box p={4} bg="gray.50" borderRadius="md" boxShadow="md">
            <Heading as="h2" size="md" mb={4} color="teal.500">
              Past Workouts
            </Heading>
            <List spacing={2}>
              {workouts.map((workout) => (
                <ListItem key={workout.id}>
                  {workout.name} - {moment(workout.date).format("MMMM Do YYYY")}{" "}
                  at {moment(workout.time, "HH:mm").format("h:mm a")}
                </ListItem>
              ))}
            </List>
          </Box>
        </VStack>
      </Box>
    </Box>
  );
};

export default History;
