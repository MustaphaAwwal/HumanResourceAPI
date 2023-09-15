package main

import (
	"context"
	"fmt"
	"hng-stage2/controllers"
	"hng-stage2/routes"
	"hng-stage2/services"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	l := log.New(os.Stdout, "human-api", log.LstdFlags)
	dbUrl := os.Getenv("DatabaseUrl")
	clientOption := options.Client().ApplyURI(dbUrl)

	client, err := mongo.Connect(context.Background(), clientOption)
	if err != nil {
        l.Fatal(err)
    }

	defer client.Disconnect(context.Background())
	fmt.Println("Connected to mongoDB")

	ctx := context.TODO()

	// mistakenly name the database hng-stage1 but no problem
	collection := client.Database("hng-stage1").Collection("Human")
	humanService := services.NewHumanService(collection, ctx)
	humanController := controllers.NewHumanController(humanService)
	humanRouter := routes.NewHumanControllerRoute(humanController)

	
	server := gin.Default()
	router := server.Group("/")

	humanRouter.HumanRoute(router)
	l.Println("Starting server on port 80")
	
	server.Run("localhost:80")	
	
}
