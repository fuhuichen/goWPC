package models

import (
	"goWPC/db"
	"goWPC/forms"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID           bson.ObjectId `json:"id" bson:"_id,omitempty"`
	FirstName    string         `json:firstname`
	LastName     string         `json:lastname`
	Email        string         `json:email`
	Company      string         `json:company`
	Mobile       string         `json:mobile`
}

type UserModel struct{}

var server = "127.0.0.1"

var dbConnect = db.NewConnection(server)

func (m *UserModel) Create(data forms.CreateUserCommand) error {
	collection := dbConnect.Use("wpc", "Users")
	err := collection.Insert(bson.M{"firstname": data.FirstName,"lastname": data.LastName,
					"email": data.Email,"company":data.Company,"mobile":data.Mobile})
	return err
}

func (m *UserModel) Find() (list []User, err error) {
	collection := dbConnect.Use("wpc", "users")
	err = collection.Find(bson.M{}).All(&list)
	return list, err
}

func (m *UserModel) Get(id string) (user User, err error) {
	collection := dbConnect.Use("wpc", "users")
	err = collection.FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
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
