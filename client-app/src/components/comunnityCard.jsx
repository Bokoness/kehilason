/* eslint-disable */
import {
  Avatar,
  Box,
  Button,
  Card,
  CardActions,
  CardContent,
  CardHeader,
  Grid,
  IconButton,
  List,
  ListItem,
  ListItemAvatar,
  ListItemText,
  Typography,
} from "@mui/material";
import { red } from "@mui/material/colors";
import ArrowForwardIosIcon from "@mui/icons-material/ArrowForwardIos";
import { Folder } from "@mui/icons-material";
import PersonIcon from "@mui/icons-material/Person";
import { styled } from "@mui/material/styles";
import React, { useState } from "react";
import InsertPhotoOutlinedIcon from "@mui/icons-material/InsertPhotoOutlined";
const IconAndTextLine = ({ text, Icon }) => {
  return (
    <ListItem>
      <ListItemAvatar>
        <Icon fontSize={"small"} />
      </ListItemAvatar>
      <ListItemText primary={text} />
    </ListItem>
  );
  // PersonIcon
};
const Demo = styled("div")(({ theme }) => ({
  backgroundColor: theme.palette.background.paper,
}));
function generate(element) {
  return [0, 1, 2].map((value) =>
    React.cloneElement(element, {
      key: value,
    })
  );
}
const MyList = () => {
  const [dense, setDense] = useState(false);
  const [secondary, setSecondary] = useState(false);
  return (
    <Grid item xs={12} md={6}>
      <Typography sx={{ mt: 4, mb: 2 }} variant="h6" component="div">
        Avatar with text
      </Typography>
      <Demo>
        <List dense={dense}>
          {generate(
            <ListItem>
              <ListItemAvatar>
                <Folder fontSize={"small"} />
              </ListItemAvatar>
              <ListItemText
                primary="Single-line item"
                secondary={secondary ? "Secondary text" : null}
              />
            </ListItem>
          )}
        </List>
      </Demo>
    </Grid>
  );
};
const ComunnityCard = () => {
  return (
    <Card sx={{ minWidth: 275 }}>
      <CardHeader
        dir="rtl"
        style={{ marginLeft: "12px" }}
        avatar={
          <Avatar aria-label="recipe">
            <InsertPhotoOutlinedIcon />
          </Avatar>
        }
        title="שם הקהילה"
      ></CardHeader>
      <MyList />
      <CardContent>
        {/* <Typography sx={{ fontSize: 14 }} color="text.secondary" gutterBottom>
          Word of the Day
        </Typography> */}
        {/* <IconAndTextLine
          text={"bla bla"}
          icon={<PersonIcon />}
        ></IconAndTextLine> */}
        <Typography variant="h5" component="div">
          be nev o lent
        </Typography>
        <Typography sx={{ mb: 1.5 }} color="text.secondary">
          adjective
        </Typography>
        <Typography variant="body2">
          well meaning and kindly.
          <br />
          {'"a benevolent smile"'}
        </Typography>
      </CardContent>
      <CardActions>
        <Button variant="contained" size="small">
          to community <ArrowForwardIosIcon></ArrowForwardIosIcon>
        </Button>
      </CardActions>
    </Card>
  );
};
export default ComunnityCard;
