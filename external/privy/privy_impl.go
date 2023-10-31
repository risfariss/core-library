package privy

import (
	"bitbucket.org/kawancicil/core-library/constant"
	"bitbucket.org/kawancicil/core-library/external/privy/request"
	"bitbucket.org/kawancicil/core-library/external/privy/response"
	"bitbucket.org/kawancicil/core-library/rest_api"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"net/http"
)

type PrivyUtils struct {
	/*todo fill this block if need connection with another interface*/
	restApi rest_api.RestApi
}

func InitPrivyUtils(restApi rest_api.RestApi) Privy {
	return &PrivyUtils{
		restApi: restApi,
	}
}

func (p *PrivyUtils) Registration(in request.RegisterPrivyRequest, merchantKey string, privyUrl string, privyUserName string, privyPassword string) (out response.PrivyAPIResponse, err error) {

	identityJson, _ := json.Marshal(in.Identity)

	/* SET HEADER */
	requestHeader := map[string]string{
		"Content-Type": "multipart/form-data",
		"Merchant-Key": merchantKey,
	}

	/* SET FILE MULTIPART */
	var requestFile []rest_api.FileParams

	fileKtp := rest_api.FileParams{}
	fileKtp.File = in.KtpByte
	fileKtp.FileName = in.Ktp
	fileKtp.ParamName = "ktp"

	fileKtpSelfie := rest_api.FileParams{}
	fileKtpSelfie.File = in.KtpSelfieByte
	fileKtpSelfie.FileName = in.KtpSelfie
	fileKtpSelfie.ParamName = "selfie"

	requestFile = append(requestFile, fileKtp)
	requestFile = append(requestFile, fileKtpSelfie)

	/* SET FORM DATA */
	requestFormData := map[string]string{
		"email":    in.Email,
		"phone":    in.PhoneNumber,
		"identity": string(identityJson),
	}

	params := rest_api.Params{
		Url:               privyUrl + constant.PathPrivyRegister,
		BasicAuthUsername: privyUserName,
		BasicAuthPassword: privyPassword,
		RequestHeader:     requestHeader,
		FormData:          requestFormData,
		File:              requestFile,
		Resty:             resty.New(),
	}
	resp, err := p.restApi.HitApiService(constant.PostRegisterPrivy, params)

	respMarshal, _ := json.Marshal(resp)

	if err := json.Unmarshal(respMarshal, &out); err != nil {
		log.Println("Could not decode body:", err)
		return out, err
	}

	return
}

func (p *PrivyUtils) UploadDocumentToPrivy(in request.UploadDocumentPrivyRequest, merchantKey string, privyUrl string, privyUserName string, privyPassword string) (out response.PrivyAPIResponse, err error) {

	/* SET HEADER */
	requestHeader := map[string]string{
		"Content-Type": "multipart/form-data",
		"Merchant-Key": merchantKey,
	}

	/* SET FORM DATA */
	owner, _ := json.Marshal(in.Owner)
	recipients, _ := json.Marshal(in.Recipients)
	requestFormData := map[string]string{
		"documentTitle": in.DocumentTitle,
		"docType":       constant.DocTypeParallel,
		"owner":         string(owner),
		"recipients":    string(recipients),
	}

	/* SET FILE MULTIPART */
	var requestFile []rest_api.FileParams

	fileReq := rest_api.FileParams{}
	fileReq.File = in.Document
	fileReq.FileName = in.DocumentTitle
	fileReq.ParamName = constant.ParamNameDocument

	requestFile = append(requestFile, fileReq)

	params := rest_api.Params{
		Url:               privyUrl + constant.PathPrivyUploadDocument,
		BasicAuthUsername: privyUserName,
		BasicAuthPassword: privyPassword,
		RequestHeader:     requestHeader,
		FormData:          requestFormData,
		File:              requestFile,
		Resty:             resty.New(),
	}
	resp, err := p.restApi.HitApiService(constant.PostRegisterPrivy, params)

	if err != nil && err.Error() == "" {
		respMarshal, _ := json.Marshal(resp)
		_ = json.Unmarshal(respMarshal, &out)
		return
	}

	respMarshal, _ := json.Marshal(resp)
	_ = json.Unmarshal(respMarshal, &out)

	return

}

