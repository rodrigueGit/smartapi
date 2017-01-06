/* SmartAPI send command
 */

package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
  "time"
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
  Sent  time.Time
}

var options Options
var address string

func parseOptions() {
	options.InputFile = flag.String("i", "input", "a file")
	flag.Parse()
	options.AddressFile = flag.Args()
}

func readAddress() {
	buf := bytes.NewBuffer(nil)
	f, err := os.Open(options.AddressFile[0])
	if err != nil {
		panic(err)
	}
	io.Copy(buf, f)
	f.Close()
	address = string(buf.Bytes())
}

func readIPAddress() {
  url := "http://127.0.0.1:8080/address"
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

  fmt.Println("response Status:", resp.Status)
  fmt.Println("response Headers:", resp.Header)
  body, _ := ioutil.ReadAll(resp.Body)
  fmt.Println("response Body:", string(body))
}

func main() {
	parseOptions()
	readAddress()
	fmt.Println("Input file: ", *options.InputFile)
	fmt.Println("Address file: ", options.AddressFile[0])
	fmt.Println("Address: ", address)
}
