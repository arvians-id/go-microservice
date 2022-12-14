package config

import "github.com/spf13/viper"

type Config struct {
	Port                string `mapstructure:"PORT"`
	AuthSvcUrl          string `mapstructure:"AUTH_SERVICE"`
	ProductSvcUrl       string `mapstructure:"PRODUCT_SERVICE"`
	UserSvcUrl          string `mapstructure:"USER_SERVICE"`
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
	AwsAccessKeyId      string `mapstructure:"AWS_ACCESS_KEY_ID"`
	AwsSecretKey        string `mapstructure:"AWS_SECRET_KEY"`
	AwsRegion           string `mapstructure:"AWS_REGION"`
	AwsBucket           string `mapstructure:"AWS_BUCKET"`
}

func LoadConfig(filenames ...string) (c *Config, err error) {
	if filenames != nil {
		viper.AddConfigPath(filenames[0])
	} else {
		viper.AddConfigPath("../../config/envs")
	}

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
