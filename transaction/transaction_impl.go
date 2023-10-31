package transaction

import (
	"strconv"
	"time"
)

type TransactionUtils struct {
	/*todo fill this block if need connection with another interface*/
}

func InitTransactionUtils() Transaction {
	return &TransactionUtils{
	}
}


func(t *TransactionUtils) GenerateClientRef(seq int) string {
	var result = ""
	var sequence = ""
	dateToday := time.Now().Format("020106")

	if seq < 10 {
		sequence = "000" + strconv.Itoa(seq)
	} else if seq < 100 {
		sequence = "00" + strconv.Itoa(seq)
	} else if seq < 1000 {
		sequence = "0" + strconv.Itoa(seq)
	} else if seq < 10000 {
		sequence = "" + strconv.Itoa(seq)
	}
	result = "1" + dateToday + sequence
	return result
}