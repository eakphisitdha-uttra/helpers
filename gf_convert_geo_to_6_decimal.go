package helpers

import (
	"fmt"
	"strconv"
)

func ConvertGeoTo6Decimal(str string) (string, error) {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%.6f", f), nil
}
