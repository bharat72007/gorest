package main

import (
	"fmt"
	g "github.com/gorest"
)

type Post struct {
	title  string
	body   string
	userId int
}

func main() {
	rest := g.New()
	querystr := map[string]string{"postId": "1"}
	request, _ := rest.Header(g.ContentType, g.JsonContentType).Base("https://jsonplaceholder.typicode.com/").Path("comments").Query(querystr).Get().Request()
	fmt.Println(request)
	resp, _ := rest.Send(request, "", "")
	fmt.Println(rest.ResponseBodyString(resp))

	//POST Method
	postr := Post{title: "foo",
		body:   "bar",
		userId: 1}
	rest2 := g.New()
	request2, _ := rest2.Header(g.ContentType, g.JsonContentType).Base("http://jsonplaceholder.typicode.com/").Path("posts").Post(postr).Request()
	fmt.Println(request2)
	resp2, _ := rest2.Send(request2, "", "")
	fmt.Println(rest.ResponseBodyString(resp2))
}
