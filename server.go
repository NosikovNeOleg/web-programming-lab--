package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Product struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

func main() {
	e := echo.New()

	var products []Product

	file, err := os.Open("products.json")
	if err != nil {
		fmt.Println(err)
	}

	bytes, _ := ioutil.ReadAll(file)
	json.Unmarshal(bytes, &products)

	// Регистрируем маршруты для статичных файлов
	e.Static("/public", "public")
	e.Static("/assets", "public/assets")

	// Пример обработчика GET с получением параметров
	e.GET("/testget", func(c echo.Context) error {
		name := c.QueryParam("name")

		return c.String(http.StatusOK, "Добрый день, "+name)
	})

	e.GET("/products/update", func(c echo.Context) error {
		file, err := os.Open("products.json")
		if err != nil {
			fmt.Println(err)
		}

		var oldCount = len(products)

		bytes, _ := ioutil.ReadAll(file)
		json.Unmarshal(bytes, &products)

		var newCount = len(products)

		var message = "Updated!"

		var result = newCount - oldCount
		if result > 0 {
			message += " Added " + string(result) + " products"
		} else if result < 0 {
			message += " Removed " + string(result) + " products"
		}
		return c.String(http.StatusOK, message)
	})

	e.GET("/products/get", func(c echo.Context) error {

		return c.String(http.StatusOK, fmt.Sprint(products))
	})

	e.GET("/products/get", func(c echo.Context) error {
		page, err := strconv.Atoi(c.QueryParam("page"))
		if err == nil && page < 0 {
			return c.String(http.StatusOK, "No response")
		} else if err != nil {
			return c.String(http.StatusOK, "page is not an integer!")
		}

		var fistProduct = page * 4
		var lastProduct = fistProduct + 4

		if lastProduct > len(products) {
			lastProduct = len(products)
		}

		var response, _ = json.Marshal(products[fistProduct:lastProduct])

		return c.String(http.StatusOK, string(response))
	})

	e.GET("/products/pages", func(c echo.Context) error {

		var pagesCount = len(products) / 4
		if len(products)%4 != 0 {
			pagesCount++
		}

		return c.String(http.StatusOK, fmt.Sprint(pagesCount))
	})

	// Пример обработчика запроса POST с получением параметров
	e.POST("/testpost", func(c echo.Context) error {
		json_map := make(map[string]interface{})
		err := json.NewDecoder(c.Request().Body).Decode(&json_map)
		if err != nil {
			return err
		}

		fmt.Println(json_map)
		name := json_map["name"].(string)
		v := map[string]interface{}{
			"response": "Добрый день, " + name,
		}

		return c.JSON(http.StatusOK, v)
	})

	// Основной обработчик GET / - отдает файл index.html
	e.GET("*", func(c echo.Context) error {
		return c.File("index.html")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
