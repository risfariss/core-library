package rest_api

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

type UtilHttpRequest struct {
	GET                               func(params Params) (out map[string]interface{}, err error)
	POST                              func(params Params) (out map[string]interface{}, err error)
	PUT                               func(params Params) (out map[string]interface{}, err error)
	GETWithoutBasicAuth               func(params Params) (out map[string]interface{}, err error)
	POSTWithoutBasicAuth              func(params Params) (out map[string]interface{}, err error)
	GETWithoutBasicAuthSingleResponse func(params Params) (out string, err error)
	POSTFormDataRegister              func(params Params) (out map[string]interface{}, err error)
	POSTWithFileUrl                   func(params Params) (out []byte, err error)
	POSTParam                         func(params Params) (out []byte, header http.Header, err error)
	Post                              func(params Params) (out map[string]interface{}, err error)
}

func NewUtilHttpRequest() *UtilHttpRequest {
	return &UtilHttpRequest{
		GET:                               GET,
		POST:                              POST,
		PUT:                               PUT,
		GETWithoutBasicAuth:               GETWithoutBasicAuth,
		POSTWithoutBasicAuth:              POSTWithoutBasicAuth,
		GETWithoutBasicAuthSingleResponse: GETWithoutBasicAuthSingleResponse,
		POSTFormDataRegister:              POSTFormDataRegister,
		POSTWithFileUrl:                   POSTWithFileUrl,
		POSTParam:                         POSTParam,
		Post:                              Post,
	}
}

type Params struct {
	Url               string
	BasicAuthUsername string
	BasicAuthPassword string
	RequestBody       map[string]interface{}
	RequestParams     map[string]string
	RequestHeader     map[string]string
	Resty             *resty.Client
	BearerToken       string
	File              []FileParams
	FormData          map[string]string
}

type FileParams struct {
	File      []byte
	FileName  string
	ParamName string
}

func NewParams(url string, basicAuthUsername string, basicAuthPassword string, requestBody map[string]interface{},
	requestParams map[string]string, requestHeader map[string]string, bearerToken string) (out Params) {
	out.Url = url
	out.BasicAuthUsername = basicAuthUsername
	out.BasicAuthPassword = basicAuthPassword
	out.RequestBody = requestBody
	out.RequestParams = requestParams
	out.RequestHeader = requestHeader
	out.Resty = resty.New()
	out.BearerToken = bearerToken
	return out
}

func GET(params Params) (out map[string]interface{}, err error) {
	params.Resty.SetTimeout(time.Second * 10)
	params.Resty.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	params.Resty.SetDebug(true)
	host := fmt.Sprint(params.Url)
	resp, err := params.Resty.R().
		SetBasicAuth(params.BasicAuthUsername, params.BasicAuthPassword).
		SetQueryParams(params.RequestParams).
		SetHeaders(params.RequestHeader).
		Get(host)
	if err != nil {
		return
	}

	_ = json.Unmarshal(resp.Body(), &out)

	if resp.StatusCode() >= http.StatusBadRequest {
		return out, errors.New("")
	}
	return
}

func POST(params Params) (out map[string]interface{}, err error) {
	//params.Resty.SetTimeout(time.Second * 10)
	params.Resty.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	params.Resty.SetDebug(true)
	payload, _ := json.Marshal(params.RequestBody)
	log.Println("payload---> ", string(payload))
	host := fmt.Sprint(params.Url)
	resp, err := params.Resty.R().
		SetBasicAuth(params.BasicAuthUsername, params.BasicAuthPassword).
		SetQueryParams(params.RequestParams).
		SetHeaders(params.RequestHeader).
		SetBody(bytes.NewBuffer(payload)).
		SetFormData(params.FormData).
		Post(host)
	if err != nil {
		return
	}

	_ = json.Unmarshal(resp.Body(), &out)

	if resp.StatusCode() >= http.StatusBadRequest {
		return out, errors.New("")
	}

	return
}

