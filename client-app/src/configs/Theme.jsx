// import {
//   createTheme,
//   //  ThemeProvider
// } from "@mui/material/styles";
// import { heIL } from "@mui/material/locale";

// const theme = createTheme(
//   {
//     palette: {
//       primary: { main: "#1976d2" },
//     },
//   },
//   heIL
// );
const Provider = ({ children }) => {
  return children;
  // <ThemeProvider theme={theme}>
  //   <>{children}</>
  // </ThemeProvider>
};

export default Provider;
