package config

type DatabaseConfig struct {
	Type     string // e.g., "postgres", "inmemory"
	Host     string
	User     string
	Password string
	SSLMode  string
	DBName   string
}

var InMemory DatabaseConfig = DatabaseConfig{Type: "inmemory", Host: "N/A", User: "", Password: ""}

var Postgres DatabaseConfig = DatabaseConfig{Type: "postgres", DBName: "YourDBName", Host: "localhost", User: "postgres", Password: "yourpassword", SSLMode: "disable"}