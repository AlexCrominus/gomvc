package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	createFlag = flag.String("create", "", "Create the MVC structure at the specified path")
	deleteFlag = flag.String("delete", "", "Delete the MVC structure at the specified path")
	helpFlag   = flag.Bool("h", false, "Show help")
)

func createDir(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.MkdirAll(path, os.ModePerm)
	}
	return nil
}

func createFile(path, content string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		file, err := os.Create(path)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = file.WriteString(content)
		return err
	}
	return nil
}

func setupMVC(rootPath string) error {
	// Prompt for project name for go mod init
	fmt.Print("Enter the project name for Go module initialization (e.g., github.com/username/project): ")
	reader := bufio.NewReader(os.Stdin)
	projectName, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	projectName = strings.TrimSpace(projectName)

	// Initialize Go module
	cmd := exec.Command("go", "mod", "init", projectName)
	cmd.Dir = rootPath
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to initialize go module: %v", err)
	}
	fmt.Printf("Initialized Go module: %s\n", projectName)

	dirs := []string{
		filepath.Join(rootPath, "cmd/api"),
		filepath.Join(rootPath, "controller"),
		filepath.Join(rootPath, "models"),
		filepath.Join(rootPath, "pkg"),
		filepath.Join(rootPath, "config"),
		filepath.Join(rootPath, "views"),
		filepath.Join(rootPath, "router"),
		filepath.Join(rootPath, "middleware"),
	}

	for _, dir := range dirs {
		if err := createDir(dir); err != nil {
			return err
		}
	}

	// Generate main.go with dynamic import path
	mainGoContent := fmt.Sprintf(`package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"%s/router"
)

func main() {
	fmt.Println("Starting the Gin server...")
	r := gin.Default()
	router.InitializeRoutes(r)
	r.Run(":8080")
}
`, projectName)
	if err := createFile(filepath.Join(rootPath, "cmd/api", "main.go"), mainGoContent); err != nil {
		return err
	}

	// Controller: sample controller with a function that handles a request
	controllerContent := `package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// HomeController handles requests for the home route
func HomeController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello from HomeController!"})
}
`
	if err := createFile(filepath.Join(rootPath, "controller", "home_controller.go"), controllerContent); err != nil {
		return err
	}

	// Models: defining a sample struct for data
	modelContent := `package models

// User represents a sample user model
type User struct {
	ID    int    ` + "`json:\"id\"`" + `
	Name  string ` + "`json:\"name\"`" + `
	Email string ` + "`json:\"email\"`" + `
}
`
	if err := createFile(filepath.Join(rootPath, "models", "user.go"), modelContent); err != nil {
		return err
	}

	// Utility: generic utility function
	pkgContent := `package pkg

import "fmt"

// PrintMessage prints a message to the console
func PrintMessage(msg string) {
	fmt.Println(msg)
}
`
	if err := createFile(filepath.Join(rootPath, "pkg", "utility.go"), pkgContent); err != nil {
		return err
	}

	// Router: sets up routes and includes middleware
	routerContent := fmt.Sprintf(`package router

import (
	"github.com/gin-gonic/gin"
	"%s/controller"
	"%s/middleware"
)

// InitializeRoutes sets up the application's routes
func InitializeRoutes(r *gin.Engine) {
	r.Use(middleware.RequestLogger())

	r.GET("/", controller.HomeController)
}
`, projectName, projectName)
	if err := createFile(filepath.Join(rootPath, "router", "router.go"), routerContent); err != nil {
		return err
	}

	// Middleware: request logger as an example
	middlewareContent := `package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// RequestLogger logs each request with method, path, and duration
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		duration := time.Since(startTime)
		fmt.Printf("[%s] %s %s %v\n", time.Now().Format(time.RFC3339), c.Request.Method, c.Request.URL.Path, duration)
	}
}
`
	if err := createFile(filepath.Join(rootPath, "middleware", "request_logger.go"), middlewareContent); err != nil {
		return err
	}

	return nil
}

func deleteMVC(rootPath string) error {
	dirs := []string{"cmd", "controller", "models", "pkg", "config", "views", "router", "middleware"}

	for _, dir := range dirs {
		if err := os.RemoveAll(filepath.Join(rootPath, dir)); err != nil {
			return err
		}
	}

	// Remove go.mod if it exists
	goModPath := filepath.Join(rootPath, "go.mod")
	if _, err := os.Stat(goModPath); err == nil {
		if err := os.Remove(goModPath); err != nil {
			return fmt.Errorf("failed to delete go.mod: %v", err)
		}
		fmt.Println("Deleted go.mod file.")
	}

	return nil
}

func showHelp() {
	fmt.Println("Usage: gomvc [OPTIONS]")
	fmt.Println("\nOptions:")
	fmt.Println("  -create <path>\tCreate the MVC structure at the specified path")
	fmt.Println("  -delete <path>\tDelete the MVC structure at the specified path")
	fmt.Println("  -h\t\t\tShow this help message")
}

func main() {
	flag.Parse()

	if *helpFlag {
		showHelp()
		return
	}

	if *createFlag != "" {
		fmt.Println("Creating MVC structure...")
		if err := setupMVC(*createFlag); err != nil {
			fmt.Printf("Error setting up MVC structure: %v\n", err)
		} else {
			fmt.Println("MVC structure created successfully!")
		}
	} else if *deleteFlag != "" {
		fmt.Println("Deleting MVC structure...")
		if err := deleteMVC(*deleteFlag); err != nil {
			fmt.Printf("Error deleting MVC structure: %v\n", err)
		} else {
			fmt.Println("MVC structure deleted successfully!")
		}
	} else {
		showHelp()
	}
}
