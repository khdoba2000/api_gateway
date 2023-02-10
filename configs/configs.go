package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

const (
	// DebugMode indicates service mode is debug.
	DebugMode = "debug"
	// TestMode indicates service mode is test.
	TestMode = "test"
	// ReleaseMode indicates service mode is release.
	ReleaseMode = "release"
)

type Config struct {
	ServiceName string
	Environment string // debug, test, release
	Version     string

	HTTPPort   string
	HTTPScheme string

	DefaultOffset string
	DefaultLimit  string

	SettingsServiceHost string
	SettingsGRPCPort    string

	ObjectServiceHost string
	ObjectGRPCPort    string

	AuthServiceHost string
	AuthGRPCPort    string
}

// Load ...
func Load() Config {

	envFileName := cast.ToString(getOrReturnDefaultValue("ENV_FILE_PATH", "../.env"))

	if err := godotenv.Load(envFileName); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.ServiceName = cast.ToString(getOrReturnDefaultValue("SERVICE_NAME", "qv_admin_api_gateway"))
	config.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", DebugMode))
	config.Version = cast.ToString(getOrReturnDefaultValue("VERSION", "1.0"))

	config.HTTPPort = cast.ToString(getOrReturnDefaultValue("ADMIN_API_GETEWAY_HTTP_PORT", ":8080"))
	config.HTTPScheme = cast.ToString(getOrReturnDefaultValue("HTTP_SCHEME", "http"))

	config.DefaultOffset = cast.ToString(getOrReturnDefaultValue("DEFAULT_OFFSET", "0"))
	config.DefaultLimit = cast.ToString(getOrReturnDefaultValue("DEFAULT_LIMIT", "10"))

	config.SettingsServiceHost = cast.ToString(getOrReturnDefaultValue("SETTINGS_SERVICE_HOST", "0.0.0.0"))
	config.SettingsGRPCPort = cast.ToString(getOrReturnDefaultValue("SETTINGS_GRPC_PORT", ":9101"))

	config.ObjectServiceHost = cast.ToString(getOrReturnDefaultValue("OBJECT_SERVICE_HOST", "0.0.0.0"))
	config.ObjectGRPCPort = cast.ToString(getOrReturnDefaultValue("OBJECT_GRPC_PORT", ":9102"))

	config.AuthServiceHost = cast.ToString(getOrReturnDefaultValue("AUTH_SERVICE_HOST", "0.0.0.0"))
	config.AuthGRPCPort = cast.ToString(getOrReturnDefaultValue("AUTH_GRPC_PORT", ":9103"))

	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
