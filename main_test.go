package number_diff

import (
	"fmt"
	"testing"
)

func TestDiffWithLocale(t *testing.T) {
	var testCases = []struct {
		originalValue float64
		newValue      float64
		expected      string
		iso4217Code   string
		expectedError string
	}{
		{100.0, 101.0, "Should Find Locale", "GBP", ""},
		{100.0, 99.99, "Should Fail with unknown Locale", "CH", "locale CH not found"},
	}
	for _, tc := range testCases {
		testName := fmt.Sprintf("[%s]->%v:%v==%s", tc.iso4217Code, tc.originalValue, tc.newValue, tc.expected)

		t.Run(testName, func(t *testing.T) {
			_, err := DiffWithLocale(tc.originalValue, tc.newValue, tc.iso4217Code)

			if err != nil && tc.expectedError == "" {
				t.Errorf("Error should be nil, but was %s", err.Error())
			} else if (err != nil && tc.expectedError != "") && err.Error() != tc.expectedError {
				t.Errorf("Error should be %s, but was %s", tc.expectedError, err.Error())
			}
		})
	}
}

func TestDiffWithLocale_FormatAsMoney(t *testing.T) {
	var testCases = []struct {
		originalValue float64
		newValue      float64
		expected      string
		iso4217Code   string
		expectedError string
	}{
		{100.0, 101.0, "£1.00", "GBP", ""},
		{100.0, 101.0, "$1.00", "USD", ""},
		{100.0, 101.34, "$1.34", "USD", ""},
		{100.0, 101.34, "€1,34", "EUR", ""},
		{100.0, 99.99, "-€0,01", "EUR", ""},
		{100.0, 99.99, "-$0.01", "USD", ""},
		{100.0, 99.99, "-CHF 0.01", "CHF", ""},
	}
	for _, tc := range testCases {
		testName := fmt.Sprintf("[%s]->%v:%v==%s", tc.iso4217Code, tc.originalValue, tc.newValue, tc.expected)

		t.Run(testName, func(t *testing.T) {
			diffResult, err := DiffWithLocale(tc.originalValue, tc.newValue, tc.iso4217Code)

			if err != nil {
				t.Errorf("Error should be nil, but was %s", err.Error())
			}

			actual := diffResult.FormatDiffAsMoney()

			if actual != tc.expected {
				t.Errorf("EXPECTED: \"%s\", but got ACTUAL: \"%s\"", tc.expected, actual)
			}
		})
	}
}

func TestDiffWithLocale_FormatDiffAsDecimal(t *testing.T) {
	var testCases = []struct {
		originalValue       float64
		newValue            float64
		minFractionalDigits int
		expected            string
		iso4217Code         string
		expectedError       string
	}{
		{100.0, 101.0, 2, "1.00", "GBP", ""},
		{100.0, 101.0, 0, "1.00", "GBP", ""},
		{100.0, 101.34, 0, "1.34", "USD", ""},
		{100.0, 101.34, 4, "1,3400", "EUR", ""},
		{100.0, 99.99, 2, "-0,01", "EUR", ""},
		{100.0, 99.99, 2, "-0.01", "USD", ""},
		{100.0, 99.99, 2, "-0.01", "CHF", ""},
	}
	for _, tc := range testCases {
		testName := fmt.Sprintf("[%s]->%v:%v==%s", tc.iso4217Code, tc.originalValue, tc.newValue, tc.expected)

		t.Run(testName, func(t *testing.T) {
			diffResult, err := DiffWithLocale(tc.originalValue, tc.newValue, tc.iso4217Code)

			if err != nil {
				t.Errorf("Error should be nil, but was %s", err.Error())
			}

			actual := diffResult.FormatDiffAsDecimal(tc.minFractionalDigits)

			if actual != tc.expected {
				t.Errorf("EXPECTED: \"%s\", but got ACTUAL: \"%s\"", tc.expected, actual)
			}
		})
	}
}

