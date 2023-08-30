package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Status bool   		`json:"status"`
	Message string 		`json:"message"`
	Data interface{}	`json:"data"`
}

type News struct {
	Title string 		`json:"title"`
	Content string 		`json:"content"`
}

func main(){
	e := echo.New()
	e.GET("/news", GetNewsController)
	e.POST("/news", AddNewsController)
	e.GET("/news/:id", GetDetailNewsController)
	e.Start(":8000")
}

func AddNewsController(c echo.Context) error {
	var newsRequest News
	c.Bind(&newsRequest)

	return c.JSON(http.StatusCreated, BaseResponse{
		Status: true,
		Message: "success",
		Data: newsRequest,
	})
}

func GetDetailNewsController(c echo.Context) error {
	var id = c.Param("id")
	var data = News{"A", "A"}

	return c.JSON(http.StatusOK, BaseResponse{
		Status: true,
		Message: id,
		Data: data,
	})
}

func GetNewsController(c echo.Context) error {
	country := c.QueryParam("country")
	
	
	var data []News

	// dummy
	data = append(data, News{"A", "A"})
	data = append(data, News{"B", "B"})

	return c.JSON(http.StatusOK, BaseResponse{
		Status: true,
		Message: country,
		Data: data,
	})
}





