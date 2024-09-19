import { Box, VStack, Text, Button, Heading } from "@chakra-ui/react";
import { useNavigate } from "react-router-dom";
import homePageImage from "../../assets/img/Home_page.jpg";

const Welcome = () => {
  const navigate = useNavigate();

  const handleGetStarted = () => {
    navigate("/login");
  };

  return (
    <Box
      bgImage={`url(${homePageImage})`}
      bgSize="cover"
      bgPosition="center"
      minH="100vh"
      display="flex"
      alignItems="center"
      justifyContent="center"
      position="relative"
      _before={{
        content: '""',
        position: "absolute",
        top: 0,
        left: 0,
        right: 0,
        bottom: 0,
        bgGradient: "linear(to-r, rgba(0, 0, 0, 0.7), rgba(0, 0, 0, 0.1))",
        zIndex: 0,
      }}
    >
      <VStack
        spacing={8}
        p={8}
        bg="rgba(255, 255, 255, 0.9)"
        borderRadius="lg"
        zIndex={1}
        boxShadow="xl"
      >
        <Heading
          as="h1"
          size="2xl"
          textAlign="center"
          color="teal.500"
          fontWeight="bold"
        >
          Welcome to Workout Tracker
        </Heading>
        <Text
          fontSize="xl"
          textAlign="center"
          color="gray.700"
          fontWeight="semibold"
        >
          Track your workouts, monitor your progress, and achieve your fitness
          goals with our easy-to-use workout tracker.
        </Text>

        <Button
          colorScheme="teal"
          size="lg"
          transition="background-color 0.3s, transform 0.3s"
          _hover={{ bg: "teal.600", transform: "scale(1.1)" }}
          onClick={handleGetStarted}
          fontWeight="bold"
        >
          Get Started
        </Button>
      </VStack>
    </Box>
  );
};

export default Welcome;
