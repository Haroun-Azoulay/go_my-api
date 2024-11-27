# <p align="center">My_Api - API REST GOlang with jwt token </p>


## Context
This project is a RESTful API built in Go (Golang) to demonstrate how to create a basic API for a library system. It is designed as a learning resource for students in an IT school (level: Bac +2).

The API provides essential functionalities like managing books, users, and authentication via JWT (JSON Web Tokens). It's a practical example to teach the fundamentals of API development, authentication, and interaction with a database using Go.

Students will explore:

- Structuring a Go project for API development.
- Implementing CRUD (Create, Read, Update, Delete) operations for resources (books and users).
- Securing API endpoints using JWT for role-based access (admin and user privileges).
 - Managing data persistence using a relational database (MariaDB).

The project can be extended to cover more advanced topics such as middleware, testing, and deployment.


## Subject

 Develop an API to manage a library system where users can:
    
  
- Users:
    Register and log in to the system.
    Access the list of available books.
    Update their user information (restricted actions based on roles).
-Books:
    Add new books (admin only).
    Update or delete books (admin only).
    View the list of books (open to all users).
    
    Students will learn how to implement the above features while following industry standards for API design and security.

## Screens

![Cover](https://github.com/Haroun-Azoulay/go_my-api/blob/main/img/my-api.png)

## 🛠️ Tech Stack

- Language: [Go](https://go.dev/doc/)(Golang version go1.23.3)
- Framework: [Gin Gonic for API routing](https://github.com/gin-gonic/gin?tab=readme-ov-file)  
- Database: MariaDB for storing users and books
- Authentication: JSON Web Tokens (JWT) for secure API acces
- Containerization: Docker for easy setup and deployment.

## 🧐 Features    

- Authentication: Secure login and JWT-based role management
- Books Management: CRUD operations for books
- Role-based Access: Admin-specific actions like adding or removing books
- User Management: Create and update user information

## 🛠️ Install Dependencies    

1. If not already done, [install Docker Compose](https://docs.docker.com/compose/install/) (v2.10+)
2. Run `bash script --build` to set up and start a fresh Golang project (you can enter --help if you need help)
  




## ➤ Routes References
You have all the POSTMAN calls on the project

        
        
## ❤️ Support  
A simple star to this project repo is enough to keep me motivated on this project for days. If you find your self very much excited with this project let me know with a tweet.

If you have any questions, feel free to reach out to me on.
        
## 🙇 Author
#### Haroun Azoulay

- Github: [@Haroun-Azoulay](https://github.com/Haroun-Azoulay)