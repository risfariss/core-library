package response

type PrivyAPIResponse struct {
	Code    float64      `json:"code"`
	Errors  []ErrorPrivy `json:"errors"`
	Message string       `json:"message"`
	Data    PrivyAPIData `json:"data"`
}

type ErrorPrivy struct {
	Field    string   `json:"field"`
	Messages []string `json:"messages"`
}

type PrivyAPIData struct {
	UserToken   string `json:"userToken,omitempty"`
	Email       string `json:"email,omitempty"`
	Phone       string `json:"phone,omitempty"`
	Status      string `json:"status,omitempty"`
	DocToken    string `json:"docToken,omitempty"`
	UrlDocument string `json:"urlDocument,omitempty"`
	Recipients  []struct {
		PrivyId         string `json:"privyId,omitempty"`
		Type            string `json:"type,omitempty"`
		EnterpriseToken string `json:"enterpriseToken,omitempty"`
		SignatoryStatus string `json:"signatoryStatus,omitempty"`
	} `json:"recipients,omitempty"`
	Download struct {
		Url       string `json:"url"`
		ExpiredAt string `json:"expiredAt"`
	} `json:"download,omitempty"`
	DocumentStatus string `json:"documentStatus,omitempty"`
}
