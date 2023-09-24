var express = require('express');
var fs = require('fs');
var http = require('http');
var https = require('https');
const axios = require('axios');

function credenciales(){
    var rutas_credenciales = require("./rutas_certs.json");

    return {
        'key': fs.readFileSync(rutas_credenciales.llave),
        'cert': fs.readFileSync(rutas_credenciales.certificado)
    }
}



var app = express()


app.get('/monthly-taxes', function(req,res) {

    const URL = 'https://dummy.restapiexample.com/api/v1/employees'
    axios.get(URL).then(
        ({data}) => {
            res.send(data)
        }
    )
});

var servidor_http = http.createServer(app)
var servidor_https = https.createServer(credenciales(), app)
servidor_http.listen(8080)
servidor_https.listen(8443)