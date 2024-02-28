package server

import (
	gin "github.com/gin-gonic/gin"
	"github.com/go-api/src/infrastructure/adapter"
	"github.com/go-api/src/infrastructure/factory"
	"github.com/go-api/src/infrastructure/routes"
)

func Up() {
	// REMOVE GIN LOGS FROM TERMINAL
	gin.SetMode(gin.ReleaseMode)

	// SET UP AN GIN ROUTER
	router := gin.Default()
	router.Use(corsMiddleware())

	// CONNECT WITH DATABASE
	db := adapter.Connect()

	// db.Migrator().DropTable(&entities.UserEntity{}, &entities.TaskEntity{})

	// MIGRATE MODELS.
	// db.AutoMigrate(&entities.UserEntity{}, &entities.TaskEntity{})

	// CREATING AND INJECTING INSTANCES
	userController, tasksController := factory.InjectDepedencies(db)

	// REGISTER THE ROUTES AND HANDLES
	routes.Register(router, userController, tasksController)

	// RUN THE SERVER
	router.Run()
}

// CORS CONFIGURATION
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