func TestDiffWithLocale_FormatAbsDiffAsDecimal(t *testing.T) {
	var testCases = []struct {
		originalValue       float64
		newValue            float64
		minFractionalDigits int
		expected            string
		iso4217Code         string
		expectedError       string
	}{
		{100.0, 101.0, 2, "1.00", "GBP", ""},
		{100.0, 101.0, 0, "1.00", "GBP", ""},
		{100.0, 101.34, 0, "1.34", "USD", ""},
		{100.0, 101.34, 4, "1,3400", "EUR", ""},
		{100.0, 99.99, 2, "0,01", "EUR", ""},
		{100.0, 99.99, 2, "0.01", "USD", ""},
		{100.0, 99.99, 2, "0.01", "CHF", ""},
	}
	for _, tc := range testCases {
		testName := fmt.Sprintf("[%s]->%v:%v==%s", tc.iso4217Code, tc.originalValue, tc.newValue, tc.expected)

		t.Run(testName, func(t *testing.T) {
			diffResult, err := DiffWithLocale(tc.originalValue, tc.newValue, tc.iso4217Code)

			if err != nil {
				t.Errorf("Error should be nil, but was %s", err.Error())
			}

			actual := diffResult.FormatAbsDiffAsDecimal(tc.minFractionalDigits)

			if actual != tc.expected {
				t.Errorf("EXPECTED: \"%s\", but got ACTUAL: \"%s\"", tc.expected, actual)
			}
		})
	}
}

func TestDiffWithLocale_FormatPctDiff(t *testing.T) {
	var testCases = []struct {
		originalValue       float64
		newValue            float64
		minFractionalDigits int
		expected            string
		iso4217Code         string
		expectedError       string
	}{
		{100.0, 101.0, 2, "1.00%", "GBP", ""},
		{100.0, 101.0, 0, "1.00%", "GBP", ""},
		{100.0, 101.3499, 0, "1.35%", "USD", ""},
		{100.0, 101.3499, 4, "1,3499%", "EUR", ""},
		{100.0, 99.99, 2, "-0,01%", "EUR", ""},
		{100.0, 99.99, 2, "-0.01%", "USD", ""},
		{100.0, 99.99, 2, "-0.01%", "CHF", ""},
	}

	for _, tc := range testCases {
		testName := fmt.Sprintf("[%s]->%v:%v==%s", tc.iso4217Code, tc.originalValue, tc.newValue, tc.expected)

		t.Run(testName, func(t *testing.T) {
			diffResult, err := DiffWithLocale(tc.originalValue, tc.newValue, tc.iso4217Code)

			if err != nil {
				t.Errorf("Error should be nil, but was %s", err.Error())
			}

			actual := diffResult.FormatPctDiff(tc.minFractionalDigits)

			if actual != tc.expected {
				t.Errorf("EXPECTED: \"%s\", but got ACTUAL: \"%s\"", tc.expected, actual)
			}
		})
	}
}

func TestDiffWithLocale_SignSymbol(t *testing.T) {
	var testCases = []struct {
		originalValue       float64
		newValue            float64
		minFractionalDigits int
		expected            string
		iso4217Code         string
	}{
		{100.0, 101.0, 2, "+", "GBP"},
		{100.0, 50.0, 4, "-", "EUR"},
		{100.0, 100., 4, "", "EUR"},
	}

	for _, tc := range testCases {
		testName := fmt.Sprintf("[%s]->%v:%v==%s", tc.iso4217Code, tc.originalValue, tc.newValue, tc.expected)

		t.Run(testName, func(t *testing.T) {
			diffResult, err := DiffWithLocale(tc.originalValue, tc.newValue, tc.iso4217Code)

			if err != nil {
				t.Errorf("Error should be nil, but was %s", err.Error())
			}

			actual := diffResult.SignSymbol()

			if actual != tc.expected {
				t.Errorf("EXPECTED: \"%s\", but got ACTUAL: \"%s\"", tc.expected, actual)
			}
		})
	}
}

func TestDiff_DefaultsToUSD(t *testing.T) {
	var testCases = []struct {
		originalValue float64
		newValue      float64
		expected      string
	}{
		{100.0, 101.0, "$1.00"},
		{100.0, 101.34, "$1.34"},
	}
	for _, tc := range testCases {
		testName := fmt.Sprintf("%v:%v==%s", tc.originalValue, tc.newValue, tc.expected)

		t.Run(testName, func(t *testing.T) {
			diffResult, err := Diff(tc.originalValue, tc.newValue)

			if err != nil {
				t.Errorf("Error should be nil, but was %s", err.Error())
			}

			actual := diffResult.FormatDiffAsMoney()

			if actual != tc.expected {
				t.Errorf("EXPECTED: \"%s\", but got ACTUAL: \"%s\"", tc.expected, actual)
			}
		})
	}
}
