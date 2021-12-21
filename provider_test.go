//go:build provider
// +build provider

package main

import (
	"fmt"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"github.com/pact-foundation/pact-go/utils"
	"testing"
)
type Settings struct {
	Host            string
	ProviderName    string
	BrokerBaseURL   string
	BrokerUsername  string // Basic authentication
	BrokerPassword  string // Basic authentication
	ConsumerName    string
	ConsumerVersion string // a git sha, semantic version number
	ProviderVersion string
}
func createPact() (pact *dsl.Pact, cleanUp func()) {
	pact = &dsl.Pact{
		Consumer:                 "CounterClient",
		Provider:                 "CounterApi",
		DisableToolValidityCheck: true,
		PactFileWriteMode:        "merge",
		LogDir:                   "./pacts/logs",
	}

	cleanUp = func() { pact.Teardown() }
	return pact, cleanUp
}
func TestCounterProvider(t *testing.T) {
	pact := dsl.Pact{
		Consumer: "CounterClient",
		Provider: "CounterApi",
		DisableToolValidityCheck: true,
	}
	port,_ := utils.GetFreePort()
	go StartCounterApi(port)
	_,err := pact.VerifyProvider(t,
		types.VerifyRequest{
			ProviderBaseURL:            fmt.Sprintf("http://localhost:%d", port),
			PactURLs:                   []string{"https://eneskzlcn.pactflow.io/pacts/provider/CounterApi/consumer/CounterClient/version/7e13c3f624fa969c4a22ccb8a0500f5127e11d97"},
			PublishVerificationResults: true,
			ProviderVersion:            "1.0.0",
			BrokerToken:                "L0IzB6WxiCRX7sEdAQoWlQ",
			Tags: []string{"main"},
			StateHandlers: map[string]types.StateHandler{
				"counter exist": func() error {
					return nil
				},
				"counter increases": func() error {
					return nil
				},
				"counter decreases": func() error {
					return nil
				},
			},
		})
	if err != nil {
		t.Fatal(err)
	}

}
