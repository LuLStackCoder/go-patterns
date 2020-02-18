package account

import (
	`fmt`
	"testing"
)

var acc account = account{"Alexei", "89482441243"}

var testString string = `{
	"nam": "Alexei",
	"phone": "89482441243"
}`

func testInfo(t *testing.T) {
	got := acc.info()
	fmt.Println(got)
	if got != testString {
		t.Errorf("Test Info FAILED: wrong output")
	}
}
