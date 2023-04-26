package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateTodo struct {
	Username string `json:"username" binding:"rquired"`
	Title    string `json:"title" binding:"rquired"`
	Message  string `json:"message" binding:"rquired"`
}

func GetAllLists(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"To Do Lists": "todoLists"})
}
func CreateTodoList(c *gin.Context){
	var input CreateTodo
	err := c.ShouldBindJSON(&input)
	if err != nil{
		
	}
}

func DeleteList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": true})
}

func Upload(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"filepath": "filepath"})
}
