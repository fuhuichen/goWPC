package controllers

import (
	"fmt"
	"goWPC/forms"
	"goWPC/models"
	"github.com/gin-gonic/gin"
)

var userModel = new(models.UserModel)

type UserController struct{}

func (user *UserController) Create(c *gin.Context) {
	var data forms.CreateUserCommand
	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"code":1, "message": "INVALID_PARAMETERS"})
		c.Abort()
		return
	}

	_, err := userModel.GetByName(data.FirstName, data.LastName,data.Company)
	fmt.Println("get by name err:>", err)
	if err == nil {
		c.JSON(406, gin.H{"code":2, "message": "USER_EXIST"})
		c.Abort()
		return
	}

	id,err := userModel.Create(data)
	if err != nil {
		c.JSON(406, gin.H{"code":5, "message": "OPERATION_FAIL"})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"code":0, "message": "SUCCESS","id":id})
}

func (user *UserController) Find(c *gin.Context) {
	list, err := userModel.Find()
	if err != nil {
		c.JSON(404, gin.H{"message": "Find Error", "error": err.Error()})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"data": list})
	}
}

func (user *UserController) Get(c *gin.Context) {
	id := c.Param("id")
	profile, err := userModel.Get(id)
	if err != nil {
		c.JSON(404, gin.H{"message": "User not found", "error": err.Error()})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"data": profile})
	}
}

func (user *UserController) Update(c *gin.Context) {
	id := c.Param("id")
	data := forms.UpdateUserCommand{}

	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"message": "Invalid Parameters"})
		c.Abort()
		return
	}
	err := userModel.Update(id, data)
	if err != nil {
		c.JSON(406, gin.H{"message": "User Could Not Be Updated", "error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"message": "User Updated"})
}

func (user *UserController) Delete(c *gin.Context) {
	var data forms.DeleteUserCommand
	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"code":1, "message": "INVALID_PARAMETERS"})
		c.Abort()
		return
	}
	err := userModel.Delete(data.ID)
	if err != nil {
		c.JSON(406, gin.H{"code":5, "message": "OPERATION_FAIL"})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"code":0, "message": "SUCCESS"})
}

func (user *UserController) UpdateImage(c *gin.Context) {
	var data forms.UpdateUserImageCommand
	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"code":1, "message": "INVALID_PARAMETERS"})
		c.Abort()
		return
	}
	err := userModel.UpdateImage(data)
	if err != nil {
		c.JSON(406, gin.H{"code":5, "message": "OPERATION_FAIL"})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"code":0, "message": "SUCCESS"})
}
