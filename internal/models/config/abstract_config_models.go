package configmodels

type Config struct {
	Server     Server     `mapstructure:"server"`
	Postgresql Postgresql `mapstructure:"postgresql"`
}
