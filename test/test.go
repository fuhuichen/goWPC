package main

import (
  "encoding/json"
  "fmt"
	"net/http"
  "bytes"
  "io/ioutil"
  "bufio"
  "encoding/base64"
  "os"
  "encoding/csv"
  "log"
)

type FRSLoginResponse struct {
	SessionID   string `json:"sessionId"`
}
type FRSCreateUserResponse struct {
	Message   string `json:"message"`
  PersonID   string `json:"persion_id"`
}
type FRSVerifyResponse struct {
	Message   string `json:"message"`
  VerifyFaceID   string `json:"verify_face_id"`
}


type GeneralResponse struct {
	Code     int `json:"code"`
	Message  string `json:"message"`
}

type CreateUserResponse struct {
	Code     int `json:"code"`
	Message  string `json:"message"`
  ID      string `json:"id"`
}

type User struct {
  ID           string  `json:"id" bson:"id"`
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
  PersonID           string         `json:"personid" bson:"personid"`
}
type ListUserResponse struct {
	Code     int `json:"code"`
	Message  string `json:"message"`
  UserList     []User  `json:"userList"`
}


func callAPI(url string,jsonStr []byte)(response string){
  fmt.Println("URL:>", url)
  req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
  req.Header.Set("Content-Type", "application/json")
  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
      panic(err)
  }
  defer resp.Body.Close()
  //fmt.Println("response Status:", resp.Status)
  //fmt.Println("response Headers:", resp.Header)
  body, _ := ioutil.ReadAll(resp.Body)
  fmt.Println("response Body:", string(body))
  response = string(body);
  return response;
}

func deleteUser(id string ) (response string){
        s := fmt.Sprintf("{\"id\":\"%s\"}",id)
      //    fmt.Println("Create Response:>",s )
        var jsonStr = []byte(s);
        url := "http://172.22.20.97:80/api/user/delete"
        return callAPI(url,jsonStr);
}

func findUser(id string ) (response string){
        s := fmt.Sprintf("{\"id\":\"%s\"}",id)
      //    fmt.Println("Create Response:>",s )
        var jsonStr = []byte(s);
        url := "http://172.22.20.97:80/api/user/info"
        return callAPI(url,jsonStr);
}

func findUserFace(id string ) (response string){
        s := fmt.Sprintf("{\"id\":\"%s\"}",id)
      //    fmt.Println("Create Response:>",s )
        var jsonStr = []byte(s);
        url := "http://172.22.20.97:80/api/user/face"
        return callAPI(url,jsonStr);
}

func updateUser(id string, email string ,title string,registered bool,counterRegistered bool ) (response string){
        s := fmt.Sprintf("{\"id\":\"%s\",\"email\":\"%s\",\"title\":\"%s\",\"registered\":%t,\"counterRegistered\":%t}",id,email,title, registered,counterRegistered)
      //    fmt.Println("Create Response:>",s )
        var jsonStr = []byte(s);
        url := "http://172.22.20.97:80/api/user/update"
        return callAPI(url,jsonStr);
}

func updateBoothCheck(id string ,boothName string, checked bool) (response string){
        s := fmt.Sprintf("{\"id\":\"%s\",\"boothName\":\"%s\",\"checked\":%t}",id,boothName,checked)
        fmt.Println("Create Response:>",s )
        var jsonStr = []byte(s);
        url := "http://172.22.20.97:80/api/user/updateCheck"
        return callAPI(url,jsonStr);
}

func createUser(firstname string, lastname string, company string, title string,
        email string ,mobile string ,extend1 string,extend2 string ) (response string){
        s := fmt.Sprintf("{\"firstname\":\"%s\",\"lastname\":\"%s\",\"company\":\"%s\",\"title\":\"%s\",\"email\":\"%s\",\"mobile\":\"%s\", \"extend1\":\"%s\",\"extend2\":\"%s\"}",
                  firstname,lastname,company,title, email,mobile,extend1,extend2)
          fmt.Println("Create Response:>",s )
        var jsonStr = []byte(s);
        url := "http://172.22.20.97:80/api/user/create"
        return callAPI(url,jsonStr);
}

