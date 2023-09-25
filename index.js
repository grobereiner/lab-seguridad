var express = require("express");
var fs = require("fs");
var http = require("http");
var https = require("https");
const axios = require("axios");

const jwt = require("jsonwebtoken");
const bcrypt = require("bcrypt");
const bodyParser = require("body-parser");

var MongoClient = require("mongodb").MongoClient;
var url = "mongodb://0.0.0.0:27017";

async function agregar_usuario(object) {
  const client = new MongoClient(url);

    var resp = { msg: "Registro exitoso" };

  try {
    await client.connect();

    const db = client.db("l-seguridad"); 
    const collection = db.collection("users");

    const existingUser = await collection.findOne({
      username: object.username,
    });

    if (!existingUser) {
      const result = await collection.insertOne(object);
      console.log(`Inserted user with username: ${object.username}`);
    } else {
      console.log(`Username ${object.username} already exists.`);
      resp = { msg: "Usuario existente" };
    }
} catch (error) {
    console.error("Error:", error);
    resp = { msg: "Usuario existente" };
  } finally {
    client.close();
  }
  return resp;
}

function credenciales() {
  var rutas_credenciales = require("./rutas_certs.json");

  return {
    key: fs.readFileSync(rutas_credenciales.llave),
    cert: fs.readFileSync(rutas_credenciales.certificado),
  };
}

async function retornar_usuario(username){
    const client = new MongoClient(url);

    try {
      await client.connect();
  
      const db = client.db('l-seguridad'); 
      const collection = db.collection('users');
  
      const user = await collection.findOne({ username });
  
      if (user) {
        console.log(`User with username ${username} found:`, user);
        return user;
      } else {
        console.log(`User with username ${username} not found.`);
        return null;
      }
    } catch (error) {
      console.error('Error:', error);
    } finally {
      client.close();
    }
}

var app = express();

app.use(bodyParser.json());

app.post("/register", async (req, res) => {
  const { username, password } = req.body;
  const hashedPassword = await bcrypt.hash(password, 10);

  const nuevo_usuario = {
    username,
    password: hashedPassword,
  };

  var resp = await agregar_usuario(nuevo_usuario);

  res.status(201).json(resp);
});

app.post("/login", async (req, res) => {
  const { username, password } = req.body;

  const usuario = await retornar_usuario(username);
  if (!usuario) {
    return res.status(401).json({ msg: "Credenciales incorrectas" });
  }

  const passwordVerify = await bcrypt.compare(password, usuario.password);
  if (!passwordVerify) {
    return res.status(401).json({ msg: "Credenciales incorrectas" });
  }

  const token = jwt.sign({ userId: usuario.id }, "ejemplo");
  res.status(200).json({ token });
});

app.get("/monthly-taxes", function (req, res) {
  const URL = "https://dummy.restapiexample.com/api/v1/employees";
  axios.get(URL).then(({ data }) => {
    res.send(data);
  });
});

var servidor_http = http.createServer(app);
var servidor_https = https.createServer(credenciales(), app);
servidor_http.listen(8080);
servidor_https.listen(8443);
