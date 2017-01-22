package models

// Config is configuration loaded from toml
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
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
	Paths     []string // directories to load media from
	CacheSize int      // cache size in bytes (1024 * 1024 = 1 MB)
}