func listUser(keyword string ) (response string){
        s := fmt.Sprintf("{\"keyword\":\"%s\"}",keyword)
          fmt.Println("Create Response:>",s )
        var jsonStr = []byte(s);
        url := "http://172.22.20.97:80/api/user/list"
        return callAPI(url,jsonStr);
}
/*
* @apiParam     {String}  orderList.name      餐點名稱
* @apiParam     {Number}  orderList.amount    餐點數量
* @apiParam     {String}  orderList.ice       餐點冰度
* @apiParam     {String}  orderList.sugar     餐點甜度
* @apiParam     {String}  orderList.size      餐點大小
* @apiParam     {String}  orderList.padding   附加品項
*/
func createOrder(id string ) (response string){
        s := fmt.Sprintf("{\"userId\":\"%s\",\"orderList\":[{\"name\":\"珍珠奶茶\",\"amount\":1,\"ice\":\"少冰\",\"sugar\":\"半糖\",\"size\":\"Large\",\"padding\":\"珍珠\"}]}",id)
        fmt.Println("Create Response:>",s )
        var jsonStr = []byte(s);
        url := "http://172.22.20.97:80/api/order/create"
        return callAPI(url,jsonStr);
}

func  listOrder(time int64 ) (response string){
        s := fmt.Sprintf("{\"time\":%d}",time)
        var jsonStr = []byte(s);
        url := "http://172.22.20.97:80/api/order/list"
        return callAPI(url,jsonStr);
}

func lastOrder(id string ) (response string){
        s := fmt.Sprintf("{\"userId\":\"%s\"}",id)
        var jsonStr = []byte(s);
        url := "http://172.22.20.97:80/api/order/last"
        return callAPI(url,jsonStr);
}

func isBonus(id string ) (response string){
        s := fmt.Sprintf("{\"userId\":\"%s\"}",id)
        var jsonStr = []byte(s);
        url := "http://172.22.20.97:80/api/order/isSpecialBonus"
        return callAPI(url,jsonStr);
}
func setBonus(id string ) (response string){
        s := fmt.Sprintf("{\"userId\":\"%s\"}",id)
        var jsonStr = []byte(s);
        url := "http://172.22.20.97:80/api/order/setSpecialBonus"
        return callAPI(url,jsonStr);
}

func listProducts() (response string){
        s := fmt.Sprintf("{}")
        var jsonStr = []byte(s);
        url := "http://172.22.20.97:80/api/product/list"
        return callAPI(url,jsonStr);
}

func updateProduct(productId string, available bool) (response string){
        s := fmt.Sprintf("{\"productId\":\"%s\",\"available\":%t}",productId, available)
        var jsonStr = []byte(s);
        url := "http://172.22.20.97:80/api/product/update"
        return callAPI(url,jsonStr);
}





func frsLogin(username string,password string) (response string){
        s := fmt.Sprintf("{\"username\":\"%s\",\"password\":\"%s\"}",  username,password)
        //  fmt.Println("Create Response:>",s )
        var jsonStr = []byte(s);
        url := "http://172.22.27.167:80/users/login"
        return callAPI(url,jsonStr);
}


func frsCreateUser(sessionId string, id string,file string) (response string){
        f, _ := os.Open(file)
        reader := bufio.NewReader(f)
        content, _ := ioutil.ReadAll(reader)
        encoded := base64.StdEncoding.EncodeToString(content)
        s := fmt.Sprintf("{\"session_id\": \"%s\",\"person_info\": {\"fullname\": \"name\",\"employeeno\": \"%s\",\"cardno\": \"%s\",\"group_list\":[],\"department_list\":[]},\"image\":\"%s\"}",
            sessionId,id,id,encoded)
      //    fmt.Println("Create Response:>",s )
        var jsonStr = []byte(s);
        url := "http://172.22.27.167:80/frs/cgi/createperson"
        return callAPI(url,jsonStr);
}




