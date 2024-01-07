package singstat

// Define a struct for the plugin connection configuration
type SingStatConfig struct {
	// Add any configuration parameters your plugin needs
}

// ConfigInstance returns a pointer to the connection config struct
func ConfigInstance() interface{} {
	return &SingStatConfig{}
}
