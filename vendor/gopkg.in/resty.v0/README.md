# resty  [![Build Status](https://travis-ci.org/go-resty/resty.svg?branch=master)](https://travis-ci.org/go-resty/resty)  [![GoCover](http://gocover.io/_badge/github.com/go-resty/resty)](http://gocover.io/github.com/go-resty/resty)  [![GoDoc](https://godoc.org/github.com/go-resty/resty?status.svg)](https://godoc.org/github.com/go-resty/resty)  [![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

Simple HTTP and REST client for Go inspired by Ruby rest-client. Provides notable features - robust request body input, auto marshal & unmarshal, request and response middlewares, custom & extensible redirect policy (multiple policies can be applied), etc.

#### Features
* Get, Post, Put, Delete, Head, Patch and Options
* Simple methods/chainable methods for settings and request
* Request Body can be `string`, `[]byte`, `struct` and `map`
  * Auto detect `Content-Type`
* Response object gives you more possibility
  * Access as `[]byte` array - `response.Body` OR Access as `string` - `response.String()`
  * Know your `response.Time()` and when we `response.ReceivedAt`
  * Have a look [godoc](https://godoc.org/github.com/go-resty/resty#Response)
* Automatic marshal and unmarshal for `JSON` and `XML` content type
  * Default is `JSON`, if you supply `struct/map` without header `Content-Type`
* Easy to upload single or multiple files via `multipart/form-data`
* Cookies for your request and CookieJar support
* Authorization option of `Basic` and `Bearer` token
* Set request `ContentLength` value for all request or particular request
* Choose between HTTP and RESTful mode. Default is RESTful
  * `HTTP` - default upto 10 redirects and no automatic response unmarshal
  * `RESTful` - default no redirects and automatic response unmarshal for `JSON` & `XML`
* Custom [Root Certificates](https://godoc.org/github.com/go-resty/resty#Client.SetRootCertificate) and Client [Certificates](https://godoc.org/github.com/go-resty/resty#Client.SetCertificates)
* Save HTTP response into File, similar to `curl -o` flag. More info [SetOutputDirectory](https://godoc.org/github.com/go-resty/resty#Client.SetOutputDirectory) & [SetOutput](https://godoc.org/github.com/go-resty/resty#Request.SetOutput).
* Client settings like `Timeout`, `RedirectPolicy`, `Proxy` and `TLSClientConfig`
* Client API design 
  * Have client level settings & options and also override at Request level if you want to
  * [Request](https://godoc.org/github.com/go-resty/resty#Client.OnBeforeRequest) and [Response](https://godoc.org/github.com/go-resty/resty#Client.OnAfterResponse) middleware
  * Create Multiple clients if want to `resty.New()`
  * goroutine concurrent safe
  * Debug mode - clean and informative logging presentation
  * Gzip - I'm not doing anything here. Go does it automatically
* Well tested client library

resty tested with Go v1.2 and above.

#### Included Batteries
  * Redirect Policies
    * NoRedirectPolicy
    * FlexibleRedirectPolicy
    * DomainCheckRedirectPolicy
    * etc. [more info](redirect.go)
  * Persist Cookies into file in JSON format from resty client (upcoming)
  * etc.

## Installation
#### Stable
```sh
go get gopkg.in/resty.v0
```
#### Latest
```sh
go get github.com/go-resty/resty
```

## Usage
Following samples will assist you to become as much comfortable as possible with resty library. Resty comes with ready to use DefaultClient.

Import resty into your code and refer it as `resty`.
```go
import "gopkg.in/resty.v0"
```

#### Simple GET
```go
// GET request
resp, err := resty.R().Get("http://httpbin.org/get")

// explore response object
fmt.Printf("\nError: %v", err)
fmt.Printf("\nResponse Status Code: %v", resp.StatusCode())
fmt.Printf("\nResponse Status: %v", resp.Status())
fmt.Printf("\nResponse Time: %v", resp.Time())
fmt.Printf("\nResponse Recevied At: %v", resp.ReceivedAt)
fmt.Printf("\nResponse Body: %v", resp)
// more...

/* Output
Error: <nil>
Response Status Code: 200
Response Status: 200 OK
Response Time: 644.290186ms
Response Recevied At: 2015-09-15 12:05:28.922780103 -0700 PDT
Response Body: {
  "args": {}, 
  "headers": {
    "Accept-Encoding": "gzip", 
    "Host": "httpbin.org", 
    "User-Agent": "go-resty v0.1 - https://github.com/go-resty/resty"
  }, 
  "origin": "0.0.0.0", 
  "url": "http://httpbin.org/get"
}
*/
```
#### Enhanced GET
```go
resp, err := resty.R().
      SetQueryParams(map[string]string{
          "page_no": "1", 
          "limit": "20",
          "sort":"name",
          "order": "asc",
          "random":strconv.FormatInt(time.Now().Unix(), 10),
      }).
      SetHeader("Accept", "application/json").
      SetAuthToken("BC594900518B4F7EAC75BD37F019E08FBC594900518B4F7EAC75BD37F019E08F").
      Get("/search_result")


// Sample of using Request.SetQueryString method
resp, err := resty.R().
      SetQueryString("productId=232&template=fresh-sample&cat=resty&source=google&kw=buy a lot more").
      SetHeader("Accept", "application/json").
      SetAuthToken("BC594900518B4F7EAC75BD37F019E08FBC594900518B4F7EAC75BD37F019E08F").
      Get("/show_product")
```

#### Various POST method combinations
```go
// POST JSON string
// No need to set content type, if you have client level setting
resp, err := resty.R().
      SetHeader("Content-Type", "application/json").
      SetBody(`{"username":"testuser", "password":"testpass"}`).
      SetResult(&AuthSuccess{}).    // or SetResult(AuthSuccess{}).
      Post("https://myapp.com/login")
      
// POST []byte array
// No need to set content type, if you have client level setting
resp, err := resty.R().
      SetHeader("Content-Type", "application/json").
      SetBody([]byte(`{"username":"testuser", "password":"testpass"}`)).
      SetResult(&AuthSuccess{}).    // or SetResult(AuthSuccess{}).
      Post("https://myapp.com/login")
      
// POST Struct, default is JSON content type. No need to set one
resp, err := resty.R().
      SetBody(User{Username: "testuser", Password: "testpass"}).
      SetResult(&AuthSuccess{}).    // or SetResult(AuthSuccess{}).
      SetError(&AuthError{}).       // or SetError(AuthError{}).
      Post("https://myapp.com/login")
      
// POST Map, default is JSON content type. No need to set one
resp, err := resty.R().
      SetBody(map[string]interface{}{"username": "testuser", "password": "testpass"}).
      SetResult(&AuthSuccess{}).    // or SetResult(AuthSuccess{}).
      SetError(&AuthError{}).       // or SetError(AuthError{}).
      Post("https://myapp.com/login")
      
// POST of raw bytes for file upload. For example: upload file to Dropbox
file, _ := os.Open("/Users/jeeva/mydocument.pdf")
fileBytes, _ := ioutil.ReadAll(file)

// See we are not setting content-type header, since go-resty automatically detects Content-Type for you
resp, err := resty.R().
      SetBody(fileBytes).
      SetContentLength(true).          // Dropbox expects this value
      SetAuthToken("<your-auth-token>").
      SetError(&DropboxError{}).       // or SetError(DropboxError{}).
      Post("https://content.dropboxapi.com/1/files_put/auto/resty/mydocument.pdf") // for upload Dropbox supports PUT too

// Note: resty detects Content-Type for request body/payload if content type header is not set.
//   * For struct and map data type defaults to 'application/json'
//   * Fallback is plain text content type
```

#### Sample PUT 
You can use various combinations of `PUT` method call like demonstrated for `POST`.
```go
// Note: This is one sample of PUT method usage, refer POST for more combination

// Request goes as JSON content type
// No need to set auth token, error, if you have client level settings
resp, err := resty.R().
      SetBody(Article{
        Title: "go-resty", 
        Content: "This is my article content, oh ya!",
        Author: "Jeevanandam M",
        Tags: []string{"article", "sample", "resty"},
      }).
      SetAuthToken("C6A79608-782F-4ED0-A11D-BD82FAD829CD").
      SetError(&Error{}).       // or SetError(Error{}).
      Put("https://myapp.com/article/1234")
```

#### Sample PATCH 
You can use various combinations of `PATCH` method call like demonstrated for `POST`.
```go
// Note: This is one sample of PUT method usage, refer POST for more combination

// Request goes as JSON content type
// No need to set auth token, error, if you have client level settings
resp, err := resty.R().
      SetBody(Article{
        Tags: []string{"new tag1", "new tag2"},
      }).
      SetAuthToken("C6A79608-782F-4ED0-A11D-BD82FAD829CD").
      SetError(&Error{}).       // or SetError(Error{}).
      Patch("https://myapp.com/articles/1234")
```

#### Sample DELETE, HEAD, OPTIONS
```go
// DELETE a article
// No need to set auth token, error, if you have client level settings
resp, err := resty.R().
      SetAuthToken("C6A79608-782F-4ED0-A11D-BD82FAD829CD").
      SetError(&Error{}).       // or SetError(Error{}).
      Delete("https://myapp.com/articles/1234")

// DELETE a articles with payload/body as a JSON string
// No need to set auth token, error, if you have client level settings
resp, err := resty.R().
      SetAuthToken("C6A79608-782F-4ED0-A11D-BD82FAD829CD").
      SetError(&Error{}).       // or SetError(Error{}).
      SetHeader("Content-Type", "application/json").
      SetBody(`{article_ids: [1002, 1006, 1007, 87683, 45432] }`).
      Delete("https://myapp.com/articles")

// HEAD of resource
// No need to set auth token, if you have client level settings
resp, err := resty.R().
      SetAuthToken("C6A79608-782F-4ED0-A11D-BD82FAD829CD").
      Head("https://myapp.com/videos/hi-res-video")

// OPTIONS of resource
// No need to set auth token, if you have client level settings
resp, err := resty.R().
      SetAuthToken("C6A79608-782F-4ED0-A11D-BD82FAD829CD").
      Options("https://myapp.com/servers/nyc-dc-01")
```

### Multipart File(s) upload
```go
// Single file scenario
resp, err := resty.R().
      SetFile("profile_img", "/Users/jeeva/test-img.png").
      Post("http://myapp.com/upload")

// Multiple files scenario
resp, err := resty.R().
      SetFiles(map[string]string{
        "profile_img": "/Users/jeeva/test-img.png", 
        "notes": "/Users/jeeva/text-file.txt",
      }).
      Post("http://myapp.com/upload")

// Multipart of form fields and files
resp, err := resty.R().
      SetFiles(map[string]string{
        "profile_img": "/Users/jeeva/test-img.png", 
        "notes": "/Users/jeeva/text-file.txt",
      }).
      SetFormData(map[string]string{
        "first_name": "Jeevanandam", 
        "last_name": "M",
        "zip_code": "00001", 
        "city": "my city",
        "access_token": "C6A79608-782F-4ED0-A11D-BD82FAD829CD",
      }).
      Post("http://myapp.com/profile")
```

#### Sample Form submision
```go
// just mentioning about POST as an example with simple flow
// User Login
resp, err := resty.R().
      SetFormData(map[string]string{
        "username": "jeeva", 
        "password": "mypass",
      }).
      Post("http://myapp.com/login")

// Followed by profile update
resp, err := resty.R().
      SetFormData(map[string]string{
        "first_name": "Jeevanandam", 
        "last_name": "M",
        "zip_code": "00001", 
        "city": "new city update",
      }).
      Post("http://myapp.com/profile")
```

#### Save HTTP Response into File
```go
// Setting output directory path, If directory not exists then resty creates one!
// This is optional one, if you're planning using absoule path in 
// `Request.SetOutput` and can used together.
resty.SetOutputDirectory("/Users/jeeva/Downloads")

// HTTP response gets saved into file, similar to curl -o flag
_, err := resty.R().
          SetOutput("plugin/ReplyWithHeader-v5.1-beta.zip").
          Get("http://bit.ly/1LouEKr")

// OR using absolute path
// Note: output directory path is not used for absoulte path
_, err := resty.R().
          SetOutput("/MyDownloads/plugin/ReplyWithHeader-v5.1-beta.zip").
          Get("http://bit.ly/1LouEKr")
```

#### Request and Response Middleware
Resty provides middleware ability to manipulate for Request and Response. It is more flexible than callback approach.
```go
// Registering Request Middleware 
resty.OnBeforeRequest(func(c *resty.Client, req *resty.Request) error {
    // Now you have access to Client and current Request object
    // manipulate it as per your need

    return nil  // if its success otherwise return error
  })

// Registering Response Middleware
resty.OnAfterResponse(func(c *resty.Client, resp *resty.Response) error {
    // Now you have access to Client and current Response object
    // manipulate it as per your need

    return nil  // if its success otherwise return error
  })
```

#### Redirect Policy
Resty provides few ready to use redirect policy(s) also it supports multiple policies together.
```go
// Assign Client Redirect Policy. Create one as per you need
resty.SetRedirectPolicy(resty.FlexibleRedirectPolicy(15))

// Wanna multiple policies such as redirect count, domain name check, etc
resty.SetRedirectPolicy(resty.FlexibleRedirectPolicy(20), 
                        resty.DomainCheckRedirectPolicy("host1.com", "host2.org", "host3.net"))
```

##### Custom Redirect Policy
Implement [RedirectPolicy](redirect.go#L20) interface and register it with resty client. Have a look [redirect.go](redirect.go) for more information.
```go
// Using raw func into resty.SetRedirectPolicy
resty.SetRedirectPolicy(resty.RedirectPolicyFunc(func(req *http.Request, via []*http.Request) error {
  // Implement your logic here

  // return nil for continue redirect otherwise return error to stop/prevent redirect
  return nil
}))

//---------------------------------------------------

// Using struct create more flexible redirect policy
type CustomRedirectPolicy struct {
  // variables goes here
}

func (c *CustomRedirectPolicy) Apply(req *http.Request, via []*http.Request) error {
  // Implement your logic here
  
  // return nil for continue redirect otherwise return error to stop/prevent redirect
  return nil
}

// Registering in resty
resty.SetRedirectPolicy(CustomRedirectPolicy{/* initialize variables */})
```

#### Custom Root Certificates and Client Certifcates
```go
// Custom Root certificates, just supply .pem file.
// you can add one or more root certificates, its get appended
resty.SetRootCertificate("/path/to/root/pemFile1.pem")
resty.SetRootCertificate("/path/to/root/pemFile2.pem")
// ... and so on!

// Adding Client Certificates, you add one or more certificates
// Sample for creating certificate object
// Parsing public/private key pair from a pair of files. The files must contain PEM encoded data.
cert1, err := tls.LoadX509KeyPair("certs/client.pem", "certs/client.key")
if err != nil {
  log.Fatalf("ERROR client certificate: %s", err)
}
// ...

// You add one or more certificates
resty.SetCertificates(cert1, cert2, cert3)
```

#### Proxy Settings
Default `Go` supports Proxy via environment variable `HTTP_PROXY`. Resty provides support via `SetProxy` & `RemoveProxy`.
Choose as per your need.
```go
// Setting a Proxy URL and Port
resty.SetProxy("http://proxyserver:8888")

// Want to remove proxy setting
resty.RemoveProxy()
```

#### Choose REST or HTTP mode
```go
// HTTP mode
resty.SetHTTPMode() 

// REST mode. Default one
resty.SetRESTMode()
```

#### Wanna Multiple Clients
```go
// Here you go!
// Client 1
client1 := resty.New()
client1.R().Get("http://httpbin.org")
// ...

// Client 2
client2 := resty.New()
client1.R().Head("http://httpbin.org")
// ...

// Bend it as per your need!!!
```

#### Remaining Client Settings & its Options
```go
// Unique settings at Client level
//--------------------------------
// Enable debug mode
resty.SetDebug(true)

// Using you custom log writer
logFile, _ := os.OpenFile("/Users/jeeva/go-resty.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
resty.SetLogger(logFile)

// Assign Client TLSClientConfig
// One can set custom root-certificate. Refer: http://golang.org/pkg/crypto/tls/#example_Dial
resty.SetTLSClientConfig(&tls.Config{ RootCAs: roots })

// or One can disable security check (https)
resty.SetTLSClientConfig(&tls.Config{ InsecureSkipVerify: true })

// Set client timeout as per your need
resty.SetTimeout(time.Duration(1 * time.Minute))


// You can override all below settings and options at request level if you want to
//--------------------------------------------------------------------------------
// Host URL for all request. So you can use relative URL in the request
resty.SetHostURL("http://httpbin.org")

// Headers for all request
resty.SetHeader("Accept", "application/json")
resty.SetHeaders(map[string]string{
        "Content-Type": "application/json",
        "User-Agent": "My cutsom User Agent String",
      })

// Cookies for all request
resty.SetCookie(&http.Cookie{
      Name:"go-resty",
      Value:"This is cookie value",
      Path: "/",
      Domain: "sample.com",
      MaxAge: 36000,
      HttpOnly: true,
      Secure: false,
    })
resty.SetCookies(cookies)

// URL query parameters for all request
resty.SetQueryParam("user_id", "00001")
resty.SetQueryParams(map[string]string{ // sample of those who use this manner
      "api_key": "api-key-here",
      "api_secert": "api-secert",
    })
resty.R().SetQueryString("productId=232&template=fresh-sample&cat=resty&source=google&kw=buy a lot more")

// Form data for all request. Typically used with POST and PUT
resty.SetFormData(map[string]string{
    "access_token": "BC594900-518B-4F7E-AC75-BD37F019E08F",
  })

// Basic Auth for all request
resty.SetBasicAuth("myuser", "mypass")

// Bearer Auth Token for all request
resty.SetAuthToken("BC594900518B4F7EAC75BD37F019E08FBC594900518B4F7EAC75BD37F019E08F")

// Enabling Content length value for all request
resty.SetContentLength(true)

// Registering global Error object structure for JSON/XML request
resty.SetError(&Error{})    // or resty.SetError(Error{})
```

### Versioning
* resty releases versions according to [Semantic Versioning](http://semver.org)
 
### Contributing
Welcome! If you find any improvement or issue you want to fix. Feel free to send a pull request, I like pull requests that include tests case for fix/enhancement. Did my best to bring pretty good code coverage and feel free to write tests.

BTW, I'd like to know what you think about go-resty. Kindly open an issue or send me an email; it'd mean a lot to me.

### License
resty released under MIT license, refer [LICENSE](LICENSE) file.
