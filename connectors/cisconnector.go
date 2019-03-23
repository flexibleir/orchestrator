package connectors

import "github.com/flexibleir/orchestrator/data/logindetails"

// CiSConnector - Pci connector for communicating with CiS container
type CiSConnector struct{}

// Run - Implementation of interface method baseconnector
func (cis *CiSConnector) Run(logindetails.LoginDetails) string {
	return ""
}
