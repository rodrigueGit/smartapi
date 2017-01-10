/* SmartAPI send command
 */

package main

import (
	"bytes"
	"flag"
	"fmt"
  "time"
  "net/http"
  "encoding/json"
  "io/ioutil"
)

//Options structure
type Options struct {
	InputFile   *string
	AddressFile []string
}

//Delivery structure
type Delivery struct {
	FileName string
	Data     []byte
  Sent     time.Time
}

// Address structure to be used in the address registry
type Address struct {
	ID     string    `json:"id"`
	IP      string    `json:"ip"`
	Updated time.Time `json:"updated"`
}

var options Options
var address string
var addressResp Address
var delivery Delivery

func parseOptions() {
	options.InputFile = flag.String("i", "input", "a file")
	flag.Parse()
	address = flag.Args()[0]
}

func readIPAddress() {
  urlBuffer := bytes.NewBufferString("http://smartapi-155012.appspot.com/address/")
  urlBuffer.WriteString(address)
  reqStr := urlBuffer.String()
  fmt.Printf("Address length: %d", len(address))
  fmt.Println("URL:", reqStr)

  req, err := http.NewRequest("GET", reqStr, nil)
  if err!=nil {
    panic(err)
  }
  client := &http.Client{}
  req.Header.Set("Content-Type", "application/json; charset=UTF-8")
  resp, err := client.Do(req)
  if err!=nil {
    panic(err)
  }

  defer resp.Body.Close()
  err = json.NewDecoder(resp.Body).Decode(&addressResp)
	if err != nil {
		panic(err)
	}

  fmt.Println("response Status:", resp.Status)
  fmt.Println("response Headers:", resp.Header)
  fmt.Println("response IP:", addressResp.IP)
}

func sendFile() {
  delivery.FileName = *options.InputFile
  delivery.Sent = time.Now()
  data, err := ioutil.ReadFile(delivery.FileName)
  if err!=nil {
    panic(err)
  }
  delivery.Data = data

  url := "http://smartapirec.appspot.com/file"

  json, err := json.Marshal(delivery)
  if err != nil {
    panic(err)
  }

  req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))
  if err != nil {
    panic(err)
  }
  req.Header.Set("Content-Type", "application/json")

  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
    panic(err)
  }
  defer resp.Body.Close()
}

func main() {
	parseOptions()
  readIPAddress()
  sendFile()
}
