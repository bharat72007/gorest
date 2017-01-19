package gorest

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

//All Constants which are related to headers.

const (
	ContentType     = "Content-Type"
	JsonContentType = "application/json"
	TextContentType = "application/text"
	XmlContentType  = "application/xml"
	timeOutSeconds  = 10
)

// Rest Entity encapsulate Client properties and executes http requests.
//It is implemented by *http.Client. End User can make use of Rest Entity to perform API invocations.

type Rest struct {
	httpClient  http.Client
	baseurl     string
	verb        string
	headers     http.Header
	uriparams   []string
	url         string
	queryvalues url.Values
	payload     io.Reader
	formdata    url.Values
}

//New Rest Client can be initialized from Here.
//Rest Client can be used to Invoke REST API's.
//Example: New().Base("some-uri.com").Path(param).Get().Request()
//Default client creation is ignored. //TODO
func New() *Rest {
	return &Rest{
		httpClient:  http.Client{Timeout: time.Second * timeOutSeconds},
		headers:     make(map[string][]string),
		queryvalues: make(url.Values),
	}
}

//Function to attach Header to Client, which in turns used with Request.
//Headers are added in the form of Key-Value Pair
func (client *Rest) WithHeader(key, value string) *Rest {
	client.headers.Add(key, value)
	return client
}

//Path Params aka URI Params can be added to the Base URI.
//Example: base_uri/{param1}/{param2} ==> Client.Base(BASEURI).Path(param1).Path(param2) etc...
func (client *Rest) Path(param string) *Rest {
	if client.baseurl == "" {
		panic("BASE URL Not Present")
	}
	client.uriparams = append(client.uriparams, param)
	return client
}

//BASE URL can be attached to Rest Client
//Example: Client.Base("https://some-service/app")
func (client *Rest) Base(baseurl string) *Rest {
	v, err := url.Parse(baseurl)
	if err != nil {
		panic("Url is incorrect")
	}
	client.baseurl = v.String()
	return client
}

//GET HTTP/HTTPS method, In case API request of the form GET, inorder to retrieve resources from server
func (client *Rest) Get() *Rest {
	client.verb = "GET"
	return client
}

//POST HTTP/HTTPS method, In case API request of the form POST, in order to create new resource at server
//In this method we can also attach Payload to be send to Server.
func (client *Rest) Post(payload interface{}) *Rest {
	client.verb = "POST"
	if payload != nil {
		client.WithPayload(payload)
	}
	return client
}

//PUT HTTP/HTTPS method, In case API request of the form PUT, in order to update the existing resource at server completely(except Primary Keys or Attributes which
// identify resource uniquely)
//In this method we can also attach updated Payload to be send to Server.
func (client *Rest) Put(payload interface{}) *Rest {
	client.verb = "PUT"
	if payload != nil {
		client.WithPayload(payload)
	}
	return client
}

//DELETE HTTP/HTTPS method, In case API request of the form delete, in order to delete the existing resource at server
//In this method we can also attach updated Payload to be send to Server.
func (client *Rest) Delete(payload ...interface{}) *Rest {
	client.verb = "DELETE"
	if payload != nil {
		client.WithPayload(payload)
	}
	return client
}

//PATCH HTTP/HTTPS method,In case API request of the form patch, in order to partial update the existing resource  properties at server
//In this method we can attach the partial Payload fro updates.
func (client *Rest) Patch(payload interface{}) *Rest {
	client.verb = "PATCH"
	if payload != nil {
		client.WithPayload(payload)
	}
	return client
}

//Head HTTP/HTTPS method, In case API request of the form Head, in order reterive Headers
func (client *Rest) Head() *Rest {
	client.verb = "HEAD"
	return client
}

//OPTIONS HTTP/HTTPS method, In case API request of the form OPTIONS, which will provide all supported HTTP VERBS API.
func (client *Rest) Option(payload interface{}) *Rest {
	client.verb = "OPTIONS"
	if payload != nil {
		client.WithPayload(payload)
	}
	return client
}

