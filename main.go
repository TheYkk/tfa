package main

import (
	"bufio"
	"flag"
	"fmt"
	go2fa "github.com/theykk/2fa-go"
	"os"
)

func main() {
	key := ""
	flag.StringVar(&key, "key", "", "--key GEZDGNBV (Base32 encoded secret)")
	flag.Parse()
	// If user didn't provide the secret key with flag, get from stdin
	if key == "" {
		// Stdin scanner for forwarded outputs like `echo "GEZDGNBV" | tfa`
		scanner := bufio.NewScanner(os.Stdin)
		// Wait for input
		for scanner.Scan() {
			key = scanner.Text()
		}
	}
	// Print info to stderr so when we copy the output of cli we won't take the stderr data
	fmt.Fprint(os.Stderr, "Your 2fa code: ")
	fmt.Print(go2fa.GetTOTPToken(key))

}
