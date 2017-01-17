package gorest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	ContentType     = "Content-Type"
	JsonContentType = "application/json"
	TextContentType = "application/text"
	XmlContentType  = "application/xml"
	timeOutSeconds  = 10
)

type Rest struct {
	httpClient  http.Client
	baseurl     string
	verb        string
	headers     http.Header
	uriparams   []string
	url         string
	queryvalues url.Values
	payload     io.Reader
}

func New() *Rest {
	return &Rest{
		httpClient:  http.Client{Timeout: time.Second * timeOutSeconds},
		headers:     make(map[string][]string),
		queryvalues: make(url.Values),
	}
}

func (client *Rest) Get() *Rest {
	client.verb = "GET"
	return client
}

func (client *Rest) Post(payload interface{}) *Rest {
	client.verb = "POST"
	if payload != nil {
		client.addPayload(payload)
	}
	return client
}

func (client *Rest) Put(payload interface{}) *Rest {
	client.verb = "PUT"
	if payload != nil {
		client.addPayload(payload)
	}
	return client
}

func (client *Rest) Delete() *Rest {
	client.verb = "DELETE"
	return client
}

func (client *Rest) Patch(payload interface{}) *Rest {
	client.verb = "PATCH"
	if payload != nil {
		client.addPayload(payload)
	}
	return client
}

func (client *Rest) Head() *Rest {
	client.verb = "HEAD"
	return client
}

func (client *Rest) Trace() *Rest {
	client.verb = "TRACE"
	return client
}

func (client *Rest) Path(param string) *Rest {
	if client.baseurl == "" {
		panic("BASE URL Not Present")
	} else if strings.Contains(param, "/") {
		panic("Incorrect PATH Param")
	}
	client.uriparams = append(client.uriparams, param)
	return client
}

func (client *Rest) Base(baseurl string) *Rest {
	v, err := url.Parse(baseurl)
	if err != nil {
		panic("Url is incorrect")
	}
	client.baseurl = v.String()
	return client
}

func (client *Rest) Header(key, value string) *Rest {
	client.headers.Add(key, value)
	return client
}

func (client *Rest) addPayload(payload interface{}) {
	var b []byte
	b, _ = json.Marshal(payload)
	client.payload = bytes.NewBuffer(b)

}

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

func (client *Rest) Request() (*http.Request, error) {
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
	fmt.Println(requrl.String())

	req, err := http.NewRequest(client.verb, requrl.String(), client.payload)
	if err != nil {
		panic("Request Object is not proper")
	}

	//Add headers to the Request
	for key, values := range client.headers {
		for _, value := range values {
			req.Header.Add(key, value)
		}

	}
	return req, err
}

func (client *Rest) ResponseBodyString(response *http.Response) string {
	responsedata, _ := ioutil.ReadAll(response.Body)
	return string(responsedata)
}

func (client *Rest) Send(req *http.Request, successM, failureM interface{}) (*http.Response, error) {
	response, err := client.httpClient.Do(req)

	if err != nil {
		panic("Send request Failed")
	}
	return response, err
}
