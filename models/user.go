package models

import (
	"fmt"
	"goWPC/db"
	"goWPC/forms"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID           bson.ObjectId  `json:"id" bson:"_id,omitempty"`
	FirstName    string         `json:firstname`
	LastName     string         `json:lastname`
	Email        string         `json:email`
	Company      string         `json:company`
	Mobile       string         `json:mobile`
	Extend1      string 			  `json:"extend1"`
	Extend2      string 				`json:"extend2"`
}

type UserModel struct{}

var server = "127.0.0.1"
var dbConnect = db.NewConnection(server)

func (m *UserModel) Create(data forms.CreateUserCommand) (id bson.ObjectId, err error ){
	collection := dbConnect.Use("wpc", "users")
	id = bson.NewObjectId()
	err = collection.Insert(bson.M{"_id":id,"firstname": data.FirstName,"lastname": data.LastName,
					"email": data.Email,"company":data.Company,"mobile":data.Mobile})
  fmt.Println("Create Response:>",err)
	return id,err
}

func (m *UserModel) List(keyword string ) (list []User, err error) {
	collection := dbConnect.Use("wpc", "users")
	var query = bson.M{}
		if keyword != "" {
		var keysarch =".*"+keyword+".*" ;
		query = bson.M{ "$or": []bson.M{{"email":bson.M{"$regex":keysarch}},
		{"firstname":bson.M{"$regex":keysarch}},
		{"lastname":bson.M{"$regex":keysarch}},
		{"company":bson.M{"$regex":keysarch}}}}
	}
	fmt.Println("List Query :>",query)
	err = collection.Find(query).All(&list)
	return list, err
}

func (m *UserModel) GetByName(
		firstname string,
		lastname string,
		company string) (user User, err error) {
	collection := dbConnect.Use("wpc", "users")
	err = collection.Find(
		bson.M{"firstname":firstname,"lastname":lastname,"company":company}).One(&user)
	return user, err
}

func (m *UserModel) GetByPersonID(
		personid string) (user User, err error) {
	collection := dbConnect.Use("wpc", "users")
	err = collection.Find(
		bson.M{"personid":personid}).One(&user)
	return user, err
}

func (m *UserModel) Get(id string) (user User, err error) {
	collection := dbConnect.Use("wpc", "users")
	err = collection.FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}


func (m *UserModel) UpdateImage( f forms.UpdateUserImageCommand) (err error) {
	collection := dbConnect.Use("wpc", "users")
	data := bson.M{"$set": bson.M{"image": f.Image,"personid": f.PersonID, "face_registerd": f.FaceRegistered }}
	err = collection.UpdateId(bson.ObjectIdHex(f.ID), data)
	return err
}

func (m *UserModel) Update(id string, data forms.UpdateUserCommand) (err error) {
	collection := dbConnect.Use("wpc", "users")
	err = collection.UpdateId(bson.ObjectIdHex(id), data)
	return err
}

func (m *UserModel) Delete(id string) (err error) {
	collection := dbConnect.Use("wpc", "users")
	err = collection.RemoveId(bson.ObjectIdHex(id))

	return err
}
