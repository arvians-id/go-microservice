package config

import "github.com/spf13/viper"

type Config struct {
	Port                string `mapstructure:"PORT"`
	DBConnection        string `mapstructure:"DB_CONNECTION"`
	DBHost              string `mapstructure:"DB_HOST"`
	DBPort              string `mapstructure:"DB_PORT"`
	DBUsername          string `mapstructure:"DB_USERNAME"`
	DBPassword          string `mapstructure:"DB_PASSWORD"`
	DBDatabase          string `mapstructure:"DB_DATABASE"`
	DBSSLMode           string `mapstructure:"DB_SSL_MODE"`
	DBPoolMin           int    `mapstructure:"DATABASE_POOL_MIN"`
	DBPoolMax           int    `mapstructure:"DATABASE_POOL_MAX"`
	DBMaxIdleTimeSecond int    `mapstructure:"DATABASE_MAX_IDLE_TIME_SECOND"`
	DBMaxLifeTimeSecond int    `mapstructure:"DATABASE_MAX_LIFE_TIME_SECOND"`
	JwtSecretKey        string `mapstructure:"JWT_SECRET_KEY"`
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./internal/config/envs")
	viper.SetConfigName("dev")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)

	return
}