func POSTParam(params Params) (out []byte, header http.Header, err error) {
	params.Resty.SetTimeout(time.Second * 10)
	params.Resty.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	params.Resty.SetDebug(true)
	payload, _ := json.Marshal(params.RequestBody)
	log.Println("payload---> ", string(payload))
	host := fmt.Sprint(params.Url)
	resp, err := params.Resty.R().
		SetBasicAuth(params.BasicAuthUsername, params.BasicAuthPassword).
		SetQueryParams(params.RequestParams).
		SetBody(bytes.NewBuffer(payload)).
		Post(host)
	fmt.Println("@----- error", err)
	if err != nil {
		return
	}

	if resp.StatusCode() >= http.StatusBadRequest {
		return resp.Body(), resp.Header(), errors.New("")
	}

	return
}

func PUT(params Params) (out map[string]interface{}, err error) {
	params.Resty.SetTimeout(time.Second * 10)
	params.Resty.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	params.Resty.SetDebug(true)
	payload, _ := json.Marshal(params.RequestBody)
	log.Println("payload---> ", string(payload))
	host := fmt.Sprint(params.Url)
	var resp *resty.Response
	resp, err = params.Resty.R().
		SetBasicAuth(params.BasicAuthUsername, params.BasicAuthPassword).
		SetBody(bytes.NewBuffer(payload)).
		SetQueryParams(params.RequestParams).
		SetHeaders(params.RequestHeader).
		Put(host)
	if err != nil {
		return
	}

	_ = json.Unmarshal(resp.Body(), &out)

	if resp.StatusCode() >= http.StatusBadRequest {
		return out, errors.New("")
	}

	return
}

func DELETE(params Params) (out map[string]interface{}, err error) {
	params.Resty.SetTimeout(time.Second * 10)
	params.Resty.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	params.Resty.SetDebug(true)
	payload, _ := json.Marshal(params.RequestBody)
	log.Println("payload---> ", string(payload))
	host := fmt.Sprint(params.Url)
	var resp *resty.Response
	resp, err = params.Resty.R().
		SetBasicAuth(params.BasicAuthUsername, params.BasicAuthPassword).
		SetBody(bytes.NewBuffer(payload)).
		SetQueryParams(params.RequestParams).
		SetHeaders(params.RequestHeader).
		Delete(host)
	if err != nil {
		return
	}

	_ = json.Unmarshal(resp.Body(), &out)

	if resp.StatusCode() >= http.StatusBadRequest {
		return out, errors.New("")
	}

	return
}

func GETWithoutBasicAuthSingleResponse(params Params) (out string, err error) {
	params.Resty.SetTimeout(time.Second * 10)
	params.Resty.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	params.Resty.SetDebug(true)
	host := fmt.Sprint(params.Url)
	resp, _ := params.Resty.R().
		SetAuthToken(params.BearerToken).
		SetQueryParams(params.RequestParams).
		SetHeaders(params.RequestHeader).
		Get(host)
	//if errs != nil { notes: single response with string cannot get error, always return the response
	//	return
	//}
	out = string(resp.Body())

	if resp.StatusCode() >= http.StatusBadRequest {
		return out, errors.New("")
	}
	return
}

func GETWithoutBasicAuth(params Params) (out map[string]interface{}, err error) {
	params.Resty.SetTimeout(time.Second * 10)
	params.Resty.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	params.Resty.SetDebug(true)
	host := fmt.Sprint(params.Url)
	resp, err := params.Resty.R().
		SetBasicAuth(params.BasicAuthUsername, params.BasicAuthPassword).
		SetQueryParams(params.RequestParams).
		Get(host)

	if err != nil {
		return
	}

	_ = json.Unmarshal(resp.Body(), &out)

	if resp.StatusCode() >= http.StatusBadRequest {
		return out, errors.New("")
	}
	return
}

