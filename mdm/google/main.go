package google

import "github.com/johnmikee/sync2snipe/mdm"

type Config struct{}

// GetDevice implements mdm.Provider.
func (*Config) GetDevice(deviceID string) (*mdm.Device, error) {
	panic("unimplemented")
}

// GetUsers implements mdm.Provider.
func (*Config) GetUsers(opts *mdm.QueryOpts) ([]mdm.User, error) {
	panic("unimplemented")
}

// ListDevices implements mdm.Provider.
func (*Config) ListDevices() ([]mdm.Device, error) {
	panic("unimplemented")
}

// QueryDevices implements mdm.Provider.
func (*Config) QueryDevices(opts *mdm.QueryOpts) (mdm.DeviceResults, error) {
	panic("unimplemented")
}

// Setup implements mdm.Provider.
func (*Config) Setup(config mdm.Config) {
	panic("unimplemented")
}
