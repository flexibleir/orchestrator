package connectors

import "github.com/flexibleir/orchestrator/data/logindetails"

// PcIConnector - PcI connector for communicating with PcI container
type PcIConnector struct{}

// Run - Implementation of interface method baseconnector
func (pci *PcIConnector) Run(*logindetails.LoginDetails) string {
	return ""
}
