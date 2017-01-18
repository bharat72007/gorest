package main

import (
	"fmt"
	gorest "github.com/gorest"
)

const (
	BASEURL         = "https://jsonplaceholder.typicode.com"
	ContentType     = "Content-Type"
	JSONContentType = "application/json"
)

type Post struct {
	Id     int    `json:"id"`
	UserId int    `json:"userid"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	postsucess1 := new(Post)
	client1 := gorest.New()
	request1, _ := client1.Base(BASEURL).Header(ContentType, JSONContentType).Path("posts").Path("1").Get().Request()

	err := client1.ResponseStruct(request1, postsucess1, nil)
	fmt.Print(err)
	fmt.Println(postsucess1)
}
