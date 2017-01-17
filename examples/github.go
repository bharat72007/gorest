package main

import (
	"fmt"
	"github.com/gorest"
)

type Post struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserId string `json:"userid"`
	Id     string `json:"id"`
}

func main() {
	var post Post
	rest := gorest.New()
	request, _ := rest.Header(gorest.ContentType, gorest.JsonContentType).Base("https://jsonplaceholder.typicode.com").Path("posts").Path("1").Get().Request()
	fmt.Println(request)
	resp, _ := rest.Send(request, "", "")
	fmt.Println(rest.ResponseBodyString(resp))
	rest.Recieve(resp, &post)
	fmt.Println(post)

	/*rest := gorest.New()
	querystr := map[string]string{"postId": "1"}
	request, _ := rest.Header(gorest.ContentType, gorest.JsonContentType).Base("https://jsonplaceholder.typicode.com").Path("comment").Query(querystr).Get().Request()
	fmt.Println(request)
	resp, _ := rest.Send(request, "", "")
	fmt.Println(rest.ResponseBodyString(resp))
	comments := new([]Comment)
	rest.Recieve(resp, &comments)
	fmt.Println(comments)
	*/ //POST Method
	postr := Post{Title: "foo",
		Body:   "bar",
		UserId: "1"}
	rest2 := gorest.New()
	request2, _ := rest2.Header(gorest.ContentType, gorest.JsonContentType).Base("http://jsonplaceholder.typicode.com/").Path("posts").Post(postr).Request()
	fmt.Println(request2)
	resp2, _ := rest2.Send(request2, "", "")
	fmt.Println(rest.ResponseBodyString(resp2))
}

type Comment struct {
	PostId string `json:"postId"`
	Id     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}
