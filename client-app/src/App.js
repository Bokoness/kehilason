import "./App.css"
import { createBrowserRouter, RouterProvider } from "react-router-dom"
import { Container } from "@mui/material"
import RegisterPage from "./views/Register"
import Theme from "./configs/Theme"

function App() {
  const router = createBrowserRouter([
    {
      path: "/",
      element: <RegisterPage />,
    },
  ])
  return (
    <div className="App">
      <Theme>

      <Container style={{ paddingTop: 5 }}>
        <RouterProvider router={router} />
      </Container>
      </Theme>
    </div>
  )
}

export default App
