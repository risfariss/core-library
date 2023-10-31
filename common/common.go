package common

import "time"

type Common interface {
	DateConverterFormatYYMMDD(dateUnix int64) (out string)
	DateConvertFromUnixToTime(unix int64) (out time.Time)
	DateDiff(firstDate time.Time, secondDate time.Time) int
	DateConverterFromTimeFormatYYMMDD(time time.Time) (out string)
	DateConverterFromStringFormatDDmmYYToTime(dateString string) (out time.Time)
	DateConverterFromStringFormatYYmmDDToTime(dateString string) (out time.Time)
	TimestampToDateInBahasa(dateTimestamp int64, separator string) string
	GetDayInBahasa(dateTimestamp int64) string
	FormatCommas(num string) string
	IsValidPhoneNumber(in string) (out string, err error)
	CheckIsNumber(in string) bool
	ValidatePhoneNumber(phoneNumber string) string
	IsValidEmail(in string) (valid bool, err error)
	GetAge(dateOfBirth int64) (age string)
	GetDateTimeStart(date int64) time.Time
	GetDateTimeEnd(date int64) time.Time
	IsValidNIK(in string) (out string, err error)
	SHA256Encrypt(input []byte) (out [32]byte)
	SHA256Decrypt(input string) (out [32]byte)
	MatchPassword(userPassword string, password string) (isValid bool, err error)
	Hash(password string) ([]byte, error)
	EncryptPassword(in string) (out string, err error)
	EncryptMPIN(in string) (out string, err error)
	MatchMpin(userPin string, in string) (isValid bool)
	DateUnixToDDmmmmYYYY(dateUnix int64) (out string)
	GetDayIdFromDateUnix(dateUnix int64) string
	RupiahFormatter(amount float64) string
}
