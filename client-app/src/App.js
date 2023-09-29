import "./App.css";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import RegisterPage from "./pages/Register.jsx";
import { Container } from "@mui/material";
import Login from "./pages/Login.jsx";

function App() {
  const router = createBrowserRouter([
    {
      path: "/signup",
      element: <RegisterPage />,
    },
    {
      path: "/",
      element: <Login />,
    },
  ]);
  return (
    <div className="App">
      <Container style={{ paddingTop: 5 }}>
        <RouterProvider router={router} />
      </Container>
    </div>
  );
}

export default App;
