package config

// AppConfig application config
type AppConfig struct {
	Env            string
	DBConfig       DBConfig
	Port           string
	AllowedOrigins []string
	AllowedHeaders []string
	GinMode        string
}

// DBConfig database connection config
type DBConfig struct {
	Host        string
	Port        string `default:"5432"`
	User        string
	Pass        string
	Name        string
	SSLMode     string `default:"disable"`
	AutoMigrate bool
}
