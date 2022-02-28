package text

import (
	"testing"

	"rsc.io/quote"
)

func TestSiter(t *testing.T) {
	want := "I can eat glass and it doesn't hurt me."
	got := quote.Glass()
	if got != want {
		t.Error("Didn't get what we wanted")
	}
}