func frsVerify(sessionId string,file string) (response string){
        f, _ := os.Open(file)
        reader := bufio.NewReader(f)
        content, _ := ioutil.ReadAll(reader)
        encoded := base64.StdEncoding.EncodeToString(content)
        s := fmt.Sprintf("{\"session_id\": \"%s\",\"target_score\" : 0.8, \"request_client\" : \"Test1\",  \"action_enable\": 0,\"source_id\" : \"src1\",\"location\":\"loc1\",\"image\":\"%s\"}",
            sessionId,encoded)
      //    fmt.Println("Create Response:>",s )
        var jsonStr = []byte(s);
        url := "http://172.22.27.167:80/frs/cgi/verifyface"
        return callAPI(url,jsonStr);
}


func updateImage(id string, file string) (response string){
  f, _ := os.Open(file)
  reader := bufio.NewReader(f)
  content, _ := ioutil.ReadAll(reader)
  encoded := base64.StdEncoding.EncodeToString(content)
//  fmt.Println("ENCODED: " + encoded)
  s := fmt.Sprintf("{\"id\":\"%s\",\"image\":\"%s\"}",id,encoded)
  //fmt.Println("Create Response:>",s )
  var jsonStr = []byte(s);
  url := "http://172.22.20.97:80/api/user/updateImage"
  return callAPI(url,jsonStr);
}

