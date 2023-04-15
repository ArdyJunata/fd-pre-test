package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func LoadConfig(file string) {
	err := godotenv.Load(file)
	if err != nil {
		panic(err)
	}
}

func GetString(key ConfigKey, defaultValue string) string {
	val := os.Getenv(string(key))
	if val == "" {
		return defaultValue
	}
	return val
}

func GetUint8(key ConfigKey, defaultValue uint8) uint8 {
	val := os.Getenv(string(key))
	if val == "" {
		return defaultValue
	}

	valInt64, err := strconv.Atoi(val)
	if err != nil {
		log.Println("error when parse config key", key, "to uint8 with error", err.Error())
		return defaultValue
	}

	valUint8 := uint8(valInt64)
	if valUint8 == 0 {
		return defaultValue
	}

	return valUint8
}
