# Gorest 
<img align="right" src="resource/img/groundhog_rest.jpeg"> 

GoRest is a HTTP client library written in Go for composing and sending REST API requests.

GoRest encapsulates HTTP Request properties in Rest Client, which simplify the construction
and invocation of REST API's.

### Features

* Add Base URL to request
* Extend BaseURI to various endpoints using Path
* Attaching Headers to HTTP request
* Send HTTP Request Body as JSON Payload
* Set Query Parameters to Request
* Receive Response Body as string
* Receive success and failure responses in JSON format.
* HTTP Verbs: Get/Post/Put/Patch/Head/Options/Delete

## Install

    go get github.com/bharat72007/gorest

## Documentation

Read [GoDoc](https://godoc.org/github.com/bharat72007/gorest)

## Usage

Use GoRest to set BASE URL, HTTP Verbs, Headers, Query Params, URI Prams or Request Payload
and create a HTTP request.

```go

rest := gorest.New()
request, _ := rest.Base("https://api.example.com/").Path("param1").Query(query).Request()
response, _ := rest.Send(request)

```

### Rest Client

Creating a `New` rest client encapsulating Request Properties

```go

rest := gorest.New()

```

### Path

`Path` method can be used to extend the Base URL to various endpoints, can chain multiple path elements one after other

```go

// creates a GET request to https://api.example.com/users/1
rest := gorest.New()
req, err := rest.Base("https://api.example.com/").Path("users").Path("1").Request()

```

### Http Verbs

Use `Get`, `Post`, `Put`, `Patch`, `Delete`, `Options` or `Head` Http methods.

```go

//Create Http request for Get method
rest.Base("https://api.example.com/").Path("users").Path("1").Get().Request()

```

### Headers

Adding headers to the Http request using `WithHeader` method

```go

rest.Base("https://api.example.com/").Path("users").Path("1").WithHeader("User-Agent", "REST API Client").Get().Request()

```

### Query

User can attach Query Parameters to Request URL.
Query Parameters can be attached in one invocation, user just needs to provide all Query parameters
in {string: string} Map

```go

query := map[string]string{"searchkey1" : "value1",
            "searchkey2" : "value2",
            "searchkey3" : "value3",
        }

rest.Base("https://api.example.com/").Path("param1").Query(query)

```
So complete request URL will be

https://api.example.com/{param1}?searchkey1=value1&searchkey2=value2&searchkey3=value3

`NOTE`: Query is not a chaining method like Path, If user chain Query parameters user are directly replacing
with the latest query

```go

rest.Base("https://api.example.com/").Path("param1").Query(query1).Query(query2)

```
query2 will be part of the request URL and not query1.

### Response
Response can be seen in below mentioned formats
#### JSON Response
HTTP Request can be send to Sever and based on that Response can be collected at various structures.
Success Structure and Error Structure, If Error occurs then error message will be pumped in to Error Structure
provided by End User, similarly for Success Structure.
From the Get Request if user need to collect the JSON Response.

```go

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

```
Here user have collected the Response in the form of http.Response.
To populate Success and Error structures

```go

var post Post
var errpost PostErr
Method Response which is not tied to REST Client can be used.
Response(http.Response, interface{}, interface{})
err := Response(response,post,errpost)

```
From above invokation, post and errpost structures will be populated (based on success or error)


#### Plain Body Response

ResponseString
If user just want to look for response Body as a string, then `ResponseBodyString` method can be used

```go

content := ResponseBodyString(request)

```

### Payload

In Request which needs Body in the request, user needs to construct Payload first.
Right now only JSON format Payload is supported.
Let’s create a structure which user needs to post

```go

type Post struct {
    UserId int    `json:"userid"`
    Title  string `json:"title"`
    Body   string `json:"body"`
}
User can create a Post Payload.
post := POST{
    UserId: "12",
    Title: "any random title",
    Body: "Body content random",
}
rest.Base("https://api.example.com/").Path("param1").Post(post)

```

Similarly, for Partial Updates, user can create partial Payload and invoke Patch Verb.

``` go

partialpost := POST{
    Body: "Body content random",
}
rest.Base("https://api.example.com/").Path("1").Patch(partialpost)

```
Only Body will be updated for id : 1

### Request
From rest client after composing HTTP request properties, user can create http.Request from it.
`Request` method is used for it.

``` go

rest.Base("https://api.example.com/").Path("1").Patch(partialpost).Request() //Will create http.Request out of this Entity.

```
Other examples:

```go

rest.Base("https://api.example.com/").Path("param1").Query(query).Get().Request()
rest.Base("https://api.example.com/").Path("param1").Put(partialpayload).Request()
rest.Base("https://api.example.com/").Path("param1").Post(payload).Request()
rest.Base("https://api.example.com/").Query(query).Head().Request()

```
