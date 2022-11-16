package internal

var (
	Configuration *Config
	ServerProfile string
)

// Config properties. All configurations should be described here.
type Config struct {
	ServiceName string `yaml:"serviceName"`

	Version string `yaml:"version"`

	Server struct {
		Profile         string `yaml:"profile"`
		ShutdownTimeout int    `yaml:"shutdownTimeout"`
		HTTP            struct {
			Port string
		} `yaml:"http"`
	} `yaml:"server"`

	Integration struct {
		//RPC struct {
		//	Test RPCClientConfig `yaml:"test"`
		//} `yaml:"rpc"`

		HTTP struct {
			Block HTTPClientConfig `yaml:"block"`
		} `yaml:"http"`
	} `yaml:"integration"`
}

type RPCClientConfig struct {
	URL string
}

// HTTPClientConfig contains parameters for repeated requests to external services.
type HTTPClientConfig struct {
	URL         string `yaml:"url"`
	ContentType string `yaml:"contentType"`
	APIKey      string `env:"GETBLOCK_KEY_API"`
}
