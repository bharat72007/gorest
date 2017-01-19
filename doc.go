/*
Package gorest is a Go HTTP client library for creating and sending REST API requests.


GoRest encapsulates HTTP Request properties in Rest Client, which simplify the construction
and invokation of REST API's.

Check the examples to learn how to compose a HTTP Request using GORest and How to handle API
responses(including errors and success)

Usage
Use GORest to set BASE URL, HTTP Verbs, Headers, Query Params, URI Prams or Request Payload
and create a HTTP request.
Use a Sling to set path, method, header, query, or body properties and create an
Create Rest Client
rest := gorest.New()
Attach BASE URL
rest.Base("https://api.example.com/")
Path/URI Parameters
Path method can be used to extend the Base URL.
Lets add Path to BASEURL
rest.Base("https://api.example.com/").Path("param1")
We can add as many paths as possible by chaining them one after other.
rest.Base("https://api.example.com/").Path(param1).Path(param2).Path(param3)
So our Request URL looks like
https://api.example.com/{param1}/{param2}/{param3}
User Need not to worry about the Forward slash while constructing URL(except BASEURL)
Headers
Add or Set headers for a requests created by Rest Client
Headers are key value pairs and therefore our WithHeader API consume key value string as paramters.
rest.Base("https://api.example.com/").Path("param1").WithHeader("User-Agent", "REST API Client")
Query
We can attach Query Paramters to our Request URL.
Query Paramters can be attached in one invokation, user just needs to provide all Query parameters
in {string: string} Map
query := map[string]string{"searchkey1" : "value1",
			"searchkey2" : "value2",
			"searchkey3" : "value3",
		}
rest.Base("https://api.example.com/").Path("param1").Query(query)
So our complete request URL will be
https://api.example.com/{param1}?searchkey1=value1&searchkey2=value2&searchkey3=value3
NOTE: Query is not a chaining method like Path, If we chain Query paramters we are directly replacing
with the latest query
rest.Base("https://api.example.com/").Path("param1").Query(query1).Query(query2)
query2 will be part of the request URL and not query1.
HTTP Verbs
GORest Client supports common HTTP Verbs which includes GET/POST/PUT/DELETE/PATCH/HEAD/OPTIONS
Example GET request.
rest.Base("https://api.example.com/").Path("param1").Query(query1).Get()
This way request is attached with GET HTTP Verb.
Similarily For HEAD Request
rest.Base("https://api.example.com/").Path("param1").Head()
Payload
In Request which needs Body in the request, user needs to construct Payload first.
Right now only JSON format Payload is supported.
Lets create an structure which we needs to post
type Post struct {
	UserId int    `json:"userid"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
We will create a Post Payload.
post := POST{
	UserId: "12",
	Title: "any random title",
	Body: "Body content random",
}
So request will become.
rest.Base("https://api.example.com/").Path("param1").Post(post)
Similarily for Partial Updates we can create partial Payload and invoke Patch Verb.
partialpost := POST{
	Body: "Body content random",
}
Only Body will be updated for id : 1
rest.Base("https://api.example.com/").Path("1").Patch(partialpost)
Request
From rest client after composing HTTP request properties, we can create http.Request from it.
Request() method is used for it.
rest.Base("https://api.example.com/").Path("1").Patch(partialpost).Request() //Will create http.Request out of this Entity.
Other examples:
rest.Base("https://api.example.com/").Path("param1").Query(query).Get().Request()
rest.Base("https://api.example.com/").Path("param1").Put(partialpayload).Request()
rest.Base("https://api.example.com/").Path("param1").Post(payload).Request()
rest.Base("https://api.example.com/").Query(query).Head().Request()
Response
HTTP Request can be send to Sever and based on that Response can be collected at various structures.
Success Structure and Error Structure, If Error occurs then error message will be pumped in to Error Structure
provided by End User, similarly for Success Structure.
From the Get Request if we need to collect the JSON Response.
type Post struct {
	UserId int    `json:"userid"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	Id int    `json:"id"`
}

type PostErr struct{
	Code int `json:"code"`
	Cause string `json:"cause"`
	CauseId string `json:"causeId"`
	Errurl string `json:"errurl"`
}
request, _ := rest.Base("https://api.example.com/").Path("param1").Query(query).Get().Request()
response, _ := rest.Send(request)
Here user have collected the Response in the form of http.Response.
To populate Success and Error structures
var post Post
var errpost PostErr
Method Response which is not tied to REST Client can be used.
Response(http.Response, interface{}, interface{})
err := Response(response,post,errpost)
From above invokation, post and errpost structures will be populated(based on success or error)
ResponseString
If user just want to look for response Body as a string, then using ResponseBodyString can be used
content := ResponseBodyString(request)
*/

package gorest
