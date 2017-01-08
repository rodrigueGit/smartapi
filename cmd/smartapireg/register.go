package main

import (
	"bytes"
	"crypto/rand"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"encoding/json"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"time"
)

//Address structure used to register public key
type Address struct {
	ID      string    `json:"id"`
	IP      string    `json:"ip"`
	Updated time.Time `json:"updated"`
}

//Identity structure used to define a unique ID
type Identity struct {
	ID	string `json:"key"`
	PrivateKey ecdsa.PrivateKey `json:"privateKey"`
	Generated time.Time `json:"generated"`
}

func generateKey(id *Identity) {

	//Generate key pair based on ecdsa
	pubkeyCurve := elliptic.P256()
	key, err := ecdsa.GenerateKey(pubkeyCurve, rand.Reader) // this generates a public & private key pair

 	if err != nil {
 		panic(err)
 	}

	//Hash public key into address ID
	var pubKeyBuffer bytes.Buffer
	pubKeyBuffer.WriteString(key.PublicKey.X.String())
	pubKeyBuffer.WriteString(key.PublicKey.Y.String())
	id.PrivateKey = *key
	sha256Hash := sha256.Sum256(pubKeyBuffer.Bytes())
	fmt.Println("Base for SHA256:", hex.EncodeToString(sha256Hash[:]))

	id.ID = hex.EncodeToString(sha256Hash[:])
	id.Generated = time.Now()
}

func getAddress(address *Address) {
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			address.IP = ipv4.String()
		}
	}
}

func registerAddress(identity *Identity, address *Address) {
	address.ID = identity.ID
	address.Updated = time.Now()
	url := "http://smartapi-155012.appspot.com/address"

	json, err := json.Marshal(address)
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

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

func main() {
	var identity Identity
	var address Address

	generateKey(&identity)
	getAddress(&address)
	registerAddress(&identity, &address)
}
