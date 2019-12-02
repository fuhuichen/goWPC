package frs

import (
  "encoding/json"
  "fmt"
	"net/http"
  "bytes"
  "io/ioutil"
  "math/rand"
)

type FrsClient struct{
  IP string
}

type FRSLoginResponse struct {
	SessionID   string `json:"sessionId"`
}
type FRSCreateUserResponse struct {
	Message   string `json:"message"`
  PersonID   string `json:"person_id"`
}
type FRSVerifyResponse struct {
	Message   string `json:"message"`
  VerifyFaceID   string `json:"verify_face_id"`
}

type FRSWSResponse struct {
  VerifyFaceID   string `json:"verify_face_id"`
  PersonID   string `json:"person_id"`
}



func  callAPI(url string,jsonStr []byte)(response string){
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

func  (client *FrsClient) FrsLogin(username string,password string) (response string){
        s := fmt.Sprintf("{\"username\":\"%s\",\"password\":\"%s\"}",  username,password)
        //  fmt.Println("Create Response:>",s )
        var jsonStr = []byte(s);
        url :=   fmt.Sprintf("http://%s/users/login",client.IP )
        var res =  callAPI(url,jsonStr);
        var frsLoginRes FRSLoginResponse
        json.Unmarshal([]byte(string(res)), &frsLoginRes)
        response = frsLoginRes.SessionID
        return response
}


func  (client *FrsClient) FrsCreateUser(sessionId string, id string,encoded string) (response string){
        s := fmt.Sprintf("{\"session_id\": \"%s\",\"person_info\": {\"fullname\": \"name\",\"employeeno\": \"%s\",\"cardno\": \"%s\",\"group_list\":[],\"department_list\":[]},\"image\":\"%s\"}",
            sessionId,id,id,encoded)
      //    fmt.Println("Create Response:>",s )
        var jsonStr = []byte(s);
        url :=   fmt.Sprintf("http://%s/frs/cgi/createperson",client.IP )
        var res =  callAPI(url,jsonStr);
        fmt.Println("res :", string(res))
        var frsCreateUserRes FRSCreateUserResponse
        json.Unmarshal([]byte(string(res)), &frsCreateUserRes)
        response = ""

        if frsCreateUserRes.Message == "ok" {
            fmt.Println("res :", frsCreateUserRes.PersonID)
          response =  frsCreateUserRes.PersonID
        }
        return response
}


func  (client *FrsClient) FrsVerify(sessionId string,encoded string, rate float64) (response string){
      src := fmt.Sprintf("src%d",rand.Intn(1000))
      loc := fmt.Sprintf("loc%d",rand.Intn(1000))
        s := fmt.Sprintf("{\"session_id\": \"%s\",\"target_score\" : %f, \"request_client\" : \"Test1\",  \"action_enable\": 0,\"source_id\" : \"%s\",\"location\":\"%s\",\"image\":\"%s\"}",
            sessionId,rate,src,loc,encoded)
      //    fmt.Println("Create Response:>",s )
        var jsonStr = []byte(s);
        url :=   fmt.Sprintf("http://%s/frs/cgi/verifyface",client.IP )
        var res =  callAPI(url,jsonStr);
        var frsVerifyResponse FRSVerifyResponse
        json.Unmarshal([]byte(string(res)), &frsVerifyResponse)
        if frsVerifyResponse.Message == "ok" {
          return frsVerifyResponse.VerifyFaceID
        }
        return ""
}
/*
func main() {

    var response  = createUser("YoungHwa", "Song", "Advantech","test@gmail.com","09876543","ext1","ext2" );
    fmt.Println("Create Response:>", response)
    var createRes CreateUserResponse
    json.Unmarshal([]byte(response), &createRes)
    fmt.Printf("Create User Code : %d ID: %s", createRes.Code, createRes.ID)
    if createRes.Code != 0 {
      fmt.Printf("Fail Code : %d Message : %s", createRes.Code, createRes.Message)
      return
    }

    response  = updateImage(createRes.ID,"./photo1.jpg")
    var updateImageRes GeneralResponse
    json.Unmarshal([]byte(response), &updateImageRes)
    fmt.Printf("Umpage Image User Code : %d Message: %s", updateImageRes.Code, updateImageRes.Message)

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



  //  response  = deleteUser(createRes.ID)
  //  var deleateRes GeneralResponse
  //  json.Unmarshal([]byte(response), &deleateRes)
  //  fmt.Printf("Delete User Code : %d Message: %s", deleateRes.Code, deleateRes.Message)

    //updateImage("123","./photo1.jpg" )
}
*/
