package loan

import "time"

type Loan interface {
	GenerateLoanNo(seq int, loanCode string,year int, month time.Month, day int) string
	CheckIsFunded(loanStatus int) (isFunded bool, loanStatusId int, loanStatusName string)
	CheckIsLate(loanStatus int, dueDate int64) (isLate bool, loanStatusId int, loanStatusName string)
	CheckTenor(tenor int, tenorPeriod int) (tenorId int, tenorPeriodId int, tenorValue string)
	GenerateLoanNoV2(loanCode string) (out string)
}

