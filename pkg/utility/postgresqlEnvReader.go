package utility

import "github.com/spf13/viper"

const configName = "postgresql"
const configType = "env"
const configPath = "."

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type PostgreSQLConfig struct {
	DBDriver string `mapstructure:"DB_DRIVER"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadPostgreSqlConfig() (postgreSQLConfig *PostgreSQLConfig, err error) {
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&postgreSQLConfig)
	if err != nil {
		return nil, err
	}
	return postgreSQLConfig, nil
}
