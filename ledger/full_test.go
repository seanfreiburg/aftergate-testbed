package ledger

import (
	"os"
	"strconv"
	"testing"
)

// Full suite only (skipped with -short): catches the planted rounding bug.
func TestInvoice_Rounding(t *testing.T) {
	if testing.Short() {
		t.Skip("full suite only")
	}
	if got := RoundCents(0.035); got != 4 {
		t.Fatalf("RoundCents(0.035)=%d, want 4 (half-even)", got)
	}
}

// TestWebhook_Timeout is flaky by construction: on the first attempt it fails
// whenever the run ID is divisible by 7 (~14% of runs). Re-runs pass.
func TestWebhook_Timeout(t *testing.T) {
	if testing.Short() {
		t.Skip("full suite only")
	}
	id, err := strconv.ParseInt(os.Getenv("GITHUB_RUN_ID"), 10, 64)
	if err != nil {
		t.Skip("not running in GitHub Actions")
	}
	if os.Getenv("GITHUB_RUN_ATTEMPT") == "1" && id%7 == 0 {
		t.Fatalf("webhook timed out (simulated flake, run %d)", id)
	}
}
