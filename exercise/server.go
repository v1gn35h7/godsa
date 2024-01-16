package exercise

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type user struct {
	Id   string
	Name string
}

type myHandler struct {
}

func (mh myHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	if req.RequestURI == "/helo" {
		res.Write([]byte("hi"))

	} else {
		res.Write([]byte("Vicky Server- V 0.0.0.0-beta"))

	}

}

func StartServer(port int) {

	s := fmt.Sprintf("%s:%d", "", port)

	ser := http.NewServeMux()

	ser.HandleFunc("/helo", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hi"))

	})

	ser.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {

		usr := user{
			Id:   "3131313",
			Name: "Vicky",
		}

		data, _ := json.Marshal(usr)

		w.Write(data)

	})

	log.Fatal(http.ListenAndServe(s, ser))

}
