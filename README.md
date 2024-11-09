# gomvc CLI

gomvc is a command-line tool that automates the creation of a structured Model-View-Controller (MVC) project template in Go. It sets up a Gin-based server with a modular structure, including controllers, models, middleware, utilities, and routes. The tool also initializes a Go module based on user input and configures the project with proper import paths.
Features

    Initializes an MVC project structure with folders for controller, models, router, middleware, pkg, and more.
    Sets up a basic Gin server with routing, middleware, and sample controller logic.
    Dynamically sets the Go module name for your project.
    Provides -create and -delete commands to build or remove the project structure.
    
### Installation

To install gomvc directly using go install, run:

go install github.com/username/gomvc@latest

Make sure to replace github.com/username/gomvc with the actual repository URL where gomvc is hosted.

This command will:

    Download and build the gomvc binary.
    Place it in your $GOPATH/bin directory, which is typically included in your $PATH, so you can run gomvc from anywhere in your terminal.

If you encounter any issues, verify that your Go environment is set up correctly and that $GOPATH/bin is in your $PATH.

### Folder Structure

Here’s the folder structure that gomvc will create:

myproject/
├── cmd/
│   └── api/
│       └── main.go            # Entry point for the Gin server
├── controller/
│   └── home_controller.go      # Sample controller
├── models/
│   └── user.go                 # Sample data model
├── middleware/
│   └── request_logger.go       # Sample middleware for request logging
├── pkg/
│   └── utility.go              # Utility functions
├── router/
│   └── router.go               # Route setup
├── config/                     # Placeholder for configuration files
└── views/                      # Placeholder for views or HTML templates


Usage

After installing gomvc, you can create or delete a Go MVC project structure using the following commands.
Create a New Project

gomvc -create <path>

Replace <path> with the desired directory path for your new project. For example:

gomvc -create ./myproject

You’ll be prompted to enter a Go module name, typically in the format github.com/username/myproject. This initializes a Go module and sets up the project with your specified module name.
Example Workflow

    Run:

gomvc -create ./myproject

Enter the module name when prompted:

Enter the project name for Go module initialization (e.g., github.com/username/myproject):

Once complete, navigate to myproject/cmd/api and run the server:

    cd myproject/cmd/api
    go run main.go

This will start a Gin server running at http://localhost:8080 with a sample route.
Delete an Existing Project

gomvc -delete <path>

Replace <path> with the path to the project you want to delete. This command will remove all directories created by gomvc and delete the go.mod file.
Help

Run the following command to display help information:

gomvc -h

Explanation of Key Components
Main Components

    cmd/api/main.go: The entry point of the Gin server. It initializes routes and starts the server.
    router/router.go: Configures the routes, middleware, and links to controllers.
    controller/home_controller.go: Contains a sample controller function that responds to HTTP requests.
    models/user.go: Provides a sample data model (User) for structuring data within the application.
    middleware/request_logger.go: Logs incoming requests with method, path, and duration. This file shows how to add custom middleware to Gin.
    pkg/utility.go: A utility folder for helper functions. The sample function PrintMessage is included to demonstrate usage.

Project Initialization

When creating a project, gomvc:

    Prompts for the module name to set up Go module imports.
    Creates each folder (controller, models, middleware, etc.) with sample files.
    Configures main.go with the correct import paths using the specified module name.

Example Code Overview

Here’s an overview of what each main file does:

    main.go (in cmd/api/): Starts the Gin server and loads the configured routes.

package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "your_module_name/router"
)

func main() {
    fmt.Println("Starting the Gin server...")
    r := gin.Default()
    router.InitializeRoutes(r)
    r.Run(":8080")
}

router.go: Initializes routes and includes middleware.

package router

import (
    "github.com/gin-gonic/gin"
    "your_module_name/controller"
    "your_module_name/middleware"
)

func InitializeRoutes(r *gin.Engine) {
    r.Use(middleware.RequestLogger())
    r.GET("/", controller.HomeController)
}

home_controller.go: A sample controller to handle the home route.

package controller

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func HomeController(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "Hello from HomeController!"})
}

request_logger.go: Logs each request with the HTTP method, path, and response time.

    package middleware

    import (
        "fmt"
        "time"
        "github.com/gin-gonic/gin"
    )

    func RequestLogger() gin.HandlerFunc {
        return func(c *gin.Context) {
            startTime := time.Now()
            c.Next()
            duration := time.Since(startTime)
            fmt.Printf("[%s] %s %s %v\n", time.Now().Format(time.RFC3339), c.Request.Method, c.Request.URL.Path, duration)
        }
    }

License

This project is licensed under the MIT License.
