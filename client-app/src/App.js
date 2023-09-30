import "./App.css"
import { createBrowserRouter, RouterProvider } from "react-router-dom"
import { Container } from "@mui/material"
import RegisterPage from "./views/Register"

function App() {
  const router = createBrowserRouter([
    {
      path: "/",
      element: <RegisterPage />,
    },
  ])
  return (
    <div className="App">
      <Container style={{ paddingTop: 5 }}>
        <RouterProvider router={router} />
      </Container>
    </div>
  )
}

export default App
