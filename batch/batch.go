package main

import (
    "encoding/json"
    "encoding/csv"
    "os"
    "fmt"
    "io"
    "net/http"
    "bytes"
    "io/ioutil"
    "bufio"
    "encoding/base64"
    "log"
)

type GeneralResponse struct {
	Code     int `json:"code"`
	Message  string `json:"message"`
}

type CreateUserResponse struct {
	Code     int `json:"code"`
	Message  string `json:"message"`
  ID      string `json:"id"`
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

func createUser(firstname string, lastname string, company string, title string,
        email string ,mobile string ,extend1 string,extend2 string ) (response string){
        s := fmt.Sprintf("{\"firstname\":\"%s\",\"lastname\":\"%s\",\"company\":\"%s\",\"title\":\"%s\",\"email\":\"%s\",\"mobile\":\"%s\", \"extend1\":\"%s\",\"extend2\":\"%s\"}",
                  firstname,lastname,company,title, email,mobile,extend1,extend2)
          fmt.Println("Create Response:>",s )
        var jsonStr = []byte(s);
        url := "http://127.0.0.1:9741/api/user/create"
        return callAPI(url,jsonStr);
}

func updateImage(id string, file string) (response string){
  fmt.Println("Updagte image ID Response:>", id )
  fmt.Println("FileName Response:>", file )
  f, _ := os.Open(file)
  reader := bufio.NewReader(f)
  content, _ := ioutil.ReadAll(reader)
  encoded := base64.StdEncoding.EncodeToString(content)
//  fmt.Println("ENCODED: " + encoded)
  s := fmt.Sprintf("{\"id\":\"%s\",\"image\":\"%s\"}",id,encoded)
  //fmt.Println("Create Response:>",s )
  var jsonStr = []byte(s);
  url := "http://127.0.0.1:9741/api/user/updateImage"
  return callAPI(url,jsonStr);
}

func ReadCsvFile(filePath string)  {
    // Load a csv file.
    f, _ := os.Open(filePath)

    // Create a new reader.
    r := csv.NewReader(f)
    csvfile, err := os.Create("output.csv")

	  if err != nil {
		    log.Fatalf("failed creating file: %s", err)
	  }
	  csvwriter := csv.NewWriter(csvfile)

  //	for _, row := range rows {
	//	_ = csvwriter.Write(row)
	//}
    for {
        record, err := r.Read()
        // Stop at EOF.
        if err == io.EOF {
            break
        }

        if err != nil {
            panic(err)
        }
        // Display record.
        // ... Display record length.
        // ... Display all individual elements of the slice.
        fmt.Println(record)

        //fmt.Println(len(record))
        if len(record)>15 && record[1]!="First Name" {
          var response  = createUser(record[1], record[2],record[3],record[4],record[0],"","","" );
          var createRes CreateUserResponse
          json.Unmarshal([]byte(response), &createRes)
          fmt.Printf("Create User Code : %d ID: %s", createRes.Code, createRes.ID)
          if record[5]!= ""  {
            response  = updateImage(createRes.ID,"./data/ProfilePicture_" +  record[5] + ".jpg")
            var updateImageRes GeneralResponse
            json.Unmarshal([]byte(response), &updateImageRes)
            fmt.Printf("Umpage Image User Code : %d Message: %s", updateImageRes.Code, updateImageRes.Message)
          }

        }
        if record[1]!="First Name" {
          newRow := append(record,"XXXXXXXX")
          _ = csvwriter.Write(newRow)
        }else{
          newRow := append(record,"QRCode")
          _ = csvwriter.Write(newRow)
        }


      //for value := range record {
      //      fmt.Printf("  %v\n", record[value])
      //  }
    }
    csvwriter.Flush()
    csvfile.Close()
}



func main() {
    ReadCsvFile("./data/Member_List-Test.csv")
}
