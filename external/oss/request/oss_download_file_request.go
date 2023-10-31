package request

type OssDownloadFileRequest struct {
	DocumentUrl  string `json:"document_url"`
	DocumentPath string `json:"document_path"`
}
