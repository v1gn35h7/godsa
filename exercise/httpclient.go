package exercise

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type response interface{}

func HttpClient() {
	// https://hub.dummyapis.com/employee?noofRecords=10&idStarts=1001
	res, err := http.Get("https://dummyjson.com/comments")

	if err != nil {
		panic(err)
	}

	var r response

	err = json.NewDecoder(res.Body).Decode(&r)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(r)

}
