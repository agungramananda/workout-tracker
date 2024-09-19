import { Box, Flex, Heading, Button } from "@chakra-ui/react";
import { useNavigate } from "react-router-dom";

const Navbar = () => {
  const navigate = useNavigate();

  const handleLogout = () => {
    localStorage.removeItem("token");
    navigate("/");
  };

  return (
    <Box bg="teal.500" p={4} color="white">
      <Flex align="center" justifyContent="space-between">
        <Heading as="h1" size="lg">
          Workout Tracker
        </Heading>
        <Flex ml="auto">
          <>
            <Button
              mx="4"
              colorScheme="teal"
              variant="outline"
              color="white"
              onClick={() => navigate("/home")}
            >
              Home
            </Button>
            <Button
              colorScheme="teal"
              variant="outline"
              color="white"
              onClick={handleLogout}
            >
              Logout
            </Button>
          </>
        </Flex>
      </Flex>
    </Box>
  );
};

export default Navbar;
