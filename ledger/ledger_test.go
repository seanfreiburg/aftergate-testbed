package ledger

import "testing"

// Fast suite: runs on every PR (go test -short). Does not catch rounding bugs.
func TestRoundCents_Whole(t *testing.T) {
	if got := RoundCents(1.00); got != 100 {
		t.Fatalf("RoundCents(1.00)=%d, want 100", got)
	}
}
