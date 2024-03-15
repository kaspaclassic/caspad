package utils

import (
	"fmt"

	"github.com/kaspaclassic/caspad/domain/consensus/utils/constants"
)

// FormatKas takes the amount of leors as uint64, and returns amount of CAS with 8  decimal places
func FormatKas(amount uint64) string {
	res := "                   "
	if amount > 0 {
		res = fmt.Sprintf("%19.8f", float64(amount)/constants.LeorPerPyrin)
	}
	return res
}
