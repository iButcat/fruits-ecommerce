# fruits-ecommerce

# Purpose

This project purpose was for an coding challenge from a company. It's a full
stack application which is dockerized. 

# Design Patterns 

This project was a good opportunity to follow some good design patterns. I tried to make something clean by using some layers. As a controller layer which call a service layer which call a repository layer ! 

# Tech used 
# Backend 

- Go gin 
- Gorm 
- Viper
- Jwt-Go

# Frontend

- React
- React-Boostrap
- React-Dom-Router

# Database 

- Postgresql

# Devops 

- Docker 

# How to use the project

# With Docker

```docker-compose up --build```

# Without docker 

```go install ```

if you did not migrate all the models you should uncommented 
the noot function in main and run: 

```go run *.go```

if you did so you can comment both noot function and simply use:

```go run main.go```


About the frontend you need to install all the dependencies as well.

````yarn install```

And then start the project: 

```yarn start```

# Endpoints 

I'm going to add all the endpoints ASAP or implement SWAGGER. 

# Architecture Documents 

Not the most complete documentation, it's still usefull. Both images correspond to the first and second step of the C4 model.

![context](https://github.com/iButcat/fruits-ecommerce/blob/main/architecture/cinemo-Context.jpg)

![container](https://github.com/iButcat/fruits-ecommerce/blob/main/architecture/cinemo-Container%20.jpg)