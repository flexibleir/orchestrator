package connectors

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/flexibleir/orchestrator/gen/client/compliance/models"

	"github.com/flexibleir/orchestrator/data/logindetails"
	apiclient "github.com/flexibleir/orchestrator/gen/client/compliance/client"
	"github.com/flexibleir/orchestrator/gen/client/compliance/client/compliance"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// PcIConnector - PcI connector for communicating with PcI container
type PcIConnector struct{}

// Run - Implementation of interface method baseconnector
func (pci *PcIConnector) Run(login *logindetails.LoginDetails) string {

	// create the transport
	transport := httptransport.New("localhost:80", "/", nil)

	// create the API client, with the transport
	client := apiclient.New(transport, strfmt.Default)

	// to override the host for the default client
	// apiclient.Default.SetTransport(transport)

	complianceType := "PcI"
	body := &models.Createjob{
		Compliancetype: &complianceType,
		Hostname:       &login.HostName,
		Username:       &login.UserName,
		Password:       &login.Password,
	}
	var timeTillContextDeadline = time.Now().Add(time.Hour)
	ctx, _ := context.WithDeadline(context.Background(), timeTillContextDeadline)

	params := &compliance.CreateParams{Body: body, Context: ctx}

	// make the request to get all items
	resp, err := client.Compliance.Create(params)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", resp.Payload)
	return ""
}
