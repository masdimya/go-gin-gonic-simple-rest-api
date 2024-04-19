package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
	password string
}

func main() {
	var users = []User{}
	var idCounter int = 1

	r := gin.Default()

	r.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, users)
	})

	r.POST("/users", func(ctx *gin.Context) {
		var user User

		ctx.ShouldBindJSON(&user)

		user.Id = idCounter
		users = append(users, user)
		idCounter += 1

		ctx.JSON(http.StatusCreated, user)
	})

	r.GET("/users/:id", func(ctx *gin.Context) {
		var user User

		paramId := ctx.Param("id")
		userId, err := strconv.Atoi(paramId)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status": "invalid parameter",
			})
			return
		}

		for _, v := range users {
			if v.Id == userId {
				user = v
				break
			}
		}

		if user.Id == 0 {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status": "data not found",
			})

			return
		}

		ctx.JSON(http.StatusOK, user)
	})

	r.PATCH("/users/:id", func(ctx *gin.Context) {
		var user User

		ctx.ShouldBindJSON(&user)

		paramId := ctx.Param("id")
		userId, err := strconv.Atoi(paramId)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status": "invalid parameter",
			})
			return
		}

		for i, v := range users {
			if v.Id == userId {

				user.Id = v.Id

				users[i].Name = user.Name
				users[i].Address = user.Address
				users[i].Age = user.Age
				users[i].Email = user.Email

				break
			}
		}

		if user.Id == 0 {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status": "data not found",
			})

			return
		}

		ctx.JSON(http.StatusCreated, user)
	})

	r.DELETE("/users/:id", func(ctx *gin.Context) {

		var user User

		paramId := ctx.Param("id")
		userId, err := strconv.Atoi(paramId)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"status": "invalid parameter",
			})
			return
		}

		indexToRemove := -1

		for i, v := range users {
			if v.Id == userId {
				indexToRemove = i
				user = v
				break
			}
		}

		if indexToRemove == -1 {
			ctx.JSON(http.StatusNotFound, gin.H{
				"status": "data not found",
			})

			return

		}

		users = append(users[:indexToRemove], users[indexToRemove+1:]...)

		ctx.JSON(http.StatusOK, user)
	})

	r.Run()
}
