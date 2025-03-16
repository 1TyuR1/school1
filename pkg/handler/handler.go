package handler

import "github.com/labstack/echo/v4"

type Handler struct {
}

func (h *Handler) InitRoutes() *echo.Echo {
	router := echo.New()

	auth := router.Group("/auth")
	{
		auth.POST("/auth", h.logIn)
	}

	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createLists)
			lists.GET("/", h.getAllLists)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)
			lists.GET("/:id", h.getListById)

			items := lists.Group(":id/items")
			{
				items.POST("/", h.createItems)
				items.GET("/", h.getAllItems)
				items.PUT("/:id", h.updateItem)
				items.DELETE("/:id", h.deleteItem)
				items.GET("/:id", h.getItemById)
			}
		}
	}

	return router
}
