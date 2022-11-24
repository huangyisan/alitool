package test

import (
	"github.com/spf13/viper"
)

// getEnv let every test function can get ENV by viper
func GetEnv() {
	viper.AutomaticEnv() // read in environment variables that match
	viper.SetEnvPrefix("ALITOOL")
}
