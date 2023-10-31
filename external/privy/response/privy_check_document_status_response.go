package response

type PrivyCheckDocumentStatusResponse struct {
	Code    int                `json:"code"`
	Data    DataDocumentStatus `json:"data"`
	Message string             `json:"message"`
}

type Recipients struct {
	PrivyID         string `json:"privyId"`
	Type            string `json:"type"`
	SignatoryStatus string `json:"signatoryStatus"`
}
type DataDocumentStatus struct {
	DocToken       string                                   `json:"docToken"`
	Recipients     []Recipients                             `json:"recipients"`
	DocumentStatus string                                   `json:"documentStatus"`
	URLDocument    string                                   `json:"urlDocument"`
	Download       PrivyCheckDocumentStatusDownloadResponse `json:"download"`
}

type PrivyCheckDocumentStatusDownloadResponse struct {
	Url       string `json:"url"`
	ExpiredAt string `json:"expiredAt"`
}
