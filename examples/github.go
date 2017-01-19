package main

import (
	"fmt"
	gorest "github.com/gorest"
)

//CREATED A Post Structure
type Post struct {
	Id     int    `json:"id"`
	UserId int    `json:"userid"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

const (
	examplebaseurl = "https://jsonplaceholder.typicode.com/"
)

func main() {
	//GET Request with URI Params and Query String
	//Response Contains List of JSON Objects
	rest1 := gorest.New()
	querystr1 := map[string]string{"postId": "1"}
	request1, _ := rest1.WithHeader(gorest.ContentType, gorest.JsonContentType).Base(examplebaseurl).Path("comments").Query(querystr1).Get().Request()
	response1, _ := rest1.Send(request1)
	var posts []Post
	gorest.Response(response1, &posts, nil)
	//PRINT Utility
	PrintUtil(request1.URL, posts)

	//GET Request with URI Params
	//Response Contains Single JSON Object
	var post Post
	rest2 := gorest.New()
	request2, _ := rest2.WithHeader(gorest.ContentType, gorest.JsonContentType).Base(examplebaseurl).Path("posts").Path("3").Get().Request()
	response2, _ := rest2.Send(request2)
	gorest.Response(response2, &post, nil)
	//PRINT Utility
	PrintUtil(request2.URL, post)

	//POST Request to Create new Post for a Specific User
	//Response it will create new Pos for that User.
	newpost := Post{
		Title:  "New Post Added",
		UserId: 1,
		Body:   "Body for a new Post is not that good",
	}
	rest3 := gorest.New()
	request3, _ := rest3.WithHeader(gorest.ContentType, gorest.JsonContentType).Base(examplebaseurl).Path("posts").Post(newpost).Request()
	response3, _ := rest3.Send(request3)
	gorest.Response(response3, &post, nil)
	//PRINT Utility
	PrintUtil(request3.URL, post)

	//PUT Request to Update already exisitng id Post
	//Response it will update all the fields except "id".
	updatepost := Post{
		Title:  "Update Post ",
		UserId: 10000,
		Body:   "Body for a Updated Post",
	}
	rest4 := gorest.New()
	request4, _ := rest4.WithHeader(gorest.ContentType, gorest.JsonContentType).Base(examplebaseurl).Path("posts").Path("1").Put(updatepost).Request()
	response4, _ := rest4.Send(request4)
	gorest.Response(response4, &post, nil)
	//PRINT Utility
	PrintUtil(request4.URL, post)

	//PATCH Request to Update userId exisitng id = "1" Post
	//Response it will update userId the field.
	patchpost := Post{
		UserId: 50000,
	}
	rest5 := gorest.New()
	request5, _ := rest5.WithHeader(gorest.ContentType, gorest.JsonContentType).Base(examplebaseurl).Path("posts").Path("1").Patch(patchpost).Request()
	response5, _ := rest5.Send(request5)
	gorest.Response(response5, &post, nil)
	//PRINT Utility
	PrintUtil(request5.URL, post)

	//DELETE Request to delete post
	//Response it will delete post
	rest6 := gorest.New()
	request6, _ := rest6.WithHeader(gorest.ContentType, gorest.JsonContentType).Base(examplebaseurl).Path("posts").Path("1").Delete().Request()
	response6, _ := rest6.Send(request6)
	gorest.Response(response6, &post, nil)
	//PRINT Utility
	PrintUtil(request6.URL, response6.Status)
}

func PrintUtil(request, posts interface{}) {
	fmt.Printf("Requested URL: %v \n", request)
	fmt.Printf("Response Body: %v \n", posts)
	fmt.Println("#########################################")
}
