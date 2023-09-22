package client

import (
	"github.com/johnmikee/sync2snipe/mdm"
	"github.com/johnmikee/sync2snipe/mdm/google"
	"github.com/johnmikee/sync2snipe/mdm/intune"
	"github.com/johnmikee/sync2snipe/mdm/jamf"
	"github.com/johnmikee/sync2snipe/mdm/kandji"
)

// Config represents the configuration for the client.
type Config struct {
	MDMProvider mdm.Provider
}

// MDM represents the MDM client.
type MDM struct {
	MDM    mdm.MDM
	Config mdm.Config
}

// New creates a new MDM provider based on the provided MDM configuration.
// It returns the MDM provider instance.
func New(m *MDM) mdm.Provider {
	config := Config{
		MDMProvider: createMDMProvider(m.MDM),
	}

	config.MDMProvider.Setup(m.Config)

	return config.MDMProvider
}

// createMDMProvider creates and returns an MDM provider based on the provided MDM type.
func createMDMProvider(providerName mdm.MDM) mdm.Provider {
	switch providerName {
	case mdm.Jamf:
		return &jamf.Config{}
	case mdm.Kandji:
		return &kandji.Config{}
	case mdm.Intune:
		return &intune.Config{}
	case mdm.Google:
		return &google.Config{}
	default:
		return nil
	}
}
