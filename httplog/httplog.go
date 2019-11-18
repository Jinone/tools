package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"time"
)

func Work(url string,ip string, header string, host string, method string) (_infos []string, exist bool, num int64) {
	db, err := sql.Open("sqlite3", "./log.db")
	var infos []string
	if err != nil {
		log.Fatal(err)
	}
	if method == "insert" {
		sql := fmt.Sprintf("INSERT INTO logs (url,ua,host,times,ip) values ('%s','%s','%s','%d','%s')", url, header, host, time.Now().Unix(),ip)
		_, err := db.Exec(sql)
		if err != nil {
			fmt.Println("insert error")
			log.Fatalln(err)
		}}else if method == "query"{
		rows, err := db.Query("select * from logs")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			var id int
			var url string
			var ua string
			var host string
			var times string
			var ip string
			rows.Scan(&id,&url,&ua,&host,&times,&ip)
			info := fmt.Sprintf("\nID: %d\nreq_url: %s\nUserAgent: %s\nHost: %s\ntimes: %s\nIp: %s\n",id,url,ua,host,times,ip)
			infos = append(infos,info)
		}
		return infos, true, 1
	}else if method == "del" {
		row, err := db.Exec("delete from logs")
		if err != nil {
			fmt.Println("del error")
			log.Fatalln(err)
			id, err := row.RowsAffected()
			if err != nil {
				log.Fatalln(err)
			}
			return _infos,true, id
		}
	}else {
		sql := fmt.Sprintf("select host from logs where host = '%s'",host)
		row, err := db.Query(sql)
		if err != nil {
			log.Fatal(err)
		}
		defer row.Close()
			if row.Next(){
				return _infos, true, 1
			}else {
				return _infos, false, 1
		}
	}
	return _infos, true, 1
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"good luck!")
	Work(r.RequestURI,r.RemoteAddr,r.UserAgent(),r.Host,"insert")
	fmt.Println(r.RequestURI)
}

func Admin(w http.ResponseWriter, r *http.Request) {
	a, _, _ := Work("1","1","1", "1", "query")
	for _, y := range a {
		fmt.Fprintf(w,y)
	}
}

func Del(w http.ResponseWriter, r *http.Request) {
	_, _, c := Work("1","1","1", "1", "del")
		fmt.Fprintf(w,"del %d success",c)
}

func Api(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	_, b, _ := Work("1","1", "1",r.Form["url"][0], "api")
	if b{
		fmt.Fprint(w,"true")
	}else{
		fmt.Fprint(w,"false")
	}

}

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/admin", Admin)
	http.HandleFunc("/del", Del)
	http.HandleFunc("/api", Api)
	fmt.Println("start server")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
