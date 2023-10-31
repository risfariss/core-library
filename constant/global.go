package constant

const (
	//Time Template
	TimeTemplateFormatDDmmYYYY      = "02-01-2006"
	TimeTemplateFormatYYYYmmDD      = "2006-01-02"
	TimeTemplateFormatDDmmYYYhhMMss = "02-Jan-2006 15:04:05"

	HeaderXAuthenticatedUserId = "X-Authenticated-Userid"
	HeaderXRequestId           = "X-REQUEST-ID"
	CountryCode                = "+62"
	CountryCodeWithoutPlus     = "62"

	PostRegisterPrivy = "post_register_privy_rest"

	/* Service privy_rest */
	PathPrivyUploadDocument = "/document/upload"
	PathPrivyRegister       = "/registration"

	DocTypeParallel   = "Parallel"
	ParamNameDocument = "document"

	RabbitMQDialUrl          = "rabbitMQDialUrl"
	RabbitMQQueueHttpRequest = "httpRequest"
)

func GetMonthInBahasa(name string) string {
	mappingObject := map[string]string{
		"January":   "Januari",
		"February":  "Februari",
		"March":     "Maret",
		"April":     "April",
		"May":       "Mei",
		"June":      "Juni",
		"July":      "Juli",
		"August":    "Agustus",
		"September": "September",
		"October":   "Oktober",
		"November":  "November",
		"December":  "Desember",
	}
	result, _ := mappingObject[name]
	return result
}

func GetDaysInBahasa(name string) string {
	mappingObject := map[string]string{
		"Sunday":    "Minggu",
		"Monday":    "Senin",
		"Tuesday":   "Selasa",
		"Wednesday": "Rabu",
		"Thursday":  "Kamis",
		"Friday":    "Jumat",
		"Saturday":  "Sabtu",
	}
	result, _ := mappingObject[name]
	return result
}

var GET_MONTH_ID = map[string]string{
	"January":   "Januari",
	"February":  "Februari",
	"March":     "Maret",
	"April":     "April",
	"May":       "Mei",
	"June":      "Juni",
	"July":      "Juli",
	"August":    "Agustus",
	"September": "September",
	"October":   "Oktober",
	"November":  "November",
	"December":  "Desember",
}

var GET_DAY_ID = map[string]string{
	"Sunday":    "Minggu",
	"Monday":    "Senin",
	"Tuesday":   "Selasa",
	"Wednesday": "Rabu",
	"Thursday":  "Kamis",
	"Friday":    "Jumat",
	"Saturday":  "Sabtu",
}
