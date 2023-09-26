const express = require("express");
const app = express();
const cors = require("cors");
const jwt = require("jsonwebtoken");
const port = 3000;

app.use(express.json());
app.use(cors());
app.use(express.urlencoded({ extended: true }));

app.get("/api/auth/signup", (req, res) => {});
app.get("/api/auth/login", (req, res) => {
  res.cookie("token", "some token");
  res.json({ success: true });
});

app.listen(port, () => {
  console.log("running");
});
