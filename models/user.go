package models

import (
	"errors"
	"reflect"
	"fmt"
	"time"
	"goWPC/db"
	"goWPC/forms"
	"gopkg.in/mgo.v2/bson"
)

type CheckRecord struct {
	BoothName     string         			`json:"boothName" bson:"boothName"`
	Checked   		bool						    `json:"checked" bson:"checked"`
	Time       		int64						    `json:"time" bson:"time"`
}

type User struct {
	ID           bson.ObjectId  `json:"id" bson:"_id,omitempty"`
	FirstName    string         `json:"firstname" bson:"firstname"`
	LastName     string         `json:"lastname" bson:"lastname"`
	Email        string         `json:"email"  bson:"email"`
	Company      string         `json:"company" bson:"company"`
	Title        string     	  `json:"title" bson:"title"`
	Mobile       string         `json:"mobile" bson:"mobile"`
	Extend1      string 			  `json:"extend1" bson:"extend1"`
	Extend2      string 				`json:"extend2" bson:"extend2"`
	Country      string 				`json:"country" bson:"country"`
	Registered         bool         `json:"registered" bson:"registered"`
	CounterRegistered  bool          `json:"counterRegistered" bson:"counterRegistered"`
	RegisterTime       int64        `json:"registerTime" bson:"registerTime"`
	FaceRegistered     bool          `json:"faceRegistered" bson:"face_registerd"`
	CheckList          []CheckRecord 	`json:"checkList" bson:"checkList"`
	PersonID           string         `json:"personid" bson:"personid"`
}
type Face struct {
	Image  string      `json:"image" bson:"image"`
}

type UserModel struct{}

var server = "127.0.0.1"
var dbConnect = db.NewConnection(server)

func (m *UserModel) Create(data forms.CreateUserCommand) (id bson.ObjectId, err error ){
	collection := dbConnect.Use("wpc", "users")
	id = bson.NewObjectId()
	err = collection.Insert(bson.M{"_id":id,"firstname": data.FirstName,"lastname": data.LastName,
					"title":data.Title,"email": data.Email,"company":data.Company,"mobile":data.Mobile,
					"extend1":data.Extend1,"extend2":data.Extend2,"country":data.Country})
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
		company string,
		email string) (user User, err error) {
	collection := dbConnect.Use("wpc", "users")
	err = collection.Find(
		bson.M{"email":email,"lastname":lastname,
			"firstname":firstname,"company":company}).One(&user)
	return user, err
}

func (m *UserModel) GetByPersonID(
		personid string) (user User, err error) {
	collection := dbConnect.Use("wpc", "users")

	err = collection.Find(
		bson.M{"personid":personid}).One(&user)
	return user, err
}

func FillStruct(data map[string]interface{}, result interface{}) {
    t := reflect.ValueOf(result).Elem()
    for k, v := range data {
        val := t.FieldByName(k)
        val.Set(reflect.ValueOf(v))
    }
}

func (m *UserModel) Get(id string) (user User, err error) {
	if !bson.IsObjectIdHex(id) {
		err =errors.New("Invalid ObjectIDHex");
    return user, err
  }
	collection := dbConnect.Use("wpc", "users")
	fmt.Println("Find Result-:>",id )
	err = collection.FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

func (m *UserModel) GetFace(id string) (face Face, err error) {
	if !bson.IsObjectIdHex(id) {
		err =errors.New("Invalid ObjectIDHex");
		return face, err
	}
	collection := dbConnect.Use("wpc", "users")
	err = collection.FindId(bson.ObjectIdHex(id)).One(&face)
	return face, err
}


func (m *UserModel) UpdateCheck( f forms.UpdateCheckCommand) (err error) {
	collection := dbConnect.Use("wpc", "users")
	var user User
	err = collection.FindId(bson.ObjectIdHex(f.ID)).One(&user)
	if err != nil {
		return err
	}
	fmt.Println("Iser :>",user)
  cList := []bson.M{}
	find := false
	for _, element := range user.CheckList {
    fmt.Println(element)
		if element.BoothName == f.BoothName {
			//fmt.Println("Update booth :>"+element.BoothName)
			ctime := time.Now().Unix()
			cList  = append(cList,bson.M{"boothName": f.BoothName, "checked":element.Checked ,"time":ctime} )
			find = true
		}else{
			//fmt.Println("Old booth :>"+ element.BoothName)
			cList  = append(cList,bson.M{"boothName": element.BoothName, "checked":element.Checked,"time":element.Time} )
		}
  }
	if find ==false  {
	//	fmt.Println("Not found :>"+f.BoothName)
		cList  = append(cList,bson.M{"boothName": f.BoothName, "checked":f.Checked } )
	}

	data := bson.M{"$set": bson.M{"checkList": cList }}
		fmt.Println("Data :>",data)
	err = collection.UpdateId(bson.ObjectIdHex(f.ID), data)
	return err
}

func (m *UserModel) UpdateImage( f forms.UpdateUserImageCommand , personID string) (err error) {
	collection := dbConnect.Use("wpc", "users")
	data := bson.M{"$set": bson.M{"image": f.Image,"personid": personID, "face_registerd": true }}
	err = collection.UpdateId(bson.ObjectIdHex(f.ID), data)
	return err
}

func (m *UserModel) Update(f forms.UpdateUserCommand) (err error) {
	collection := dbConnect.Use("wpc", "users")
	updateList := bson.M{}
	//if f.Registered != nil{
	//	  updateList["registered"] = f.Registered
//	}
	if f.Registered != nil {
			updateList["registered"] = f.Registered
			updateList["registerTime"] = time.Now().Unix()
	}
	if f.CounterRegistered != nil {
			updateList["counterRegistered"] = f.CounterRegistered
	}
	if f.Email != "" {
			updateList["email"] = f.Email
	}
	if f.Mobile != "" {
			updateList["mobile"] = f.Mobile
	}
	if f.Title != "" {
			updateList["title"] = f.Title
	}
	if f.Extend1 != "" {
			updateList["extend1"] = f.Extend1
	}
	if f.Extend2 != "" {
			updateList["extend2"] = f.Extend2
	}
	//fmt.Println("Cmd Data :>",f.Registered)
	data := bson.M{"$set": updateList}
	fmt.Println("Data :>",data)
	err = collection.UpdateId(bson.ObjectIdHex(f.ID), data)
	return err
}

func (m *UserModel) Delete(id string) (err error) {
	if !bson.IsObjectIdHex(id) {
		err =errors.New("Invalid ObjectIDHex");
		return  err
	}
	collection := dbConnect.Use("wpc", "users")
	err = collection.RemoveId(bson.ObjectIdHex(id))

	return err
}
