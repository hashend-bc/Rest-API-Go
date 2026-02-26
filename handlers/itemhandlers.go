package handlers

import (
	"net/http"
	"strconv"

	"go-rest-api/models"

	"github.com/labstack/echo/v4"
)

var Items []models.Item
var CurrentID = 1

// CREATE
func AddItem(c echo.Context) error {

	item := new(models.Item)

	if err := c.Bind(item); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid JSON",
		})
	}

	if item.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Item name is required",
		})
	}

	item.ID = CurrentID
	CurrentID++

	Items = append(Items, *item)

	return c.JSON(http.StatusCreated, item)
}

// READ
func GetItems(c echo.Context) error {
	return c.JSON(http.StatusOK, Items)
}

// UPDATE
func UpdateItem(c echo.Context) error {

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid ID",
		})
	}

	updatedItem := new(models.Item)
	if err := c.Bind(updatedItem); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid JSON",
		})
	}

	for i := range Items {
		if Items[i].ID == id {
			Items[i].Name = updatedItem.Name
			Items[i].Price = updatedItem.Price
			Items[i].Quantity = updatedItem.Quantity
			Items[i].Description = updatedItem.Description

			return c.JSON(http.StatusOK, Items[i])
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{
		"error": "Item not found",
	})
}

// DELETE
func DeleteItem(c echo.Context) error {

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid ID",
		})
	}

	for i := range Items {
		if Items[i].ID == id {
			Items = append(Items[:i], Items[i+1:]...)
			return c.JSON(http.StatusOK, map[string]string{
				"message": "Item deleted successfully",
			})
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{
		"error": "Item not found",
	})
}