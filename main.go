package main

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Product struct {
	ID    int    `json:"id`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

func main() {
	app := fiber.New()

	products := []Product{
		{1, "product A", 10000, 12},
		{2, "product B", 12000, 13},
		{3, "product C", 13000, 14},
	}

	app.Get("/products", func(c *fiber.Ctx) error {
		return c.JSON(products)
	})

	app.Get("/products/:id", func(c *fiber.Ctx) error {
		paramId := c.Params("id")
		id, _ := strconv.Atoi(paramId)
		var foundProduct Product
		for _, p := range products {
			if p.ID == id {
				foundProduct = p
				break
			}
		}
		return c.JSON(foundProduct)
	})

	app.Post("/products", func(c *fiber.Ctx) error {
		var newProduct Product
		c.BodyParser(&newProduct)
		newProduct.ID = len(products) + 1
		products = append(products, newProduct)

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "Product created successfully",
			"product": newProduct,
		})
	})

	app.Put("/products/:id", func(c *fiber.Ctx) error {
		id, _ := strconv.Atoi(c.Params("id"))

		var updatedProduct Product
		c.BodyParser(&updatedProduct)

		var foundIndex int = -1
		for i, p := range products {
			if p.ID == id {
				foundIndex = i
				updatedProduct.ID = foundIndex + 1
				break
			}
		}

		products[foundIndex] = updatedProduct
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": fmt.Sprintf("Product with ID %d updated successfully", id),
			"product": updatedProduct,
		})
	})

	app.Delete("/products/:id", func(c *fiber.Ctx) error {
		id, _ := strconv.Atoi(c.Params("id"))

		var foundIndex int = -1
		for i, p := range products {
			if p.ID == id {
				foundIndex = i
				break
			}
		}

		products = append(products[:foundIndex], products[foundIndex+1:]...)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": fmt.Sprintf("Product with ID %d deleted successfully", id),
		})
	})

	app.Listen(":3000")
}
