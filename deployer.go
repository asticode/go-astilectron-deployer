package astideployer

// Deployer represents an object capable of deploying astilectron projects
type Deployer struct {
	c Configuration
}

// Configuration is the deployer configuration
type Configuration struct {
	ServerPrivate ConfigurationServer `toml:"server_private"`
	ServerPublic  ConfigurationServer `toml:"server_public"`
}

type ConfigurationServer struct {
	Addr          string `toml:"addr"`
	Password      string `toml:"password"`
	ResourcesPath string `toml:"resources_path"`
	Username      string `toml:"username"`
}

// New creates a new deployer
func New(c Configuration) *Deployer {
	return &Deployer{c: c}
}
