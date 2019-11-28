package controllers

import (
	"fmt"
	"goWPC/forms"
	"goWPC/models"
	"goWPC/frs"
	"github.com/gin-gonic/gin"
	"time"
)

var userModel = new(models.UserModel)

type UserController struct{
	SessionID string
	Messages chan frs.FRSWSResponse
}

func (user *UserController) Create(c *gin.Context) {
	var data forms.CreateUserCommand
	if c.BindJSON(&data) != nil {
		c.JSON(200, gin.H{"code":1, "message": "INVALID_PARAMETERS"})
		c.Abort()
		return
	}

	_, err := userModel.GetByName(data.FirstName, data.LastName,data.Company)
	fmt.Println("get by name err:>", err)
	if err == nil {
		c.JSON(200, gin.H{"code":2, "message": "USER_EXIST"})
		c.Abort()
		return
	}

	id,err := userModel.Create(data)
	if err != nil {
		c.JSON(200, gin.H{"code":5, "message": "OPERATION_FAIL"})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"code":0, "message": "SUCCESS","id":id})
}

func (user *UserController) List(c *gin.Context) {
	var data forms.ListUserCommand
	c.BindJSON(&data)
	list, err := userModel.List(data.Keyword)
	if err != nil {
		var userList = make([]models.User, 0)
		c.JSON(200, gin.H{"code":0, "message": "SUCCESS", "userList": userList})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"code":0, "message": "SUCCESS", "userList": list})
	}
}

func (user *UserController) Find(c *gin.Context) {
	var data forms.FindUserCommand
	if c.BindJSON(&data) != nil {
		c.JSON(200, gin.H{"code":1, "message": "INVALID_PARAMETERS"})
		c.Abort()
		return
	}
	profile, err := userModel.Get(data.ID)
	if err != nil {
		c.JSON(200, gin.H{"code":3, "message": "USER_NOT_EXIST"})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"code":0, "message": "SUCCESS", "user": profile})
	}
}

func (user *UserController) FindFace(c *gin.Context) {
	var data forms.FindUserCommand
	if c.BindJSON(&data) != nil {
		c.JSON(200, gin.H{"code":1, "message": "INVALID_PARAMETERS"})
		c.Abort()
		return
	}
	profile, err := userModel.GetFace(data.ID)
	if err != nil {
		c.JSON(200, gin.H{"code":3, "message": "USER_NOT_EXIST"})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"code":0, "message": "SUCCESS", "image": profile.Image})
	}
}

func (user *UserController) Update(c *gin.Context) {
	var data forms.UpdateUserCommand
	if c.BindJSON(&data) != nil {

		c.JSON(200, gin.H{"code":1, "message": "INVALID_PARAMETERS"})
		c.Abort()
		return
	}
	err := userModel.Update(data)
	if err != nil {
		c.JSON(200, gin.H{"code":5, "message": "OPERATION_FAIL"})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"code":0, "message": "SUCCESS"})
}

func (user *UserController) Delete(c *gin.Context) {
	var data forms.DeleteUserCommand
	if c.BindJSON(&data) != nil {
		c.JSON(200, gin.H{"code":1, "message": "INVALID_PARAMETERS"})
		c.Abort()
		return
	}
	err := userModel.Delete(data.ID)
	if err != nil {
		c.JSON(200, gin.H{"code":5, "message": "OPERATION_FAIL"})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"code":0, "message": "SUCCESS"})
}

func (user *UserController) UpdateImage(c *gin.Context) {
	var data forms.UpdateUserImageCommand
	if c.BindJSON(&data) != nil {
		c.JSON(200, gin.H{"code":1, "message": "INVALID_PARAMETERS"})
		c.Abort()
		return
	}
	var frsClient = new(frs.FrsClient)
	frsClient.IP ="172.22.20.175:80";
	var frsPersonID = frsClient.FrsCreateUser(user.SessionID, data.ID,data.Image)
	fmt.Println("get person id=", frsPersonID)
	if  frsPersonID == "" {
		c.JSON(200, gin.H{"code":4, "message": "INVALID_IMAGE"})
		c.Abort()
		return
	}

	data.PersonID = frsPersonID
  data.FaceRegistered = true

	err := userModel.UpdateImage(data)
	if err != nil {
		c.JSON(200, gin.H{"code":5, "message": "OPERATION_FAIL"})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"code":0, "message": "SUCCESS"})
}


func (user *UserController) UpdateCheck(c *gin.Context) {
	//	fmt.Println("update check=")
	var data forms.UpdateCheckCommand
	if c.BindJSON(&data) != nil {
			fmt.Println("data id=", data)
		c.JSON(200, gin.H{"code":1, "message": "INVALID_PARAMETERS"})
		c.Abort()
		return
	}
	err := userModel.UpdateCheck(data)
	if err != nil {
		c.JSON(200, gin.H{"code":5, "message": "OPERATION_FAIL"})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"code":0, "message": "SUCCESS"})
}





func (user *UserController) VerifyImage(c *gin.Context) {
	var data forms.VerifyImageCommand
	if c.BindJSON(&data) != nil {
		fmt.Println("VerifyImage=", data)
		c.JSON(200, gin.H{"code":1, "message": "INVALID_PARAMETERS"})
		c.Abort()
		return
	}
	var userList = make([]models.User, 0)
	var frsClient = new(frs.FrsClient)
	frsClient.IP ="172.22.20.175:80";
	var frsVerifyFaceID = frsClient.FrsVerify(user.SessionID, data.Image)
	fmt.Println("get face id=", frsVerifyFaceID)
	if  frsVerifyFaceID == "" {
		c.JSON(200, gin.H{"code":0, "message": "SUCCESS", "userList":userList})
		c.Abort()
		return
	}
	personID := ""
	fmt.Println("Wait for face face" +frsVerifyFaceID )
	for i := 1; i < 20; i++ {
		select {
    	case msg := <-user.Messages:
					fmt.Println("received face", msg)
					if msg.VerifyFaceID == frsVerifyFaceID{
						personID = msg.PersonID
						break
					}
			default:
			}
		  //msg := <- user.Messages
	    time.Sleep(time.Duration(200)*time.Millisecond)
		//	fmt.Println("Test faceid =", msg)
	}
	if  personID == "" {
		c.JSON(200, gin.H{"code":0, "message": "SUCCESS", "userList":userList})
		c.Abort()
		return
	}
  fmt.Println("Person ID=>", personID)
  var profileList  [1]models.User
	profile, err := userModel.GetByPersonID(personID)
	//fmt.Println("get by person id err:>", err)
	if err != nil {
		c.JSON(200, gin.H{"code":0, "message": "SUCCESS", "userList":userList})
		c.Abort()
		return
	}
	fmt.Println("Find Person =>", profile)
	profileList[0] = profile
	c.JSON(200, gin.H{"code":0, "message": "SUCCESS", "userList":profileList})
}
