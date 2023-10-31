package privy

import (
	"bitbucket.org/kawancicil/core-library/external/privy/request"
	"bitbucket.org/kawancicil/core-library/external/privy/response"
)

type Privy interface {
	Registration(in request.RegisterPrivyRequest, merchantKey string, privyUrl string, privyUserName string, privyPassword string) (out response.PrivyAPIResponse, err error)
	RegistrationCicil(in request.CicilRegisterPrivyRequest, merchantKey string, privyUrl string, privyUserName string, privyPassword string) (out response.PrivyAPIResponse, err error)
	UploadDocumentToPrivy(in request.UploadDocumentPrivyRequest, merchantKey string, privyUrl string, privyUserName string, privyPassword string) (out response.PrivyAPIResponse, err error)
	CheckRegistrationStatus(privyUrl string, privyUserName string, privyPassword string, merchantKey string, privyRegistrationToken string) (out response.PrivyCheckRegistrationStatusResponse, err error)
	CheckDocumentStatus(privyUrl string, privyUserName string, privyPassword string, merchantKey string, privyDocumentToken string) (out response.PrivyCheckDocumentStatusResponse, err error)
}
