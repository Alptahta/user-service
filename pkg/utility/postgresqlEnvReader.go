package utility

import "github.com/spf13/viper"

const configName = "postgresql"
const configType = "env"

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type PostgreSQLConfig struct {
	DBDriver string `mapstructure:"DB_DRIVER"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadPostgreSqlConfig(path string) (postgreSQLConfig PostgreSQLConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&postgreSQLConfig)
	return
}
