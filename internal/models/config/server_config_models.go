package configmodels

type Server struct {
	Host       string `mapstructure:"host"`
	Port       int    `mapstructure:"port"`
	Address    string `mapstructure:"address"`
	Version    string `mapstructure:"version"`
	PrefixPath string `mapstructure:"prefix_path"`
}
