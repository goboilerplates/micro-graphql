package env

import (
	"os"
)

// LoadVariables gets environment variables: enableProdMode and serverHost.
func LoadVariables() (bool, string) {
	var (
		enableProdMode = true
		serverHost     = ":9000"
	)
	if os.Getenv("GBP_PROMODE") == "false" {
		enableProdMode = false
	}

	host := os.Getenv("GBP_HOST")
	if len(host) > 0 {
		serverHost = host
	}
	return enableProdMode, serverHost
}
