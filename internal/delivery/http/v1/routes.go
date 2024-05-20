package v1

import "github.com/gin-gonic/gin"

func (h *Handler) Routes() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Recovery())
	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := router.Group("/api/v1")
	{

		command := api.Group("/command")
		{
			command.POST("/create")
			command.GET("/")
		}

		output := api.Group("/output")
		{
			output.POST("/create")
			output.GET("/:id")
		}

	}
	return router
}
