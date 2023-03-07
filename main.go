package number_diff

import (
	"fmt"
	"github.com/leekchan/accounting"
	"math"
)

type NumberDiffResult struct {
	Diff      float64
	AbsDiff   float64
	PctDiff   float64
	Sign      int8
	localInfo accounting.Locale
	ac        accounting.Accounting
}

func (n *NumberDiffResult) FormatDiffAsDecimal(minFractionDigits int) string {
	fd := n.localInfo.FractionLength

	if minFractionDigits > n.localInfo.FractionLength {
		fd = minFractionDigits
	}

	return accounting.FormatNumberFloat64(n.Diff, fd, n.localInfo.ThouSep, n.localInfo.DecSep)
}

func (n *NumberDiffResult) FormatDiffAsMoney() string {
	return n.ac.FormatMoneyFloat64(n.Diff)
}

// FormatAbsDiffAsDecimal returns the absolute difference as a decimal number, e.g. 1.23. using the localized format (defaults to English).
func (n *NumberDiffResult) FormatAbsDiffAsDecimal(minFractionDigits int) string {
	fd := n.localInfo.FractionLength

	if minFractionDigits > n.localInfo.FractionLength {
		fd = minFractionDigits
	}

	return accounting.FormatNumberFloat64(n.AbsDiff, fd, n.localInfo.ThouSep, n.localInfo.DecSep)
}

func (n *NumberDiffResult) FormatAbsDiffAsMoney() string {
	return n.ac.FormatMoneyFloat64(n.AbsDiff)
}

// FormatPctDiff returns the percentage difference formatted using the localized format (defaults to English).
func (n *NumberDiffResult) FormatPctDiff(minFractionDigits int) string {
	fd := n.localInfo.FractionLength

	if minFractionDigits > n.localInfo.FractionLength {
		fd = minFractionDigits
	}

	f := accounting.FormatNumberFloat64(n.PctDiff*100, fd, n.localInfo.ThouSep, n.localInfo.DecSep)
	return fmt.Sprintf("%v%%", f)
}

// SignSymbol returns a string representation of the sign of the number difference.  No string means no difference.
func (n *NumberDiffResult) SignSymbol() string {
	if n.Sign == 1 {
		return "+"
	}

	if n.Sign == -1 {
		return "-"
	}

	return ""
}

// Diff calculates the difference between two numbers using the default ISO-4217 code "USD".
func Diff(originalValue, newValue float64) (*NumberDiffResult, error) {
	return DiffWithLocale(originalValue, newValue, "USD")
}

// DiffWithLocale calculates the difference between two numbers using a specific language, e.g. en-US.
func DiffWithLocale(originalValue, newValue float64, locale string) (*NumberDiffResult, error) {
	change := newValue - originalValue
	var sign int8 = 0

	if change > 0 {
		sign = 1
	} else if change < 0 {
		sign = -1
	}

	if val, ok := accounting.LocaleInfo[locale]; ok {
		ac := accounting.Accounting{
			Symbol:    val.ComSymbol,
			Precision: val.FractionLength,
			Thousand:  val.ThouSep,
			Decimal:   val.DecSep,
		}

		return &NumberDiffResult{
			ac:        ac,
			localInfo: val,
			Diff:      change,
			AbsDiff:   math.Abs(change),
			PctDiff:   change / originalValue,
			Sign:      sign,
		}, nil
	}

	return nil, fmt.Errorf("locale %s not found", locale)
}
