package main

import (
	"fmt"
	g "github.com/gorest"
)

type Post struct {
	Id     int    `json:"id"`
	UserId int    `json:"userid"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	rest := g.New()
	querystr := map[string]string{"postId": "1"}
	request, _ := rest.AddHeader(g.ContentType, g.JsonContentType).BasePath("https://jsonplaceholder.typicode.com/").URIParam("comments").SetQuery(querystr).Get().Request()
	fmt.Println(request)
	resp, _ := rest.Send(request, "", "")
	var poss []Post
	fmt.Println(rest.ResponseStructure(resp, &poss))

	//POST Method
	/*postr := Post{Title: "foo",
		Body:   "bar",
		UserId: 1}
	rest2 := g.New()
	request2, _ := rest2.AddHeader(g.ContentType, g.JsonContentType).BasePath("http://jsonplaceholder.typicode.com/").URIParam("posts").Post(postr).Request()
	fmt.Println(request2)
	resp2, _ := rest2.Send(request2, "", "")
	fmt.Println(rest.ResponseStructure(resp2, ""))
	*/
}
