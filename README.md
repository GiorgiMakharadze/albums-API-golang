# Albums API with Go and Firebase

## Overview

This project demonstrates a basic CRUD (Create, Read, Update, Delete) API built using Go (or Golang) and Firebase Firestore for managing a collection of albums. The API allows users to perform the following operations:

- Get a list of all albums
- Get details of a specific album by ID
- Create a new album
- Update an existing album by ID
- Delete an album by ID

Utilizing the speed and simplicity of Go, along with Firebase's scalable NoSQL cloud database, this API is a robust start for any album management application.

## Features

- **Go Modules**: The project utilizes Go Modules for managing dependencies, ensuring easy setup and reproducible builds.
- **Gin Web Framework**: Gin is a web framework written in Go, utilized in this project for handling HTTP requests and routing.
- **Firebase Firestore**: Firebase Firestore is a flexible, scalable database for mobile, web, and server development from Firebase and Google Cloud.
- **Input Validation**: Leveraging the `go-playground/validator` package, the API ensures all incoming data adheres to expected formats before performing operations.
- **Error Handling**: Proper error handling is implemented to guide the users with the correct status codes and error messages.

## Setup

### Prerequisites

- [Go](https://golang.org/doc/install) (v1.15 or later recommended)
- [Firebase Project](https://firebase.google.com/)
- Firebase Admin SDK private key file

### Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/GiorgiMakharadze/albums-API-golang.git
   ```
