package models

import (
	//"errors"
//	"reflect"
	"fmt"
	"time"
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

type DateItem struct {
	ID           	bson.ObjectId     `json:"id" bson:"_id,omitempty"`
	Date          string         		`json:"date" bson:"date"`
	Count         int  							`json:"count" bson:"count"`
}

type Special struct {
	UserId    	  string              `json:"userId" bson:"userId"`
	Date          string         			`json:"date" bson:"date"`
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
	OrderNumber        int64         	`json:"orderNumber" bson:"orderNumber"`
	UserId    	       string         `json:"userId" bson:"userId"`
	Time               int64          `json:"time" bson:"time"`
	UserInfo           OrderUser			`json:"orderUser" bson:"orderUser"`
	List          		 []OrderItem 	  `json:"orderList" bson:"orderList"`
}

type Product struct {
	ID           	bson.ObjectId     `bson:"_id,omitempty"`
	ProductId     string         		`json:"productId" bson:"productId"`
	Name          string         		`json:"name" bson:"name"`
	Available     bool							`json:"count" bson:"count"`
}
type ProductOut struct {
	ProductId     string         		`json:"productId" bson:"productId"`
	Name          string         		`json:"name" bson:"name"`
	Available     bool							`json:"available" bson:"available"`
}

type OrderModel struct{}

var mserver = "127.0.0.1"
var dbConnectOrder = db.NewConnection(mserver)

func (m *OrderModel) Create(user User, data forms.CreateOrderCommand) (orderNumber string,  err error ){
	collection := dbConnectOrder.Use("wpc", "Orders")
	countCollections := dbConnectOrder.Use("wpc", "Counts")
	fmt.Println("Create Data>",data)
	id := bson.NewObjectId()
	cList := []bson.M{}
	for _, element := range data.OrderList {
		cList  = append(cList,bson.M{"name": element.Name, "amount":element.Amount ,"size":element.Size,
				"ice":element.Ice,"sugar":element.Sugar,"padding":element.Padding} )
	}
	ctime := time.Now().Unix()
	date := time.Now().Format("2006-01-02")
	fmt.Println("Date >",date)
	var count = 1
	var dataInfo DateItem
	err0 := countCollections.Find(
		bson.M{"date":date}).One(&dataInfo)
	if err0 == nil {
			fmt.Println("Count >",dataInfo)
			count = (dataInfo.Count +1 )%10000
			udata := bson.M{"$set": bson.M{"count":count }}
			fmt.Println("Data :>",udata)
			var orderId = dataInfo.ID.Hex()
			fmt.Println("OrderId :>",orderId)
			countCollections.UpdateId(bson.ObjectIdHex(orderId), udata)

	}else{
			fmt.Println("Create New ")
			cid := bson.NewObjectId()
			countCollections.Insert(bson.M{"_id":cid,"date": date ,"count": count})
	}
	count = count % 10000
	orderNumber = fmt.Sprintf("%04d", count)
	userInfo := bson.M{"firstname": user.FirstName,"lastname": user.LastName,
					"title":user.Title,"email": user.Email,"company":user.Company,"mobile":user.Mobile}
	err = collection.Insert(bson.M{"_id":id,"userId": data.UserId,
		 "time":ctime,"orderList":cList, "orderUser":userInfo, "orderNumber":orderNumber})
  fmt.Println("Create Response:>",err)
	return orderNumber,err
}

func (m *OrderModel) List(time int64 ) (list []Order, err error) {
	fmt.Println("Time Start",time)
	collection := dbConnectOrder.Use("wpc", "Orders")
	var query = bson.M{"time":bson.M{"$gt":time}}
	err = collection.Find(query).All(&list)
	return list, err
}


func (m *OrderModel) Find(data forms.FindLastOrderCommand) (order []Order,  err error ){
	collection := dbConnectOrder.Use("wpc", "Orders")
	var query = bson.M{"userId":data.UserId}
	err = collection.Find(query).Sort("[{\"time\":-1}]").Limit(1).All(&order)
	return order,err
}


func (m *OrderModel) FindBonus(data forms.FindLastOrderCommand) (isOrderToday bool,  err error ){
	collection := dbConnectOrder.Use("wpc", "Bonus")
	date := time.Now().Format("2006-01-02")
	var query = bson.M{"userId":data.UserId,"date":date}
	var special []Special
	err = collection.Find(query).All(&special)
	if err != nil || special == nil || len(special) ==0 {
		 isOrderToday = false
	}else{
		isOrderToday  = true
	}
	return isOrderToday,err
}


func (m *OrderModel) SetBonus(data forms.FindLastOrderCommand) (err error ){
	collection := dbConnectOrder.Use("wpc", "Bonus")
	date := time.Now().Format("2006-01-02")
	cid := bson.NewObjectId()
	err = collection.Insert(bson.M{"_id":cid,"date": date ,"userId":data.UserId})
	return err
}


func (m *OrderModel) UpdateProductAvailable( f forms.UpdateProductCommand ) (err error) {
	collection := dbConnect.Use("wpc", "products")
	var product Product
	err = collection.Find(bson.M{"productId":f.ProductId}).One(&product)
	if err != nil{
		return err
	}
	data := bson.M{"$set": bson.M{"available": f.Available }}
	err = collection.UpdateId(product.ID, data)
	return err
}

func (m *OrderModel) ListProducts( ) (list []ProductOut,err error) {
	collection := dbConnect.Use("wpc", "products")
	var query = bson.M{}
	err = collection.Find(query).All(&list)
	return list, err
}
