package utils

import (
	"fmt"
	"time"
)

func GenerateAccountNumber(id int64) string {
	year := time.Now().Year()
	return fmt.Sprintf("BANK%d%06d", year, id)
}
