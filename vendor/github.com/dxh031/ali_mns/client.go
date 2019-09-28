package ali_mns

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"net/http"
	neturl "net/url"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gogap/errors"
	"github.com/valyala/fasthttp"
)

const (
	DefaultQueueQPSLimit int32 = 2000
	DefaultTopicQPSLimit int32 = 2000
	DefaultDNSTTL        int32 = 10
)

const (
	GLOBAL_PROXY = "MNS_GLOBAL_PROXY"
)

const (
	version = "2015-06-06"
)

const (
	DefaultTimeout int64 = 35
)

type Method string

var (
	errMapping map[string]errors.ErrCodeTemplate
)

func init() {
	initMNSErrors()
}

const (
	GET    Method = "GET"
	PUT           = "PUT"
	POST          = "POST"
	DELETE        = "DELETE"
)

type MNSClient interface {
	Send(method Method, headers map[string]string, message interface{}, resource string) (*fasthttp.Response, error)
	SetProxy(url string)

	getAccountID() (accountId string)
	getRegion() (region string)
}

type aliMNSClient struct {
	Timeout     int64
	url         *neturl.URL
	credential  Credential
	accessKeyId string
	client      *fasthttp.Client
	proxyURL    string

	accountId string
	region    string

	clientLocker sync.Mutex
}

func NewAliMNSClient(inputUrl, accessKeyId, accessKeySecret string) MNSClient {
	if inputUrl == "" {
		panic("ali-mns: message queue url is empty")
	}

	credential := NewAliMNSCredential(accessKeySecret)

	cli := new(aliMNSClient)
	cli.credential = credential
	cli.accessKeyId = accessKeyId

	var err error
	if cli.url, err = neturl.Parse(inputUrl); err != nil {
		panic("err parse url")
	}

	// 1. parse region and accountid
	pieces := strings.Split(inputUrl, ".")
	if len(pieces) != 5 {
		panic("ali-mns: message queue url is invalid")
	}

	accountIdSlice := strings.Split(pieces[0], "/")
	cli.accountId = accountIdSlice[len(accountIdSlice)-1]

	regionSlice := strings.Split(pieces[2], "-internal")
	cli.region = regionSlice[0]

	if globalurl := os.Getenv(GLOBAL_PROXY); globalurl != "" {
		cli.proxyURL = globalurl
	}

	// 2. now init http client
	cli.initFastHttpClient()

	return cli
}

func (p aliMNSClient) getAccountID() (accountId string) {
	return p.accountId
}

func (p aliMNSClient) getRegion() (region string) {
	return p.region
}

func (p *aliMNSClient) SetProxy(url string) {
	if url == p.proxyURL {
		return
	}

	p.proxyURL = url
}

func (p *aliMNSClient) initFastHttpClient() {
	p.clientLocker.Lock()
	defer p.clientLocker.Unlock()

	timeoutInt := DefaultTimeout

	if p.Timeout > 0 {
		timeoutInt = p.Timeout
	}

	timeout := time.Second * time.Duration(timeoutInt)

	p.client = &fasthttp.Client{ReadTimeout: timeout, WriteTimeout: timeout}
}

func (p *aliMNSClient) proxy(req *http.Request) (*neturl.URL, error) {
	if p.proxyURL != "" {
		return neturl.Parse(p.proxyURL)
	}
	return nil, nil
}

func (p *aliMNSClient) authorization(method Method, headers map[string]string, resource string) (authHeader string, err error) {
	if signature, e := p.credential.Signature(method, headers, resource); e != nil {
		return "", e
	} else {
		authHeader = fmt.Sprintf("MNS %s:%s", p.accessKeyId, signature)
	}

	return
}

func (p *aliMNSClient) Send(method Method, headers map[string]string, message interface{}, resource string) (*fasthttp.Response, error) {
	var xmlContent []byte
	var err error

	if message == nil {
		xmlContent = []byte{}
	} else {
		switch m := message.(type) {
		case []byte:
			{
				xmlContent = m
			}
		default:
			if bXml, e := xml.Marshal(message); e != nil {
				err = ERR_MARSHAL_MESSAGE_FAILED.New(errors.Params{"err": e})
				return nil, err
			} else {
				xmlContent = bXml
			}
		}
	}

	xmlMD5 := md5.Sum(xmlContent)
	strMd5 := fmt.Sprintf("%x", xmlMD5)

	if headers == nil {
		headers = make(map[string]string)
	}

	headers[MQ_VERSION] = version
	headers[CONTENT_TYPE] = "application/xml"
	headers[CONTENT_MD5] = base64.StdEncoding.EncodeToString([]byte(strMd5))
	headers[DATE] = time.Now().UTC().Format(http.TimeFormat)

	if authHeader, e := p.authorization(method, headers, fmt.Sprintf("/%s", resource)); e != nil {
		err = ERR_GENERAL_AUTH_HEADER_FAILED.New(errors.Params{"err": e})
		return nil, err
	} else {
		headers[AUTHORIZATION] = authHeader
	}

	var buffer bytes.Buffer
	buffer.WriteString(p.url.String())
	buffer.WriteString("/")
	buffer.WriteString(resource)

	url := buffer.String()

	// 莫名的lock 加这个是为了啥 想不通。。 推拉模式 加lock 这是直接限流请求了
	// p.clientLocker.Lock()
	// defer p.clientLocker.Unlock()
	req := fasthttp.AcquireRequest()

	req.SetRequestURI(url)
	req.Header.SetMethod(string(method))
	req.SetBody(xmlContent)

	for header, value := range headers {
		req.Header.Set(header, value)
	}

	resp := fasthttp.AcquireResponse()

	if err = p.client.Do(req, resp); err != nil {
		err = ERR_SEND_REQUEST_FAILED.New(errors.Params{"err": err})
		return nil, err
	}

	return resp, nil
}

