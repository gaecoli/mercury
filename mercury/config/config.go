package config

type GlobalConfig struct {
	App      App      `mapstructure:"app" json:"app" yaml:"app"`
	Log      Log      `mapstructure:"log" json:"log" yaml:"log"`
	Database Database `mapstructure:"database" json:"database" yaml:"database"`
}
