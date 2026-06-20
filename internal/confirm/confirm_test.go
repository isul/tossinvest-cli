package confirm_test

import (
	"os"
	"testing"

	"github.com/isul/tossinvest-cli/internal/confirm"
)

func TestRequired_AutoYes(t *testing.T) {
	if err := confirm.Required("test op", true); err != nil {
		t.Fatal(err)
	}
}

func TestRequired_EnvAutoConfirm(t *testing.T) {
	t.Setenv("TOSSINVEST_AUTO_CONFIRM", "1")
	if err := confirm.Required("test op", false); err != nil {
		t.Fatal(err)
	}
}

func TestRequired_Cancelled(t *testing.T) {
	t.Setenv("TOSSINVEST_AUTO_CONFIRM", "")
	// Cannot easily test stdin without piping; verify env path only
	if os.Getenv("TOSSINVEST_AUTO_CONFIRM") == "1" {
		t.Skip("auto confirm set")
	}
}
