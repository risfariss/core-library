package rest_api

import "net/http"

type RestApi interface {
	HitApiService(httpMethod string, params Params) (out map[string]interface{}, err error)
	HitApiServiceWithoutBasicAuth(httpMethod string, params Params) (out map[string]interface{}, err error)
	HitApiServiceWithoutBasicAuthSingleResponse(httpMethod string, params Params) (out string, err error)
	GenerateBearerToken(in string, clientId string, clientSecret string, grantType string, provisionKey string, tokenURL string) (out string, err error)
	HitApiServiceParam(httpMethod string, params Params) (out []byte, header http.Header, err error)
	HitApiServiceGeneral(httpMethod string, params Params) (out map[string]interface{}, err error)
}
