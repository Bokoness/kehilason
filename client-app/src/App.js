import './App.css';
import {createBrowserRouter, RouterProvider} from "react-router-dom";
import RegisterPage from "./views/Register"
function App() {
  const router = createBrowserRouter([
    {
      path: "/",
      element: <div>Main page</div>
    },
    {
      path: "/register",
      element: <RegisterPage/>
    }
  ])
  return (
    <div className="App">
      <RouterProvider router={router}/>
    </div>
  );
}

export default App;
