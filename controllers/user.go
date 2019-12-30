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
var orderModel = new(models.OrderModel)

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

	_, err := userModel.GetByName(data.FirstName, data.LastName,data.Company,data.Email)
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
	fmt.Println("List Users")
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
	fmt.Println("Find Iser-:>")
	var data forms.FindUserCommand
	if c.BindJSON(&data) != nil {
		c.JSON(200, gin.H{"code":1, "message": "INVALID_PARAMETERS"})
		c.Abort()
		return
	}
	profile, err := userModel.Get(data.ID)
	fmt.Println("Find MObile-:>",err)
	if err != nil {
		c.JSON(200, gin.H{"code":3, "message": "USER_NOT_EXIST"})
		c.Abort()
			return
	} else {
		c.JSON(200, gin.H{"code":0, "message": "SUCCESS", "user": profile})
		c.Abort()
			return
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

	profile, errg := userModel.Get(data.ID)
	if errg != nil {
		c.JSON(200, gin.H{"code":3, "message": "USER_NOT_EXIST"})
		c.Abort()
			return
	}
	var frsClient = new(frs.FrsClient)
	frsClient.IP ="172.22.20.175:80";
	if profile.PersonID != "" {
		fmt.Println("Delete OldPerson-:>",errg)
		frsClient.FrsDeleteUser(user.SessionID, profile.PersonID)
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
		fmt.Println("Update Image InvalidParameters")
		c.JSON(200, gin.H{"code":1, "message": "INVALID_PARAMETERS"})
		c.Abort()
		return
	}
	profile, errg := userModel.Get(data.ID)
	if errg != nil {
		c.JSON(200, gin.H{"code":3, "message": "USER_NOT_EXIST"})
		c.Abort()
			return
	}
	var frsClient = new(frs.FrsClient)
	frsClient.IP ="172.22.20.175:80";
	if profile.PersonID != "" {
		fmt.Println("Delete OldPerson-:>",profile.PersonID)
		frsClient.FrsDeleteUser(user.SessionID, profile.PersonID)
	}

	var frsPersonID = frsClient.FrsCreateUser(user.SessionID, data.ID,data.Image)
	fmt.Println("get person id=", frsPersonID)
	if  frsPersonID == "" {
		c.JSON(200, gin.H{"code":4, "message": "INVALID_IMAGE"})
		c.Abort()
		return
	}

	err := userModel.UpdateImage(data,frsPersonID)
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
	var frsVerifyFaceID = frsClient.FrsVerify(user.SessionID, data.Image,data.Threshold)
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



func (user *UserController) CreateOrder(c *gin.Context) {
	var data forms.CreateOrderCommand
	if c.BindJSON(&data) != nil {
		c.JSON(200, gin.H{"code":1, "message": "INVALID_PARAMETERS"})
		c.Abort()
		return
	}
	profile, errg := userModel.Get(data.UserId)
	if errg != nil {
		c.JSON(200, gin.H{"code":3, "message": "USER_NOT_EXIST"})
		c.Abort()
			return
	}
	orderNumber,err2 := orderModel.Create(profile,data)
	if err2 != nil {
		c.JSON(200, gin.H{"code":5, "message": "OPERATION_FAIL"})
		c.Abort()
			return
	}
	c.JSON(200, gin.H{"code":0, "message": "SUCCESS","orderNumber":orderNumber})
}


func (user *UserController) ListOrders(c *gin.Context) {
	fmt.Println("List Orders")
	var data forms.ListOrderCommand
	if c.BindJSON(&data) != nil {
		c.JSON(200, gin.H{"code":1, "message": "INVALID_PARAMETERS"})
		c.Abort()
		return
	}
	list, err := orderModel.List(data.Time)
	if err != nil || list == nil{
		var orderList = make([]models.Order, 0)
		c.JSON(200, gin.H{"code":0, "message": "SUCCESS", "orders": orderList})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"code":0, "message": "SUCCESS", "orders": list})
	}
}

func (user *UserController) FindOrder(c *gin.Context) {
	//fmt.Println("List Orders")
	var data forms.FindLastOrderCommand
	if c.BindJSON(&data) != nil {
		c.JSON(200, gin.H{"code":1, "message": "INVALID_PARAMETERS"})
		c.Abort()
		return
	}
	list, err := orderModel.Find(data)
	if err != nil || list == nil || len(list) == 0 {
		c.JSON(200, gin.H{"code":1, "message": "NO_ORDER_EXIST"})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"code":0, "message": "SUCCESS", "order": list[0]})
	}
}



func (user *UserController) FindBonus(c *gin.Context) {
	//fmt.Println("List Orders")
	var data forms.FindLastOrderCommand
	if c.BindJSON(&data) != nil {
		c.JSON(200, gin.H{"code":1, "message": "INVALID_PARAMETERS"})
		c.Abort()
		return
	}
	_, errg := userModel.Get(data.UserId)
	if errg != nil {
		c.JSON(200, gin.H{"code":3, "message": "USER_NOT_EXIST"})
		c.Abort()
			return
	}
	isSpecialBonus, err := orderModel.FindBonus(data)
	if err != nil  {
		c.JSON(200, gin.H{"code":0, "message": "SUCCESS", "order": false})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"code":0, "message": "SUCCESS", "isSpecialBonus": isSpecialBonus})
	}
}

func (user *UserController) SetBonus(c *gin.Context) {
	//fmt.Println("List Orders")
	var data forms.FindLastOrderCommand
	if c.BindJSON(&data) != nil {
		c.JSON(200, gin.H{"code":1, "message": "INVALID_PARAMETERS"})
		c.Abort()
		return
	}
	_, errg := userModel.Get(data.UserId)
	if errg != nil {
		c.JSON(200, gin.H{"code":3, "message": "USER_NOT_EXIST"})
		c.Abort()
			return
	}
	err := orderModel.SetBonus(data)
	if err != nil  {
		c.JSON(200, gin.H{"code":5, "message": "OPERATION_FAIL"})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"code":0, "message": "SUCCESS"})
	}
}


func (user *UserController) ListProducts(c *gin.Context) {

	list, err := orderModel.ListProducts()
	if err != nil || list == nil{
		var orderList = make([]models.ProductOut, 0)
		c.JSON(200, gin.H{"code":0, "message": "SUCCESS", "products": orderList})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"code":0, "message": "SUCCESS", "products": list})
	}
}


func (user *UserController) UpdateProduct(c *gin.Context) {
	//fmt.Println("List Orders")
	var data forms.UpdateProductCommand
	if c.BindJSON(&data) != nil {
		c.JSON(200, gin.H{"code":1, "message": "INVALID_PARAMETERS"})
		c.Abort()
		return
	}
	err := orderModel.UpdateProductAvailable(data)
	if err != nil  {
		c.JSON(200, gin.H{"code":5, "message": "OPERATION_FAIL"})
		c.Abort()
	} else {
		c.JSON(200, gin.H{"code":0, "message": "SUCCESS"})
	}
}
