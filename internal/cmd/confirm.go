package cmd

import "github.com/isul/tossinvest-cli/internal/confirm"

func confirmRequiredYes(yes bool, desc string) error {
	return confirm.Required(desc, yes)
}
