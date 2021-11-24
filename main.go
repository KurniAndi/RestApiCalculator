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
	router.POST("/dividend", dividendOperation)
	router.POST("/multiple", multipleOperation)

	router.Run()
}

type numberInput struct {
	Number  int
	Number2 int
	Result  int
}

func plusOperation(c *gin.Context) {
	var numberInput numberInput

	err := c.ShouldBindJSON(&numberInput)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{

		"result": numberInput.Number + numberInput.Number2,
	})
}

func minusOperation(c *gin.Context) {
	var numberInput numberInput

	err := c.ShouldBindJSON(&numberInput)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{

		"result": numberInput.Number - numberInput.Number2,
	})
}

func dividendOperation(c *gin.Context) {
	var numberInput numberInput

	err := c.ShouldBindJSON(&numberInput)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{

		"result": numberInput.Number / numberInput.Number2,
	})
}

func multipleOperation(c *gin.Context) {
	var numberInput numberInput

	err := c.ShouldBindJSON(&numberInput)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{

		"result": numberInput.Number * numberInput.Number2,
	})
}
