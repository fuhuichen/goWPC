package models

import (
	//"errors"
//	"reflect"
	"fmt"
	//"time"
	"goWPC/db"
	"goWPC/forms"
	"gopkg.in/mgo.v2/bson"
)

type OrderItem struct {
	Name          string         			`json:"name" bson:"name"`
	Amount        int64								`json:"amount" bson:"amount"`
	Size          string         			`json:"size" bson:"size"`
	Ice           string         			`json:"ice" bson:"ice"`
	Sugar         string         		  `json:"sugar" bson:"sugar"`
	Padding       string           		`json:"padding" bson:"padding"`
}


type OrderUser struct {
	FirstName    string         `json:"firstname" bson:"firstname"`
	LastName     string         `json:"lastname" bson:"lastname"`
	Email        string         `json:"email"  bson:"email"`
	Company      string         `json:"company" bson:"company"`
	Title        string     	  `json:"title" bson:"title"`
	Mobile       string         `json:"mobile" bson:"mobile"`
	Extend1      string 			  `json:"extend1" bson:"extend1"`
	Extend2      string 				`json:"extend2" bson:"extend2"`
}

type Order struct {
	ID           		   bson.ObjectId  `json:"id" bson:"_id,omitempty"`
	OrderNumber        int64         	`json:"orderNumber" bson:"orderNumber"`
	UserId    	       string         `json:"userId" bson:"userId"`
	Time               int64          `json:"time" bson:"time"`
	UserInfo           OrderUser			`json:"orderUser" bson:"orderUser"`
	List          		 []OrderItem 	  `json:"orderList" bson:"checkUser"`
}

type OrderModel struct{}

var mserver = "127.0.0.1"
var dbConnectOrder = db.NewConnection(mserver)

func (m *OrderModel) Create(user User, data forms.CreateOrderCommand) (id bson.ObjectId,  err error ){
	collection := dbConnectOrder.Use("wpc", "Orders")
	id = bson.NewObjectId()
	err = collection.Insert(bson.M{"_id":id,"userId": data.UserId})
  fmt.Println("Create Response:>",err)
	return id,err
}
