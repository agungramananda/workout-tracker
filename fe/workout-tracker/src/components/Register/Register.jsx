import { useState } from "react";
import {
  Box,
  Heading,
  Input,
  Button,
  Text,
  VStack,
  Link,
} from "@chakra-ui/react";
import { useNavigate } from "react-router-dom";

const Register = () => {
  const [username, setUsername] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [fullName, setFullName] = useState("");
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();

  const handleRegister = async () => {
    setLoading(true);
    try {
      const response = await fetch("http://localhost:8080/register", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ username, email, password, fullName }),
      });

      if (!response.ok) {
        if (response.status === 400) {
          throw new Error("Invalid request");
        }
        throw new Error("Username or email already exists");
      }

      const res = await response.json();
      const token = res?.data?.[0]?.access_token;

      localStorage.setItem("token", token);

      navigate("/home");
    } catch (err) {
      console.error(err);
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  return (
    <Box
      minH="100vh"
      display="flex"
      alignItems="center"
      justifyContent="center"
      bg="gray.100"
      p={4}
    >
      <Box
        p={8}
        bg="white"
        boxShadow="lg"
        borderRadius="lg"
        maxW="sm"
        width="full"
      >
        <VStack spacing={4} align="stretch">
          <Heading as="h2" size="lg" textAlign="center" color="teal.500">
            Workout Tracker
          </Heading>
          <Text fontSize="md" color="gray.600" textAlign="center">
            Create your account
          </Text>
          {error && (
            <Text fontSize="sm" color="red.500" textAlign="center">
              {error}
            </Text>
          )}
          <Input
            placeholder="Full Name"
            variant="filled"
            value={fullName}
            onChange={(e) => setFullName(e.target.value)}
          />
          <Input
            placeholder="Username"
            variant="filled"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
          />
          <Input
            placeholder="Email"
            type="email"
            variant="filled"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          />
          <Input
            placeholder="Password"
            type="password"
            variant="filled"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
          <Button
            colorScheme="teal"
            size="lg"
            width="full"
            onClick={handleRegister}
            isLoading={loading}
            disabled={loading}
          >
            Register
          </Button>
          <Text fontSize="sm" color="gray.600" textAlign="center">
            Already have an account?{" "}
            <Link color="teal.500" onClick={() => navigate("/login")}>
              Login
            </Link>
          </Text>
        </VStack>
      </Box>
    </Box>
  );
};

export default Register;
