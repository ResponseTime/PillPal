const express = require("express");
const app = express();
const cors = require("cors");
const jwt = require("jsonwebtoken");
const db = require("./database.js");
const port = 3000;

app.use(express.json());
app.use(cors());
app.use(express.urlencoded({ extended: true }));

app.post("/api/auth/signup", async (req, res) => {
  const { email, password } = req.body;
  try {
    let collection = await db.collection("login");
    let user = await collection.findOne({ email });
    if (user) {
      res.status(400).json({ message: "user exists login" });
    }
    const salt = await bcrypt.genSalt();
    const encrytedPass = await bcrypt.hash(password, salt);
    let newIns = await collection.insertOne({
      email: email,
      password: encrytedPass,
    });

    if (newIns.acknowledged == true) {
      res.status(201).json({ message: "user created" });
    } else {
      res.status(500).json({ message: "some error occured" });
    }
  } catch (err) {
    return res.status(500).json({ error: "Internal server error" });
  }
});

app.post("/api/auth/login", async (req, res) => {
  const { email, password } = req.body;
  try {
    let collection = await db.collection("login");
    let user = await collection.findOne({ email });
    if (!user) {
      return res.status(401).json({ error: "No user Found" });
    }
    const validPass = await bcrypt.compare(password, user.password);
    if (!validPass) {
      res.status(401).json({ message: "Password incorrect" });
      return;
    }
    const token = jwt.sign({ email }, process.env.KEY, { expiresIn: "24h" });
    res.cookie("token", token);
    res.json({ Logged: true });
  } catch (err) {
    res.status(500).json({ message: "Internal Server Error" });
  }
});

app.listen(port, () => {
  console.log("running");
});
