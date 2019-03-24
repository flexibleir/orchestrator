package connectors

import (
	"fmt"
	"log"
	"os"

	"github.com/flexibleir/orchestrator/data/logindetails"
	apiclient "github.com/flexibleir/orchestrator/gen/client/compliance/client"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// PcIConnector - PcI connector for communicating with PcI container
type PcIConnector struct{}

// Run - Implementation of interface method baseconnector
func (pci *PcIConnector) Run(*logindetails.LoginDetails) string {

	// create the transport
	transport := httptransport.New(os.Getenv("TODOLIST_HOST"), "", nil)

	// create the API client, with the transport
	client := apiclient.New(transport, strfmt.Default)

	// to override the host for the default client
	// apiclient.Default.SetTransport(transport)

	// make the request to get all items
	resp, err := client.Compliance.Create()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", resp.Payload)
	return ""
}
