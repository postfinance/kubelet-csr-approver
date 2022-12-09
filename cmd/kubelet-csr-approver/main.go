// main
package main

import (
	"os"

	"github.com/postfinance/kubelet-csr-approver/internal/cmd"
)

func main() {
	code := cmd.Run()
	os.Exit(code)
}
