package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	ip, err := getIP()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Ip:", ip)
}

func getIP() (string, error) {
	resp, err := http.Get("https://httpbin.org/ip")

	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("bad status: %s", resp.Status)
	}

	//fmt.Println("content type:", resp.Header.Get("content-type"))
	//fmt.Println("body", resp.Body)
	//io.Copy(os.Stdout, resp.Body)

	var ir IPResponse
	dec := json.NewDecoder(resp.Body)

	if err := dec.Decode(&ir); err != nil {
		//log.Fatal(err)
		return "", err
	}
	//fmt.Println("Origin:", ir.Origin)
	//fmt.Println("ip:", ir.IP)
	return ir.Origin, nil
}

type IPResponse struct {
	Origin string //must be capital letter
	//IP     string `json:"origin"` // field tag  these are backticks not single quotes
}

/*JSON <-> Go
object <-> map, struct
array <-> slice
string <-> string
array <-> slice
number <-> float64
boolean <-> bool
null <-> nil

*/
