package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DatabaseHost     string `mapstructure:"database_host"`
	DatabasePort     int    `mapstructure:"database_port"`
	DatabasePassword string `mapstructure:"database_password"`
}

// LoadConfig อ่าน Configuration จาก Environment Variable
func LoadConfig() (Config, error) {
	viper.SetEnvPrefix("MYAPP")
	viper.AutomaticEnv()

	// กำหนดค่า Default โดยใช้ Key ที่ตรงกับ Struct Tag
	viper.SetDefault("database_host", "localhost")
	viper.SetDefault("database_port", 5432)
	viper.SetDefault("database_password", "")

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return Config{}, fmt.Errorf("failed to unmarshal config: %v", err)
	}

	if config.DatabaseHost == "" {
		return Config{}, fmt.Errorf("database_host is required but not set")
	}
	return config, nil
}

func main() {
	config, err := LoadConfig()
	if err != nil {
		log.Fatalf("Cannot load config: %v", err)
	}

	fmt.Println("Database Configuration:")
	fmt.Printf("Host: %s\n", config.DatabaseHost)
	fmt.Printf("Port: %d\n", config.DatabasePort)
	fmt.Printf("Password is set: %v\n", config.DatabasePassword != "")
}