func POSTWithoutBasicAuth(params Params) (out map[string]interface{}, err error) {
	params.Resty.SetTimeout(time.Second * 10)
	params.Resty.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	params.Resty.SetDebug(true)
	payload, _ := json.Marshal(params.RequestBody)
	log.Println("payload---> ", string(payload))
	host := fmt.Sprint(params.Url)
	var resp *resty.Response

	resp, err = params.Resty.R().
		SetBody(bytes.NewBuffer(payload)).
		SetHeaders(params.RequestHeader).
		Post(host)
	if err != nil {
		return
	}

	_ = json.Unmarshal(resp.Body(), &out)
	if resp.StatusCode() >= http.StatusBadRequest {
		return out, errors.New("")
	}
	return
}

func POSTFormDataRegister(params Params) (out map[string]interface{}, err error) {

	if len(params.File) > 1 {
		return SendPostBulkFile(params)
	}

	params.Resty.SetTimeout(time.Second * 30)
	params.Resty.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	params.Resty.SetDebug(true)

	host := fmt.Sprint(params.Url)

	var document []byte
	var documentName string
	var paramName string

	for _, data := range params.File {
		documentName = data.FileName
		document = data.File
		paramName = data.ParamName
	}

	resp, err := params.Resty.R().
		SetHeaders(params.RequestHeader).
		SetBasicAuth(params.BasicAuthUsername, params.BasicAuthPassword).
		SetFileReader(paramName, documentName, bytes.NewReader(document)).
		SetFormData(params.FormData).
		Post(host)

	if err != nil {
		return
	}

	_ = json.Unmarshal(resp.Body(), &out)

	if resp.StatusCode() >= http.StatusBadRequest {
		return out, errors.New("")
	}
	return

}

func SendPostBulkFile(params Params) (out map[string]interface{}, err error) {
	/* for register privy_rest */
	params.Resty.SetTimeout(time.Second * 30)
	params.Resty.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	params.Resty.SetDebug(true)

	host := fmt.Sprint(params.Url)

	var ktp []byte
	var ktpName string
	var ktpSelfie []byte
	var ktpSelfieName string

	for _, data := range params.File {
		if data.ParamName == "ktp" {
			ktpName = data.FileName
			ktp = data.File
		} else {
			ktpSelfieName = data.FileName
			ktpSelfie = data.File
		}
	}

	resp, err := params.Resty.R().
		SetHeaders(params.RequestHeader).
		SetBasicAuth(params.BasicAuthUsername, params.BasicAuthPassword).
		SetFileReader("ktp", ktpName, bytes.NewReader(ktp)).
		SetFileReader("selfie", ktpSelfieName, bytes.NewReader(ktpSelfie)).
		SetFormData(params.FormData).
		Post(host)

	if err != nil {
		return
	}

	_ = json.Unmarshal(resp.Body(), &out)

	if resp.StatusCode() >= http.StatusBadRequest {
		return out, errors.New("")
	}
	return
}

func POSTWithFileUrl(params Params) (out []byte, err error) {
	params.Resty.SetTimeout(time.Second * 30)
	params.Resty.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	params.Resty.SetDebug(true)
	payload, _ := json.Marshal(params.RequestBody)

	host := fmt.Sprint(params.Url)
	resp, err := params.Resty.R().
		SetBasicAuth(params.BasicAuthUsername, params.BasicAuthPassword).
		SetBody(bytes.NewBuffer(payload)).
		Post(host)
	if err != nil {
		return
	}

	//_ = json.Unmarshal(resp.Body(), &out)
	out = resp.Body()
	if resp.StatusCode() >= http.StatusBadRequest {
		return out, errors.New("")
	}

	return
}

func Post(params Params) (out map[string]interface{}, err error) {
	params.Resty.SetTimeout(time.Second * 10)
	params.Resty.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	params.Resty.SetDebug(true)
	payload, _ := json.Marshal(params.RequestBody)
	host := fmt.Sprint(params.Url)
	resp, err := params.Resty.R().
		SetHeaders(params.RequestHeader).
		SetBody(bytes.NewBuffer(payload)).
		Post(host)
	if err != nil {
		return map[string]interface{}{}, err
	}

	_ = json.Unmarshal(resp.Body(), &out)

	if resp.StatusCode() >= http.StatusBadRequest {
		return out, err
	}
	return out, nil
}
