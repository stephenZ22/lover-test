package main

import (
	"math/rand"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/number", GetNumber)
	r.POST("/add", AddNumber)

	r.POST("/multiple", MultipleNumber)
	r.Run(":3002")
}

type NumberRep struct {
	Number int `json:"number"`
}

func GetNumber(c *gin.Context) {
	// n := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := rand.Intn(100)
	rep := NumberRep{n}

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
		"data": rep,
	})
}

type AddNumberForm struct {
	NumberFirst  int `json:"number_first"`
	NumberSecond int `json:"number_second"`
}

type AddNumberRep struct {
	AddNumberForm
	NumberRep
}

func AddNumber(c *gin.Context) {
	form := AddNumberForm{}
	if err := c.BindJSON(&form); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
			"data": nil,
		})

		return
	}

	result := form.NumberFirst + form.NumberSecond

	rep := AddNumberRep{
		AddNumberForm{
			NumberFirst:  form.NumberFirst,
			NumberSecond: form.NumberSecond},
		NumberRep{Number: result},
	}

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
		"data": rep,
	})
}

func MultipleNumber(c *gin.Context) {
	form := AddNumberForm{}
	if err := c.BindJSON(&form); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"msg":  err.Error(),
			"data": nil,
		})

		return
	}

	result := form.NumberFirst * form.NumberSecond

	rep := AddNumberRep{
		AddNumberForm{
			NumberFirst:  form.NumberFirst,
			NumberSecond: form.NumberSecond},
		NumberRep{Number: result},
	}

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
		"data": rep,
	})
}
