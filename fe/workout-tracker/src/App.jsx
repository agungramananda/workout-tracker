import {
  BrowserRouter as Router,
  Route,
  Routes,
  useLocation,
} from "react-router-dom";
import Register from "./components/Register/Register";
import Login from "./components/Login/Login";
import { ChakraProvider } from "@chakra-ui/react";
import Welcome from "./components/Welcome/Welcome.Jsx";
import Navbar from "./components/Navbar/Navbar";
import Home from "./components/Home/Home.Jsx";
import NewWorkout from "./components/NewWorkout/NewWorkout";
import StartWorkout from "./components/StartWorkout/StartWorkout";
import History from "./components/History/History";
import EditWorkout from "./components/EditWorkout/EditWorkout";
import NotFound from "./components/NotFound/NotFound";

function App() {
  return (
    <ChakraProvider>
      <Router>
        <ConditionalNavbar />
        <Routes>
          <Route path="/register" element={<Register />} />
          <Route path="/login" element={<Login />} />
          <Route path="/home" element={<Home />} />
          <Route path="/new" element={<NewWorkout />} />
          <Route path="/edit/:id" element={<EditWorkout />} />
          <Route path="/start-workout" element={<StartWorkout />} />
          <Route path="/history" element={<History />} />
          <Route path="/" element={<Welcome />} />
          <Route path="*" element={<NotFound />} />
        </Routes>
      </Router>
    </ChakraProvider>
  );
}

function ConditionalNavbar() {
  const location = useLocation();
  const NavbarRoutes = ["/home", "/new", "/edit", "/start-workout", "/history"];
  return NavbarRoutes.includes(location.pathname) ? <Navbar /> : null;
}

export default App;
