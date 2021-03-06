package datetime

import "github.com/MontFerret/ferret/pkg/runtime/core"

func NewLib() map[string]core.Function {
	return map[string]core.Function{
		"NOW":            Now,
		"DATE":           Date,
		"DATE_DAYOFWEEK": DateDayOfWeek,
		"DATE_YEAR":      DateYear,
		"DATE_MONTH":     DateMonth,
		"DATE_DAY":       DateDay,
	}
}
