package connectors

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/flexibleir/orchestrator/test/testutils"
)

func TestPciConnectorRun(t *testing.T) {
	loginDetails := testutils.Credential
	var connector BaseConnector
	connector = &PcIConnector{}
	returnValue := connector.Run(loginDetails)
	assert.Equal(t, "", returnValue)
}