func verifyImage( threshold float64, max int, file string) (response string){
  f, _ := os.Open(file)
  reader := bufio.NewReader(f)
  content, _ := ioutil.ReadAll(reader)
  encoded := base64.StdEncoding.EncodeToString(content)
//  fmt.Println("ENCODED: " + encoded)
  s := fmt.Sprintf("{\"threshold\":%f, \"max\":%d, \"image\":\"%s\"}",threshold,max,encoded)
//  fmt.Println("Create Response:>",s )
  var jsonStr = []byte(s);
  url := "http://172.22.20.97:80/api/fr/verification"
  return callAPI(url,jsonStr);
}
func main() {

  //  var response = verifyImage(0.5,1,"./lebron2.jpg")
  //var updateImageRes GeneralResponse
  //  json.Unmarshal([]byte(response), &updateImageRes)
  //fmt.Printf("Umpage Image User Code : %d Message: %s", updateImageRes.Code, updateImageRes.Message)
  //createOrder("5e02c2fc9eff623e2c9f85c1")
  //listOrder(1)
  //lastOrder("5e0073f41cce9e1ba0b401e1")
  //setBonus("5e02c2fc9eff623e2c9f85c1")
  //isBonus("5e02c2fc9eff623e2c9f85c1")

  csvfile, err := os.Create("output.csv")
  if err != nil {
      log.Fatalf("failed creating file: %s", err)
  }
  csvwriter := csv.NewWriter(csvfile)

  var response  = listUser("")
  var listUserRes ListUserResponse
  json.Unmarshal([]byte(response), &listUserRes)
  head  := []string{"Company", "FirstName", "LastName","Title","Email","Checked","CheckType"}
  _ = csvwriter.Write(head)
  for _, s := range listUserRes.UserList {
    newRow := make([]string, 7)
    newRow[0] = s.Company
    newRow[1] = s.FirstName
    newRow[2] = s.LastName
    if s.Title != ""{
      newRow[3] = s.Title
    }else{
      newRow[3] = "N/A"
    }
    if s.Email != ""{
      newRow[4] = s.Email
    }else{
      newRow[4] = "N/A"
    }

    if s.Registered ||  s.Email != ""{
        newRow[5] = "YES"
        if s.CounterRegistered  {
          newRow[6] = "Counter"
        }else{
          newRow[6] = "Face"
        }
    }else{
      newRow[5] = "NO"
      newRow[6] = "  "
    }

    _ = csvwriter.Write(newRow)
  //  deleteUser(s.ID)
    fmt.Println(s)
   }
   csvwriter.Flush()
    //updateBoothCheck("5ddc97509eff62678c6c0cf1" ,"BoothA", true)
    //updateBoothCheck("5ddc97509eff62678c6c0cf1" ,"BoothB", true)
    //updateBoothCheck("5ddc97509eff62678c6c0cf1" ,"BoothC", true)
    //updateBoothCheck("5ddc97509eff62678c6c0cf1" ,"BoothA", false)

//    findUser("5ddc97509eff62678c6c0cf1DDFDFDFDFDfddF")
  //  findUserFace("5ddc97509eff62678c6c0cf1")
  //  response  = verifyImage(0.7,1,"./photo3.jpg")
    //var updateImageRes GeneralResponse
    //json.Unmarshal([]byte(response), &updateImageRes)
    //fmt.Printf("Umpage Image User Code : %d Message: %s", updateImageRes.Code, updateImageRes.Message)
//"5ddc97509eff62678c6c0cf1"

    /*
    response  = createUser("Lebron", "James", "NBA", "MVP","lebron.james@gmail.com","09876543","ext1","ext2" );
    fmt.Println("Create Response:>", response)
    var createRes CreateUserResponse
    json.Unmarshal([]byte(response), &createRes)
    fmt.Printf("Create User Code : %d ID: %s", createRes.Code, createRes.ID)
    if createRes.Code != 0 {
      fmt.Printf("Fail Code : %d Message : %s", createRes.Code, createRes.Message)
      return
    }

      response  = updateImage(createRes.ID,"./lebron1.jpg")
      json.Unmarshal([]byte(response), &updateImageRes)
      fmt.Printf("Umpage Image User Code : %d Message: %s", updateImageRes.Code, updateImageRes.Message)
      */
    //  5e02c2fc9eff623e2c9f85c1"
      //  updateUser("5de4af2f1cce9e045cf42154","YYY@gmail","Manager",false,false)
      //  response  = listUser("")

        /*
    {
      response  = createUser("YuChio", "Sha", "Cloud","test2@gmail.com","09876543","ext1","ext2" );
      fmt.Println("Create Response:>", response)
      var createRes CreateUserResponse
      json.Unmarshal([]byte(response), &createRes)
      fmt.Printf("Create User Code : %d ID: %s", createRes.Code, createRes.ID)
      if createRes.Code != 0 {
        fmt.Printf("Fail Code : %d Message : %s", createRes.Code, createRes.Message)
        return
      }

      response  = updateImage(createRes.ID,"./photo2.jpg")
      var updateImageRes GeneralResponse
      json.Unmarshal([]byte(response), &updateImageRes)
      fmt.Printf("Umpage Image User Code : %d Message: %s", updateImageRes.Code, updateImageRes.Message)
    }
   */
    /*
    response  =  frsLogin("Admin","123456")
    var frsLoginRes FRSLoginResponse
    json.Unmarshal([]byte(response), &frsLoginRes)
    fmt.Printf("Login Sesson ID : %s", frsLoginRes.SessionID)

    response  =  frsCreateUser(frsLoginRes.SessionID, createRes.ID,"./photo1.jpg")
    var frsCreateUserRes FRSCreateUserResponse
    json.Unmarshal([]byte(response), &frsCreateUserRes)
      fmt.Printf("Create User Message : %s", frsCreateUserRes.Message)
    if frsCreateUserRes.Message == "ok" {
      response  =  frsVerify(frsLoginRes.SessionID,"./photo1.jpg")
      var frsVerifyResponse FRSVerifyResponse
      json.Unmarshal([]byte(response), &frsVerifyResponse)
      if frsVerifyResponse.Message == "ok" {
          fmt.Printf("Verify Face ID : %s", frsVerifyResponse.VerifyFaceID)
      }
    }
    */


    //response  = deleteUser(createRes.ID)
  ///  var deleateRes GeneralResponse
    //json.Unmarshal([]byte(response), &deleateRes)
  //  fmt.Printf("Delete User Code : %d Message: %s", deleateRes.Code, deleateRes.Message)

    //updateImage("123","./photo1.jpg" )
    /*.code
    url := "http://172.22.20.97:80/api/user/create"
    fmt.Println("URL:>", url)

    var jsonStr = []byte(`{"firstname":"A","lastname":"B","company":"Advantch"}`)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    //req.Header.Set("X-Custom-Header", "myvalue")
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
    */
    //listProducts()
    //updateProduct("1",true)
}
