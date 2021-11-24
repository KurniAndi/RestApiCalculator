package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/plus", plusOperation)
	router.POST("/minus", minusOperation)
	router.POST("/bagi", bagiOperation)
	router.POST("/kali", kaliOperation)

	router.Run()
}

type numberInput struct {
	Number  int
	Number2 int
	Hasil   int
}

func plusOperation(c *gin.Context) {
	var numberInput numberInput

	err := c.ShouldBindJSON(&numberInput)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"number":  numberInput.Number,
		"number2": numberInput.Number2,
		"hasil":   numberInput.Number + numberInput.Number2,
	})
}

func minusOperation(c *gin.Context) {
	var numberInput numberInput

	err := c.ShouldBindJSON(&numberInput)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"number":  numberInput.Number,
		"number2": numberInput.Number2,
		"hasil":   numberInput.Number - numberInput.Number2,
	})
}

func bagiOperation(c *gin.Context) {
	var numberInput numberInput

	err := c.ShouldBindJSON(&numberInput)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"number":  numberInput.Number,
		"number2": numberInput.Number2,
		"hasil":   numberInput.Number / numberInput.Number2,
	})
}

func kaliOperation(c *gin.Context) {
	var numberInput numberInput

	err := c.ShouldBindJSON(&numberInput)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"number":  numberInput.Number,
		"number2": numberInput.Number2,
		"hasil":   numberInput.Number * numberInput.Number2,
	})
}
