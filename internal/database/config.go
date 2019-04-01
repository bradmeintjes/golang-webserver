package database

// Config defines database configuration.
type Config struct {
	Driver  string `yaml:"driver"`
	ConnStr string `yaml:"connection"`
}
