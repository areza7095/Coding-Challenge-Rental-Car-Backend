# Introduction
This project is a backend project built using Golang, and the Gin, Gorm, Godotenv framework

## :ledger: Index

- [Usage](#zap-usage)
  - [Installation](#electric_plug-installation)
  - [Commands](#package-commands)
- [Development](#wrench-development)
  - [File Structure](#file_folder-file-structure)
  - [Production](#rocket-deployment)  


## :zap: Usage
How to run the project

###  :electric_plug: Installation
- Install GO Compile Daemon
- Copy .env.example and rename it to .env
- Set your port to run this project
- And set the DB_URL to your postgres Database

```
$ go get github.com/githubnemo/CompileDaemon
```

###  :package: Commands
- Commands to start the project.

```
$ CompileDaemon -command="./rental-car"
```

##  :wrench: Development

###  :file_folder: File Structure
Add a file structure here with the basic details about files, below is an example.

```
.
├── controller
│   ├── carController.go
│   └── orderController.go
├── doc
│   └── Rental Car - Go.postman_collection
├── initializers
│   ├── database.go
│   └── loadEnvVariables.go
├── migrate
│   └── migrate.go
├── models
│    └── models.go
│       
├── .env
├── .env.example
├── .gitignore
├── go.mod
├── go.sum
├── main.go
├── README
└── rental-care
```

| No | File Name | Details 
|----|------------|-------|
| 1  | main.go | Entry point
| 2 | carController.go | Controller for handling Car request 
| 3 | orderController.go | Controller for handling Order request 
| 4 | Rental Car - Go.postman_collection | Documentation of REST API
| 5  | database.go | Initializer of database
| 6  | loadEnvVariables.go | Initializer of Env Variables
| 7  | models.go | Initializer models of Relational database 
| 8  | .env | Environment Variable
| 9  | .env.example | Environment Variable Example

### :rocket: Production
For using production ready use this endpoint 
```
$ https://rental-car-backend-production.up.railway.app/rentalcar/v1/get/car
```