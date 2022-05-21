package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

const docs = `aadtoken [scope]

Exchanges scope for a JWT token.
`

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	os.Exit(0)
}

func run() error {
	if len(os.Args) != 2 {
		fmt.Print(docs)
		return nil
	}

	scope := os.Args[1]

	credential, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return fmt.Errorf("failed to instantiate credentials: %w", err)
	}
	token, err := credential.GetToken(context.Background(), policy.TokenRequestOptions{
		Scopes: []string{scope},
	})
	if err != nil {
		return fmt.Errorf("failed to retrieve the token: %w", err)
	}

	fmt.Fprintln(os.Stdout, token.Token)
	return nil
}
