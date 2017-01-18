package gorest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	ContentType     = "Content-Type"
	JsonContentType = "application/json"
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
		httpClient:  http.Client{Timeout: time.Second * 10},
		headers:     make(map[string][]string),
		queryvalues: make(url.Values),
	}
}

func (r *Rest) AddHeader(key, value string) *Rest {
	r.headers.Add(key, value)
	return r
}

func (r *Rest) URIParam(param string) *Rest {
	if r.baseurl == "" {
		panic("BASE URL Not Present")
	}
	r.uriparams = append(r.uriparams, param)
	return r
}

func (r *Rest) BasePath(baseurl string) *Rest {
	v, err := url.Parse(baseurl)
	if err != nil {
		panic("Url is incorrect")
	}
	r.baseurl = v.String()
	return r
}

func (r *Rest) Get() *Rest {
	r.verb = "GET"
	return r
}

func (r *Rest) addPayload(payload interface{}) {
	var b []byte
	b, _ = json.Marshal(payload)
	r.payload = bytes.NewBuffer(b)

}

func (r *Rest) Post(payload interface{}) *Rest {
	r.verb = "POST"
	if payload != nil {
		r.addPayload(payload)
	}
	return r
}

func (r *Rest) Request() (*http.Request, error) {
	//Add all URI params to baseurl
	r.url = r.baseurl
	for _, param := range r.uriparams {
		r.url = r.url + param
	}
	v, err := url.Parse(r.url)

	if err != nil {
		panic("Complete URL is not correct")
	}

	r.url = v.String()

	//Adding Query String
	var requrl url.URL
	requrl.Path = r.url
	requrl.RawQuery = r.queryvalues.Encode()
	fmt.Println(requrl.String())

	req, err := http.NewRequest(r.verb, requrl.String(), r.payload)
	if err != nil {
		panic("Request Object is not proper")
	}

	//Add headers to the Request
	for key, values := range r.headers {
		for _, value := range values {
			req.Header.Add(key, value)
		}

	}

	return req, err
}

func (r *Rest) SetQuery(options ...interface{}) *Rest {
	if len(options) > 0 {
		qry, ok := options[0].(map[string]string)
		if ok {
			for k, v := range qry {
				r.queryvalues.Set(k, v)
			}
		}
	}
	return r
}

func (r *Rest) Send(req *http.Request, successM, failureM interface{}) (*http.Response, error) {
	response, err := r.httpClient.Do(req)

	if err != nil {
		panic("Send request Failed")
	}
	return response, err
}

func (r *Rest) ResponseBodyString(response *http.Response, st interface{}) string {
	json.NewDecoder(response.Body).Decode(st)
	fmt.Println(st)
	responsedata, _ := ioutil.ReadAll(response.Body)
	return string(responsedata)
}

func (r *Rest) ResponseStructure(response *http.Response, st interface{}) error {
	err := json.NewDecoder(response.Body).Decode(st)
	fmt.Println(st)
	return err
}
