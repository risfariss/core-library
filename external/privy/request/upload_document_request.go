package request

type UploadDocumentPrivyRequest struct {
	Owner         Owner
	Recipients    []Recipients
	DocumentTitle string
	DocType       string
	Document      []byte
}

type Owner struct {
	PrivyId         string `json:"privyId"`
	EnterpriseToken string `json:"enterpriseToken"`
}

type Recipients struct {
	PrivyId         string `json:"privyId"`
	EnterpriseToken string `json:"enterpriseToken"`
	Type            string `json:"type"`
}

