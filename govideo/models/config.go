package models

// Config is configuration loaded from toml
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Cache    CacheConfig
	App      AppConfig
}

// ServerConfig holds server-specific configurations
type ServerConfig struct {
	Port       int
	StaticPath string
}

// DatabaseConfig holds database-specific configurations
type DatabaseConfig struct {
	URI    string
	DBName string
}

// AppConfig holds application-specific configurations
type AppConfig struct {
	Paths      []string // directories to load media from
	UserExpiry int      `toml:"user_expiry"`
}

// CacheConfig holds redis cache configuration
type CacheConfig struct {
	URI      string
	Password string
	Database string
}
