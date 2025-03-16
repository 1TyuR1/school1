package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) createItems(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
func (h *Handler) getAllItems(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
func (h *Handler) getItemById(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
func (h *Handler) updateItem(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
func (h *Handler) deleteItem(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
