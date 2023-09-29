import "./App.css";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import RegisterPage from "./pages/Register.jsx";
import { Container } from "@mui/material";
import Login from "./pages/Login.jsx";
import { useState } from "react";
import Dashboard from "./pages/Dashboard.jsx";
import ProtectedRout from "./components/ProtectedRout.jsx";
import UnProtectedRout from "./components/UnProtectedRout.jsx";

function App() {
  const [authorized, setAuthorized] = useState(false);
  const [userId, setUserId] = useState();

  useEffect(() => {
    const authorization = async () => {
      if (authorized) return;
      const user = await auth.autorized();
      if (user) {
        setAuthorized(true);
        setUserId(user._id);
      }
    };
    authorization();
  }, [authorized]);

  return (
    <div className="App">
      <Container style={{ paddingTop: 10 }}>
        <BrowserRouter>
          <Routes>
            <Route
              path="/"
              element={<UnProtectedRout authorized={authorized} />}
            >
              <Route
                path="/"
                element={<Login setAuthorized={setAuthorized} />}
              />
              <Route
                path="/register"
                element={<RegisterPage setAuthorized={setAuthorized} />}
              />
            </Route>
            <Route
              path="/home"
              element={<ProtectedRout authorized={authorized} />}
            >
              <Route
                path="/home"
                element={
                  <Dashboard
                    authorized={authorized}
                    //  userId={userId}
                  />
                }
              />
            </Route>
          </Routes>
        </BrowserRouter>
      </Container>
    </div>
  );
}

export default App;
