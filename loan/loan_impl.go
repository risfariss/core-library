package loan

import (
	"bitbucket.org/kawancicil/core-library/common"
	"bitbucket.org/kawancicil/core-library/constant"
	"fmt"
	"strconv"
	"time"
)

type LoanUtils struct {
	common common.Common
}

func InitLoanUtils(common common.Common) Loan {
	return &LoanUtils{
		common: common,
	}
}




func(c *LoanUtils) GenerateLoanNo(seq int, loanCode string,year int, month time.Month, day int) string {
	var result = ""
	var sequence = ""

	var monthString = strconv.Itoa(int(month))
	if int(month) < 10 {
		monthString = "0" + strconv.Itoa(int(month))
	}
	var dayString = strconv.Itoa(day)
	if day < 10 {
		dayString = "0" + strconv.Itoa(day)
	}
	var monthYear = monthString + strconv.Itoa(year)[2:]
	if seq < 10 {
		sequence = "000" + strconv.Itoa(seq)
	} else if seq < 100 {
		sequence = "00" + strconv.Itoa(seq)
	} else if seq < 1000 {
		sequence = "0" + strconv.Itoa(seq)
	} else if seq < 10000 {
		sequence = "" + strconv.Itoa(seq)
	}
	result = "KC-"+ loanCode + "-" + monthYear + "-" + dayString + "-" + sequence
	return result
}


func(c *LoanUtils) CheckIsFunded(loanStatus int) (isFunded bool, loanStatusId int, loanStatusName string) {
	if loanStatus == constant.PartiallyFundedId || loanStatus == constant.FullyFundedId || loanStatus == constant.InreviewId{
		return true, constant.FullyFundedId, constant.FundedName
	}
	return false, 0, ""
}

func(c *LoanUtils) CheckIsLate(loanStatus int, dueDate int64) (isLate bool, loanStatusId int, loanStatusName string) {
	if loanStatus == constant.ActiveId {
		dueDateTime := c.common.DateConvertFromUnixToTime(dueDate)
		now := time.Now()
		late := dueDateTime.Before(now)
		if !late {
			return false, 0, ""
		}
		return true, constant.LateId, constant.LateName
	}
	return false, 0, ""
}

func(c *LoanUtils) CheckTenor(tenor int, tenorPeriod int) (tenorId int, tenorPeriodId int, tenorValue string) {
	tenorMap := map[int]string{
		constant.TenorPeriodDayId:         constant.TenorPeriodDayName,
		constant.TenorPeriodWeekId:        constant.TenorPeriodWeekName,
		constant.TenorPeriodPeriodMonthId: constant.TenorPeriodPeriodMonthName,
		constant.TenorPeriodYearId:        constant.TenorPeriodYearName,
	}
	tenorId = tenor
	tenorPeriodId = tenorPeriod
	tenorValue = fmt.Sprintf(`%d %s`, tenor, tenorMap[tenorPeriod])
	return
}


func(c *LoanUtils) GenerateLoanNoV2(loanCode string) (out string){
	/* new format loan number
		ORD -> order type, yg ada saat ini ORD, TFL, MPO, dll aka loan type
		060102 -> creation date, hari ini 230525
		212345 -> random number -> unix timestamp
	*/
	DDMMYYYY := "060102" /* tahun bulan hari 2006-01-02 */
	now := time.Now().UTC()
	data := now.Format(DDMMYYYY)
	timeNow := time.Now().Unix()
	out = loanCode + "-" + data + "-" + strconv.Itoa(int(timeNow))
	return
}