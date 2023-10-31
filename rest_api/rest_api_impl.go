package rest_api

import (
	"bitbucket.org/kawancicil/core-library/constant"
	"bitbucket.org/kawancicil/core-library/rest_api/request"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type RestApiUtils struct {
	/*todo fill this block if need connection with another interface*/
	utilHttpRequest *UtilHttpRequest
}

func InitRestApiUtils(utilHttpRequest *UtilHttpRequest) RestApi {
	return &RestApiUtils{
		utilHttpRequest: utilHttpRequest,
	}
}

func (r *RestApiUtils) HitApiService(httpMethod string, params Params) (out map[string]interface{}, err error) {
	if httpMethod == http.MethodPost {
		out, err = POST(params)
	} else if httpMethod == http.MethodGet {
		out, err = GET(params)
	} else if httpMethod == http.MethodPut {
		out, err = PUT(params)
	} else if httpMethod == http.MethodDelete {
		out, err = DELETE(params)
	} else if httpMethod == constant.PostRegisterPrivy {
		out, err = POSTFormDataRegister(params)
	}
	if err != nil {
		errorReq := out["error_description"]
		if errorReq != nil {
			return out, errors.New(fmt.Sprintf("%s", errorReq))
		}
		return out, err
	}
	return
}

func (r *RestApiUtils) HitApiServiceWithoutBasicAuth(httpMethod string, params Params) (out map[string]interface{}, err error) {
	if httpMethod == http.MethodPost {
		out, err = POSTWithoutBasicAuth(params)
	} else if httpMethod == http.MethodGet {
		out, err = GETWithoutBasicAuth(params)
	}
	if err != nil {
		errorReq := out["error_description"]
		return out, errors.New(fmt.Sprintf("%s", errorReq))
	}
	return
}

func (r *RestApiUtils) HitApiServiceWithoutBasicAuthSingleResponse(httpMethod string, params Params) (out string, err error) {
	if httpMethod == http.MethodGet {
		out, err = GETWithoutBasicAuthSingleResponse(params)
	}
	if err != nil {
		return out, errors.New(err.Error())
	}
	return
}

func (r *RestApiUtils) GenerateBearerToken(in string, clientId string, clientSecret string, grantType string, provisionKey string, tokenURL string) (out string, err error) {
	requestBody := map[string]interface{}{
		"client_id":            clientId,
		"client_secret":        clientSecret,
		"grant_type":           grantType,
		"authenticated_userid": in,
		"provision_key":        provisionKey,
	}
	params := NewParams(tokenURL,
		"", "", requestBody, map[string]string{}, map[string]string{},
		"")

	resp, err := POST(params)
	if err != nil {
		errorReq := resp["error_description"]
		return out, errors.New(fmt.Sprintf("%s", errorReq))
	}
	gg, _ := json.Marshal(resp)
	var token request.BearerToken
	err = json.Unmarshal(gg, &token)
	out = token.AccessToken
	return
}

func (r *RestApiUtils) HitApiServiceParam(httpMethod string, params Params) (out []byte, header http.Header, err error) {
	if httpMethod == http.MethodPost {
		out, header, err = POSTParam(params)
	}

	return
}

func (r *RestApiUtils) HitApiServiceGeneral(httpMethod string, params Params) (out map[string]interface{}, err error) {
	if httpMethod == http.MethodPost {
		out, err = r.utilHttpRequest.Post(params)
	}
	if err != nil {
		errorReq := out["error_description"]
		return out, errors.New(fmt.Sprintf("%s", errorReq))
	}
	return
}