func initMNSErrors() {
	errMapping = map[string]errors.ErrCodeTemplate{
		"AccessDenied":                ERR_MNS_ACCESS_DENIED,
		"InvalidAccessKeyId":          ERR_MNS_INVALID_ACCESS_KEY_ID,
		"InternalError":               ERR_MNS_INTERNAL_ERROR,
		"InvalidAuthorizationHeader":  ERR_MNS_INVALID_AUTHORIZATION_HEADER,
		"InvalidDateHeader":           ERR_MNS_INVALID_DATE_HEADER,
		"InvalidArgument":             ERR_MNS_INVALID_ARGUMENT,
		"InvalidDegist":               ERR_MNS_INVALID_DEGIST,
		"InvalidRequestURL":           ERR_MNS_INVALID_REQUEST_URL,
		"InvalidQueryString":          ERR_MNS_INVALID_QUERY_STRING,
		"MalformedXML":                ERR_MNS_MALFORMED_XML,
		"MissingAuthorizationHeader":  ERR_MNS_MISSING_AUTHORIZATION_HEADER,
		"MissingDateHeader":           ERR_MNS_MISSING_DATE_HEADER,
		"MissingVersionHeader":        ERR_MNS_MISSING_VERSION_HEADER,
		"MissingReceiptHandle":        ERR_MNS_MISSING_RECEIPT_HANDLE,
		"MissingVisibilityTimeout":    ERR_MNS_MISSING_VISIBILITY_TIMEOUT,
		"MessageNotExist":             ERR_MNS_MESSAGE_NOT_EXIST,
		"QueueAlreadyExist":           ERR_MNS_QUEUE_ALREADY_EXIST,
		"QueueDeletedRecently":        ERR_MNS_QUEUE_DELETED_RECENTLY,
		"InvalidQueueName":            ERR_MNS_INVALID_QUEUE_NAME,
		"QueueNameLengthError":        ERR_MNS_QUEUE_NAME_LENGTH_ERROR,
		"QueueNotExist":               ERR_MNS_QUEUE_NOT_EXIST,
		"ReceiptHandleError":          ERR_MNS_RECEIPT_HANDLE_ERROR,
		"SignatureDoesNotMatch":       ERR_MNS_SIGNATURE_DOES_NOT_MATCH,
		"TimeExpired":                 ERR_MNS_TIME_EXPIRED,
		"QpsLimitExceeded":            ERR_MNS_QPS_LIMIT_EXCEEDED,
		"TopicAlreadyExist":           ERR_MNS_TOPIC_ALREADY_EXIST,
		"TopicNameLengthError":        ERR_MNS_TOPIC_NAME_LENGTH_ERROR,
		"TopicNotExist":               ERR_MNS_TOPIC_NOT_EXIST,
		"SubscriptionNameLengthError": ERR_MNS_SUBSRIPTION_NAME_LENGTH_ERROR,
		"TopicNameInvalid":            ERR_MNS_INVALID_TOPIC_NAME,
		"SubsriptionNameInvalid":      ERR_MNS_INVALID_SUBSCRIPTION_NAME,
		"SubscriptionAlreadyExist":    ERR_MNS_SUBSCRIPTION_ALREADY_EXIST,
		"EndpointInvalid":             ERR_MNS_INVALID_ENDPOINT,
		"SubscriberNotExist":          ERR_MNS_SUBSCRIBER_NOT_EXIST,
	}
}

func ParseError(resp ErrorResponse, resource string) (err error) {
	if errCodeTemplate, exist := errMapping[resp.Code]; exist {
		err = errCodeTemplate.New(errors.Params{"resp": resp, "resource": resource})
	} else {
		err = ERR_MNS_UNKNOWN_CODE.New(errors.Params{"resp": resp, "resource": resource})
	}
	return
}
