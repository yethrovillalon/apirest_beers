API Rest Beers

Esta API esta diseñada para ser consultas a cervezas con lenguaje GO.

Comenzando 🚀
Estas instrucciones te permitirán obtener una copia del proyecto en funcionamiento en tu máquina local para propósitos de desarrollo y pruebas.

Pre-requisitos 📋
GO
MongoDB

Instalación 🔧

Descargar de los sitios oficiales

https://golang.org/
https://www.mongodb.com/es

Ejecutando las pruebas ⚙️

http://localhost/beers GET (Lista todas las cervezas)

http://localhost/beers POST (Ingresa una nueva cerveza)
Request
{
	"id": 2,
	"name": "Stout",
	"brewery": "Kroos",
	"country": "Chile",
	"price": 10,
	"currency": "CLP"
}

http://localhost/beers/{ID} (Busca una cerveza por su Id)

http://localhost:8081/beers/{ID}/boxprice (Busca una cerveza por su Id, detalle precio caja de cervezas)

Construido con 🛠️

Visual Studio Code
Cygwin64
Robo 3T

Inicializar API

1.- Instalar MongoDB
2.- Instalar Golang
3.- Instalar Cygwin64
4.- Abrir Cygwin64, verificar la instalación de Golang (comando: go version)
5.- Inicializar MongoDB, en el caso de windows con el archivo mongod.exe. En otro sistema operativo verificar la inicializacion de mongodb.
6.- Dirigirse a la carpeta donde se encuentra el proyecto donde se descargo a traves de Cygwin64. 
7.- Ejecutar comando go run *.go
8.- Ejecutar el plan de prueba por postman


Autores ✒️
Yethro Villalon (yvillalonsilva@gmail.com)

Licencia 📄
Este proyecto está bajo la Licencia (Tu Licencia) - mira el archivo LICENSE.md para detalles
