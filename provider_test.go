package main

import (
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"testing"
)

func TestCounterProvider(t *testing.T) {
	pact := dsl.Pact{
		Consumer: "CounterClient",
		Provider: "CounterApi",
	}
	go StartCounterApi()
	_, err := pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL:            "http://localhost:3000/",
		PactURLs:                   []string{"https://eneskzlcn.pactflow.io/pacts/provider/CounterApi/consumer/CounterClient/latest"},
		PublishVerificationResults: true,
		ProviderVersion:            "1.0.0",
		BrokerToken:                "L0IzB6WxiCRX7sEdAQoWlQ",
	})
	if err != nil {
		t.Fatal(err)
	}
}
