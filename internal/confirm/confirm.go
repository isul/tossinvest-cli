package confirm

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/isul/tossinvest-cli/internal/config"
)

func Required(description string, autoYes bool) error {
	if autoYes || os.Getenv(config.EnvAutoConfirm) == "1" {
		return nil
	}
	fmt.Fprintf(os.Stderr, "Write operation: %s\n", description)
	fmt.Fprintf(os.Stderr, "Type CONFIRM to proceed: ")
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	if strings.TrimSpace(line) != "CONFIRM" {
		return fmt.Errorf("operation cancelled")
	}
	return nil
}
