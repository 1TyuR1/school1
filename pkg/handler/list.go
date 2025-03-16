package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) createLists(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
func (h *Handler) getAllLists(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
func (h *Handler) getListById(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
func (h *Handler) updateList(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
func (h *Handler) deleteList(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
