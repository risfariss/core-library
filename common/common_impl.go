package common

import (
	"bitbucket.org/kawancicil/core-library/constant"
	"crypto/sha256"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type CommonUtils struct {
	/*todo fill this block if need connection with another interface*/
}

func InitCommonUtils() Common {
	return &CommonUtils{}
}

func (c *CommonUtils) DateConverterFormatYYMMDD(dateUnix int64) (out string) {
	return time.Unix(dateUnix, 0).Format(constant.TimeTemplateFormatYYYYmmDD)
}

func (c *CommonUtils) DateConvertFromUnixToTime(unix int64) (out time.Time) {
	return time.Unix(unix, 0)
}

func (c *CommonUtils) DateDiff(firstDate time.Time, secondDate time.Time) int {
	diff := secondDate.Sub(firstDate)
	return int(diff.Hours() / 24)
}

func (c *CommonUtils) DateConverterFromTimeFormatYYMMDD(time time.Time) (out string) {
	return time.Format(constant.TimeTemplateFormatYYYYmmDD)
}

func (c *CommonUtils) DateConverterFromStringFormatDDmmYYToTime(dateString string) (out time.Time) {
	date, err := time.Parse(constant.TimeTemplateFormatDDmmYYYY, dateString)
	if err != nil {
		now := time.Date(1970, 1, 1, 1, 1, 1, 1, time.UTC)
		date, _ = time.Parse(constant.TimeTemplateFormatDDmmYYYY, now.Format(constant.TimeTemplateFormatDDmmYYYY))
	}
	out = date
	return
}

func (c *CommonUtils) DateConverterFromStringFormatYYmmDDToTime(dateString string) (out time.Time) {
	date, err := time.Parse(constant.TimeTemplateFormatYYYYmmDD, dateString)
	if err != nil {
		now := time.Date(1970, 1, 1, 1, 1, 1, 1, time.UTC)
		date, _ = time.Parse(constant.TimeTemplateFormatYYYYmmDD, now.Format(constant.TimeTemplateFormatYYYYmmDD))
	}
	out = date
	return
}

func (c *CommonUtils) TimestampToDateInBahasa(dateTimestamp int64, separator string) string {
	tm := time.Unix(dateTimestamp, 0)
	day := strconv.Itoa(tm.Day())
	month := constant.GetMonthInBahasa(tm.Month().String())
	year := strconv.Itoa(tm.Year())
	return day + separator + month + separator + year
}

func (c *CommonUtils) GetDayInBahasa(dateTimestamp int64) string {
	tm := time.Unix(dateTimestamp, 0)
	day := constant.GetDaysInBahasa(tm.Weekday().String())
	return day
}

func (c *CommonUtils) DateUnixToDDmmmmYYYY(dateUnix int64) (out string) {
	tm := time.Unix(dateUnix, 0)
	day := strconv.Itoa(tm.Day())
	month := constant.GET_MONTH_ID[tm.Month().String()]
	year := strconv.Itoa(tm.Year())
	return day + " " + month + " " + year
}

func (c *CommonUtils) GetDayIdFromDateUnix(dateUnix int64) string {
	tm := time.Unix(dateUnix, 0)
	day := constant.GET_DAY_ID[tm.Weekday().String()]
	return day
}

func (c *CommonUtils) FormatCommas(num string) string {
	re := regexp.MustCompile("(\\d+)(\\d{3})")
	for n := ""; n != num; {
		n = num
		num = re.ReplaceAllString(num, "$1.$2")
	}
	return num
}

func (c *CommonUtils) IsValidPhoneNumber(in string) (out string, err error) {
	phoneNumberRegex := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-. \\/]?)?((?:\(?\d+\)?[\-. \\/]?)*)(?:[\-. \\/]?(?:#|ext\.?|extension|x)[\-. \\/]?(\d+))?$`)
	var phone string
	if validateCountryCode := strings.HasPrefix(in, constant.CountryCode); validateCountryCode {
		phone = strings.TrimPrefix(in, constant.CountryCode)
	} else if isPrefixedByZero := strings.HasPrefix(in, "0"); isPrefixedByZero {
		phone = strings.TrimPrefix(in, "0")
	} else {
		phone = in
	}

	phoneWithPrefix := constant.CountryCode + phone
	if len(phone) < 7 || len(phone) > 19 {
		err = errors.New(constant.ErrorPhoneNumberNotValid)
		return
	}

	valid := phoneNumberRegex.MatchString(phoneWithPrefix)
	if !valid {
		err = errors.New(constant.ErrorPhoneNumberNotValid)
		return
	}
	out = phoneWithPrefix
	return
}

func (c *CommonUtils) CheckIsNumber(in string) bool {
	_, err := strconv.Atoi(in)
	if err != nil {
		return false
	} else {
		return true
	}
}

func (c *CommonUtils) ValidatePhoneNumber(phoneNumber string) string {
	var phone string
	if validateCountryCode := strings.HasPrefix(phoneNumber, constant.CountryCode); validateCountryCode {
		phone = strings.TrimPrefix(phoneNumber, constant.CountryCode)
		if isPrefixedByZero := strings.HasPrefix(phone, "0"); isPrefixedByZero {
			phone = strings.TrimPrefix(phone, "0")
		} else {
			phone = strings.TrimPrefix(phoneNumber, constant.CountryCode)
		}
	} else if isPrefixedByCountryCode := strings.HasPrefix(phoneNumber, constant.CountryCodeWithoutPlus); isPrefixedByCountryCode {
		phone = strings.TrimPrefix(phoneNumber, constant.CountryCodeWithoutPlus)
		if isPrefixedByZero := strings.HasPrefix(phone, "0"); isPrefixedByZero {
			phone = strings.TrimPrefix(phone, "0")
		} else {
			phone = strings.TrimPrefix(phoneNumber, constant.CountryCodeWithoutPlus)
		}
	} else if isPrefixedByZero := strings.HasPrefix(phoneNumber, "0"); isPrefixedByZero {
		phone = strings.TrimPrefix(phoneNumber, "0")
	} else {
		phone = phoneNumber
	}

	phoneWithPrefix := constant.CountryCode + phone

	return phoneWithPrefix
}

func (c *CommonUtils) IsValidEmail(in string) (valid bool, err error) {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if len(in) < 3 || len(in) > 50 {
		err = errors.New(constant.ErrorEmailNotValid)
		return
	}
	if !emailRegex.MatchString(in) {
		err = errors.New(constant.ErrorEmailNotValid)
		return
	}
	parts := strings.Split(in, "@")
	mx, err := net.LookupMX(parts[1])
	if err != nil || len(mx) == 0 {
		err = errors.New(constant.ErrorEmailNotValid)
		return
	}
	valid = true
	return
}

func (c *CommonUtils) GetAge(dateOfBirth int64) (age string) {
	birthDate := time.Unix(dateOfBirth, 0)
	today := time.Now()
	ages := today.Year() - birthDate.Year()
	age = strconv.Itoa(ages)
	return age
}

// Get Date with time 00:00:00
func (c *CommonUtils) GetDateTimeStart(date int64) time.Time {
	tm := time.Unix(date, 0)
	year, month, day := tm.Date()
	result := time.Date(year, month, day, 0, 0, 0, 0, tm.Location())
	return result
}

// Get Date with time 23:59:59
func (c *CommonUtils) GetDateTimeEnd(date int64) time.Time {
	tm := time.Unix(date, 0)
	year, month, day := tm.Date()
	result := time.Date(year, month, day, 23, 59, 59, 0, tm.Location())
	return result
}

func (c *CommonUtils) IsValidNIK(in string) (out string, err error) {
	if len(in) != 16 {
		err = errors.New(constant.ErrorInvalidNIK)
		return
	}

	out = in
	return
}

func (c *CommonUtils) SHA256Encrypt(input []byte) (out [32]byte) {
	out = sha256.Sum256(input)
	return
}

func (c *CommonUtils) SHA256Decrypt(input string) (out [32]byte) {
	out = sha256.Sum256([]byte(input))
	return
}

func (c *CommonUtils) MatchPassword(userPassword string, password string) (isValid bool, err error) {
	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(userPassword))
	if err != nil {
		return false, err
	}
	return true, nil
}

func (c *CommonUtils) Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func (c *CommonUtils) EncryptPassword(in string) (out string, err error) {
	hashedPassword, err := c.Hash(in)
	if err != nil {
		log.Println("Error while encrypting password", err)
		return
	}
	out = string(hashedPassword)
	return
}

func (c *CommonUtils) EncryptMPIN(in string) (out string, err error) {
	hashedMpin, err := c.Hash(in)
	if err != nil {
		log.Println("Error whie encrypting mpin", err)
		return "", err
	}
	out = string(hashedMpin)
	return
}

func (c *CommonUtils) MatchMpin(userPin string, in string) (isValid bool) {
	err := bcrypt.CompareHashAndPassword([]byte(userPin), []byte(in))
	if err != nil {
		return false
	}
	return true
}

func (c *CommonUtils) RupiahFormatter(amount float64) string {
	return "Rp" + c.FormatCommas(strconv.FormatFloat(amount, 'f', 0, 64))
}
