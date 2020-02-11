package main

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var awvsurl = "https://xx.xx.xx.xx:13443"
var apikey = "xxxxx"

type AwvsCount struct {
	ScansRunningCount int `json:"scans_running_count"`
}

func GetInfo() (Can bool){
	req, _ := http.NewRequest("GET",awvsurl + "/api/v1/me/stats",nil)
	req.Header.Set("X-Auth",apikey)
	req.Header.Set("Content-Type","application/json")
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	resp, err := (&http.Client{Transport: tr}).Do(req)
	if err != nil{
		fmt.Println(err)
	}
	defer resp.Body.Close()
	respByte, _ := ioutil.ReadAll(resp.Body)
	var s AwvsCount
	json.Unmarshal([]byte(string(respByte)), &s)
	if s.ScansRunningCount < 2{
		return true
	}else {
		return false
	}
}
func AddTarget(url string) (target_id string){
	type TargetID struct {
		TargetID string `json:"target_id"`
	}
	post := "{\"address\":\"" + url + "\",\"description\":\"" + url + "\",\"criticality\":\"" + "10" + "\"}"
	var jsonStr = []byte(post)
	req, _ := http.NewRequest("POST", awvsurl + "/api/v1/targets",bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Auth",apikey)
	req.Header.Set("Content-Type","application/json")
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	resp, err := (&http.Client{Transport: tr}).Do(req)
	if err != nil{
		fmt.Println(err)
	}
	defer resp.Body.Close()
	respByte, _ := ioutil.ReadAll(resp.Body)
	var s TargetID
	json.Unmarshal([]byte(string(respByte)), &s)
	return s.TargetID
}

func ScanTarget(url string) (success bool){
	post := "{\"target_id\":\"" + AddTarget(url) + "\",\"profile_id\":\"" + "11111111-1111-1111-1111-111111111111" + "\",\"schedule\":{\"disable\":false,\"start_date\":null,\"time_sensitive\":false}}"
	var jsonStr = []byte(post)
	req, _ := http.NewRequest("POST",awvsurl + "/api/v1/scans",bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Auth",apikey)
	req.Header.Set("Content-Type","application/json")
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	resp, err := (&http.Client{Transport: tr}).Do(req)
	if err != nil{
		fmt.Println(err)
	}
	defer resp.Body.Close()
	fmt.Println(string(jsonStr))
	if resp.StatusCode == 201{
		return true
	}else {
		return false
	}

}


func main() {
	file, err := os.Open("./target.txt")
	if err != nil{
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
RET:
		if GetInfo() {
			if ScanTarget(scanner.Text()){
				fmt.Println(scanner.Text(),"success")
			}else {
				fmt.Println(scanner.Text(),"fail")
			}
		}else {
			fmt.Println("scan is full[*]waiting")
			time.Sleep(60 * 1e9)
			goto RET
		}
	}
}
