package response

type PrivyCheckRegistrationStatusResponse struct {
	Code    int    `json:"code"`
	Data    Data   `json:"data"`
	Message string `json:"message"`
}

type Data struct {
	PrivyID     string   `json:"privyId"`
	Email       string   `json:"email"`
	Phone       string   `json:"phone"`
	ProcessedAt string   `json:"processedAt"`
	UserToken   string   `json:"userToken"`
	Status      string   `json:"status"`
	Identity    Identity `json:"identity"`
	Reject      Reject   `json:"reject"`
}

type Identity struct {
	Nama         string `json:"nama"`
	Nik          string `json:"nik"`
	TanggalLahir string `json:"tanggalLahir"`
	TempatLahir  string `json:"tempatLahir"`
}

type Reject struct {
	Code     string     `json:"code"`
	Reason   string     `json:"reason"`
	Handlers []Handlers `json:"handlers"`
}

type Handlers struct {
	Category    string                   `json:"category"`
	Handler     string                   `json:"handler"`
	FileSupport []map[string]interface{} `json:"file_support"`
}
