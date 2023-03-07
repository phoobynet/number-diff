# Number Diff

Calculate the difference between two numbers, including percentage difference, absolute difference and sign.

Some additional functions are available for formatting the results.  To localize the results, specify an [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217) code, e.g. USD, GBP, CHF, AUD, etc.

## Installation

```bash
go get github.com/phoobynet/number-diff
```

## Example

```go
import diff "github.com/phoobynet/number-diff"

func main () {
	originalValue := 100
	newValue := 101
	
	result := diff.DiffWithLocale(originalValue, newValue, "GBP")
	
	fmt.Printf("The difference was: %f\n", result.Diff)
	fmt.Printf("The difference as decimal was: %f\n", result.FormatDiffAsDecimal(2))
	fmt.Printf("The difference as money was: %s\n", result.FormatDiffAsMoney(2))
	fmt.Printf("The absolute difference was: %f\n", result.AbsDiff)
	fmt.Printf("The absolute difference as decimal was: %s\n", result.FormatAbsDiffAsDecimal(2))
	fmt.Printf("The absolute difference as money was: %s\n", result.FormatAbsDiffAsMoney())
	fmt.Printf("The percentage difference was: %f\n", result.PctDiff)
	fmt.Printf("The percentage formatted difference was: %s\n", result.FormatPctDiff(2))
	fmt.Printf("The sign was: %d\n", result.Sign)
	fmt.Printf("The sign symbol was: %s\n", result.SignSymbol())
}
```