func (p *PrivyUtils) CheckRegistrationStatus(privyUrl string, privyUserName string, privyPassword string, merchantKey string, privyRegistrationToken string) (out response.PrivyCheckRegistrationStatusResponse, err error) {
	requestHeader := map[string]string{
		"Merchant-Key": merchantKey,
		"Content-Type": "multipart/form-data",
	}

	req := request.PrivyCheckRegistrationStatusRequest{
		Token: privyRegistrationToken,
	}
	var requestBody map[string]string
	payload, _ := json.Marshal(req)
	err = json.Unmarshal(payload, &requestBody)

	params := rest_api.Params{
		Url:               privyUrl,
		RequestHeader:     requestHeader,
		BasicAuthUsername: privyUserName,
		BasicAuthPassword: privyPassword,
		FormData:          requestBody,
		Resty:             resty.New(),
	}

	resp, err := p.restApi.HitApiService(http.MethodPost, params)
	if err != nil {
		respMarshal, _ := json.Marshal(resp)
		_ = json.Unmarshal(respMarshal, &out)
		return
	}
	respMarshal, _ := json.Marshal(resp)
	_ = json.Unmarshal(respMarshal, &out)

	return
}

func (p *PrivyUtils) CheckDocumentStatus(privyUrl string, privyUserName string, privyPassword string, merchantKey string, privyDocumentToken string) (out response.PrivyCheckDocumentStatusResponse, err error) {
	requestHeader := map[string]string{
		"Merchant-Key": merchantKey,
	}

	id := fmt.Sprintf("/%s", privyDocumentToken)

	params := rest_api.Params{
		Url:               privyUrl + id,
		RequestHeader:     requestHeader,
		BasicAuthUsername: privyUserName,
		BasicAuthPassword: privyPassword,
		Resty:             resty.New(),
	}

	resp, err := p.restApi.HitApiService(http.MethodGet, params)
	if err != nil {
		respMarshal, _ := json.Marshal(resp)
		_ = json.Unmarshal(respMarshal, &out)
		return
	}
	respMarshal, _ := json.Marshal(resp)
	_ = json.Unmarshal(respMarshal, &out)
	return
}

func (p *PrivyUtils) RegistrationCicil(in request.CicilRegisterPrivyRequest, merchantKey string, privyUrl string, privyUserName string, privyPassword string) (out response.PrivyAPIResponse, err error) {

	identityJson, _ := json.Marshal(in.Identity)

	/* SET HEADER */
	requestHeader := map[string]string{
		"Content-Type": "multipart/form-data",
		"Merchant-Key": merchantKey,
	}

	/* SET FILE MULTIPART */
	var requestFile []rest_api.FileParams

	fileKtp := rest_api.FileParams{}
	fileKtp.File = in.KtpByte
	fileKtp.FileName = in.Ktp
	fileKtp.ParamName = "ktp"

	fileKtpSelfie := rest_api.FileParams{}
	fileKtpSelfie.File = in.KtpSelfieByte
	fileKtpSelfie.FileName = in.KtpSelfie
	fileKtpSelfie.ParamName = "selfie"

	requestFile = append(requestFile, fileKtp)
	requestFile = append(requestFile, fileKtpSelfie)

	/* SET FORM DATA */
	requestFormData := map[string]string{
		"email":    in.Email,
		"phone":    in.PhoneNumber,
		"identity": string(identityJson),
	}

	params := rest_api.Params{
		Url:               privyUrl + constant.PathPrivyRegister,
		BasicAuthUsername: privyUserName,
		BasicAuthPassword: privyPassword,
		RequestHeader:     requestHeader,
		FormData:          requestFormData,
		File:              requestFile,
		Resty:             resty.New(),
	}
	resp, err := p.restApi.HitApiService(constant.PostRegisterPrivy, params)

	respMarshal, _ := json.Marshal(resp)

	if err := json.Unmarshal(respMarshal, &out); err != nil {
		log.Println("Could not decode body:", err)
		return out, err
	}

	return
}
