package connectors

import "github.com/flexibleir/orchestrator/data/logindetails"

// BaseConnector - class for all connectors
type BaseConnector interface {
	Run(*logindetails.LoginDetails) string
}
