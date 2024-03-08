package cmd

import (
	"log"
	"os"
	"strings"

	"equipment-management/config"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:   "equipment-management",
		Short: "equipment-management",
		Long:  `equipment-management`,
	}
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(serverCmd)
}

func Execute() error {
	return rootCmd.Execute()
}

func initConfig() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("error while reading config file: %s", err.Error())
	}
	for _, env := range viper.AllKeys() {
		if viper.GetString(env) != "" {
			_ = os.Setenv(env, viper.GetString(env))
			_ = os.Setenv(strings.ToUpper(env), viper.GetString(env))
		}
	}

	config.InitConfig()
}