func (client *Rest) Copy() *Rest {
	client.verb = "COPY"
	return client
}

//Construct the Form Data provided by User and attach it with REST Client.
func (client *Rest) WithFormData(data map[string]string) *Rest {
	form := url.Values{}
	for k, v := range data {
		form.Add(k, v)
	}
	client.formdata = form
	client.WithHeader("Content-Type", "application/x-www-form-urlencoded")
	return client
}

//Construct Payload, to be used for POST/PUT/PATCH methods
func (client *Rest) WithPayload(payload interface{}) {
	var b []byte
	b, _ = json.Marshal(payload)
	client.payload = bytes.NewBuffer(b)
}

//From all the Rest Client Entity, HTTP request will be created. User can also provide the Authunication mode also, whether BASIC or OAUTH.
//Example: BASIC Authencticaton BasicAuth{Username:"xyz", Password:"abcd"}
func (client *Rest) Request(authoptions ...interface{}) (*http.Request, error) {
	//Add all URI params to baseurl
	if !strings.HasSuffix(client.baseurl, "/") {
		client.baseurl = client.baseurl + "/"
	}

	client.url = client.baseurl
	params := strings.Join(client.uriparams, "/")
	client.url = client.baseurl + params
	v, err := url.Parse(client.url)
	if err != nil {
		panic("Complete URL is not correct")
	}

	client.url = v.String()

	//Adding Query String
	var requrl url.URL
	requrl.Path = client.url
	requrl.RawQuery = client.queryvalues.Encode()
	var req *http.Request
	if client.formdata != nil {
		req, err = http.NewRequest(client.verb, requrl.String(), strings.NewReader(client.formdata.Encode()))
	} else {
		req, err = http.NewRequest(client.verb, requrl.String(), client.payload)
	}
	if err != nil {
		panic("Request Object is not proper")
	}

	//Add headers to the Request
	for key, values := range client.headers {
		for _, value := range values {
			req.Header.Add(key, value)
		}

	}

	for _, opt := range authoptions {
		switch v := opt.(type) {
		case *BasicAuth:
			req.SetBasicAuth(v.Username, v.Password)
		case *OAuth2: //TODO
		}
	}
	return req, err
}

//Query Paramters can be attached to URI.
//Example: ?key1=val1&key2=val2 etc. Query paramters are added in the form of {Key,Value} Pair
func (client *Rest) Query(options ...interface{}) *Rest {
	if len(options) > 0 {
		qry, ok := options[0].(map[string]string)
		if ok {
			for k, v := range qry {
				client.queryvalues.Set(k, v)
			}
		}
	}
	return client
}

//After Request Construction, http.Request will be send to the Server for that http.Client.Do method is invoked.
//From this API invokation, http Response and errors are being returned.
func (client *Rest) Send(req *http.Request) (*http.Response, error) {
	response, err := client.httpClient.Do(req)
	if err != nil {
		panic("Send request Failed")
	}
	return response, err
}

//Utility function not tied with specific Client interface, this method can be used to return response body as string
func ResponseBodyString(response *http.Response, st interface{}) string {
	if response != nil {
		responsedata, _ := ioutil.ReadAll(response.Body)
		return string(responsedata)
	}
	return nil
}

//Utility function to decode response Body as response struct or error struct (JSON Object)
//Unmarshalling process ==> To convert Response to respective JSON objects.
//In case the response status is [200-299] then send the Success Response.
func Response(response *http.Response, responsestruct, errorstruct interface{}) error {
	var err error
	if code := response.StatusCode; 200 <= code && code <= 299 {
		err = json.NewDecoder(response.Body).Decode(responsestruct)
	} else {
		err = json.NewDecoder(response.Body).Decode(errorstruct)
	}
	return err
}
