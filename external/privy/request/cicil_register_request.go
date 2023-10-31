package request

type CicilRegisterPrivyRequest struct {
	Identity      CicilIdentityRegisterRequest
	Email         string
	PhoneNumber   string
	Ktp           string
	KtpByte       []byte
	KtpSelfie     string
	KtpSelfieByte []byte
}

type CicilIdentityRegisterRequest struct {
	Nik         string `json:"nik"`
	Name        string `json:"nama"`
	DateOfBirth string `json:"tanggalLahir"`
}
