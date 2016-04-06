package main
import (
	"os"
	"fmt"
	"encoding/json"
//	"io/ioutil"
)

var Settings struct {
	AccessKey string `json:"accessKey"`
	SecretKey  string `json:"secretKey"`
	EndPoint  string `json:"endPoint"`
}

func main() {
	configFile, err := os.Open("config.json")
	if err != nil {

		fmt.Println("file not present", err.Error())
	} else {
		fmt.Println("file present")
	}

	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&Settings); err != nil {
		fmt.Println("file not present", err.Error())
	} else {
		fmt.Println("file present")
	}

	fmt.Printf("%v ** %v ** %v **\n", Settings.AccessKey, Settings.EndPoint, Settings.SecretKey)


//	file, e := ioutil.ReadFile("config.json")
//	if e != nil {
//		fmt.Printf("File error: %v\n", e)
//		os.Exit(1)
//	}
//	fmt.Printf("%s\n", string(file))
//
//	//m := new(Dispatch)
//	//var m interface{}
//	var jsontype Settings
//	json.Unmarshal(file, &jsontype)
//	fmt.Printf("Results: %v\n", jsontype)

}

