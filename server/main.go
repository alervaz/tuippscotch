package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	e := echo.New()

	e.POST("/", func(c echo.Context) error {
		var person Person
		err := json.NewDecoder(c.Request().Body).Decode(&person)
		if err != nil {
			return c.String(http.StatusBadRequest, "Could not parse body: "+err.Error())
		}

		return c.String(
			http.StatusOK,
			fmt.Sprintf(
				"The person %s that is %d years old added to the database",
				person.Name,
				person.Age,
			),
		)
	})

	e.Logger.Fatal(e.Start(":3000"))
}
