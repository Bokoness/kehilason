import * as React from "react";
import Box from "@mui/material/Box";
import TextField from "@mui/material/TextField";
import api from "../api/auth";
import { Button, Grid } from "@mui/material";

export default function BasicTextFields() {
  const [fullName, setFullName] = React.useState("");
  const [email, setEmail] = React.useState("");
  const [password, setPassword] = React.useState("");

  const handleRegister = () => {
    api.register(fullName, email, password);
  };

  return (
    <Box component="form" noValidate autoComplete="on">
      <Grid container rowSpacing={1}>
        <Grid item xs={12}>
          <TextField
            label="Full Name"
            variant="outlined"
            value={fullName}
            onChange={(e) => setFullName(e.target.value)}
          />
        </Grid>
        <Grid item xs={12}>
          <TextField
            label="Email"
            variant="outlined"
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          />
        </Grid>

        <Grid item xs={12}>
          <TextField
            label="Password"
            variant="outlined"
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />
        </Grid>
      </Grid>
      <Button onClick={handleRegister}>Sign up</Button>
    </Box>
  );
}
