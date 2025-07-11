package nyx

import (
	"log/slog"

	Nyx "github.com/MatteoMori/nyx/pkg/nyx"
	NyxShared "github.com/MatteoMori/nyx/pkg/shared"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var config NyxShared.Config

var startNyx = &cobra.Command{
	Use:     "start",
	Aliases: []string{"start"},
	Short:   "Start Nyx controller in Quality Score mode",
	Args:    cobra.ExactArgs(0), // 0 arguments
	Run: func(cmd *cobra.Command, args []string) {
		Nyx.Start(config)
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().IntP("verbosity", "v", 0, "verbosity level (0-2)")

	// Viper bindings to Flags
	// This allows Viper to read the flags set by Cobra and use them in the configuration
	viper.BindPFlag("prometheusPort", rootCmd.PersistentFlags().Lookup("prometheusPort"))
	viper.BindPFlag("verbosity", rootCmd.PersistentFlags().Lookup("verbosity"))

	// Viper deafults
	viper.SetDefault("prometheusPort", 9090) // Default port for Prometheus metrics endpoint
	viper.SetDefault("verbosity", 0)

	// Start the Nyx command
	rootCmd.AddCommand(startNyx)
}

/*
	Nyx Config

Here is where we initialize the Nyx configuration.
This configuration is used to set up the App, metrics, and other parameters.
- Ideally, all Nyx parameters should have an equivalent here so that humans can override as they want.
*/
func initConfig() {
	viper.SetConfigName("nyx") // Name of config file without extension
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/nyx") // Config file expected location

	/*
		Viper will read environment variables and use them as configuration values if they match your config keys.
		EXAMPLE:
		 export PROMETHEUSPORT=12345 --> Viper will use the value from the environment variable instead of the default or config file.
	*/
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		slog.Info("Config file loaded", "file", viper.ConfigFileUsed())
	} else {
		slog.Warn("No config file found, using environment variables and defaults", "err", err)
	}

	// Load into config struct
	if err := viper.Unmarshal(&config); err != nil {
		slog.Error("Unable to decode config into struct", "err", err)
	}

	// Apply fallback defaults if any field is missing
	//shared.ApplyDefaultConfig(&config)
}
