package request

type RegisterPrivyRequest struct {
	Identity      IdentityRegisterRequest
	Email         string
	PhoneNumber   string
	Ktp           string
	KtpByte       []byte
	KtpSelfie     string
	KtpSelfieByte []byte
}

type IdentityRegisterRequest struct {
	Nik         string `json:"nik"`
	Name        string `json:"name"`
	DateOfBirth string `json:"date_of_birth"`
}
