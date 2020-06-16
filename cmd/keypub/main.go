package main

import (
	"fmt"
	"os"

	"github.com/go-logr/zapr"
	"github.com/spf13/cobra"
	logf "github.com/theSuess/keypub/pkg/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var LogLevel int
var LogEncoder string

var rootCmd = &cobra.Command{
	Use:   "keypub",
	Short: "Keypub is the place to discover and verify public keys",
	Long:  `A Keyserver focused on usability and identity verification`,
	PersistentPreRun: func(c *cobra.Command, args []string) {
		var config zap.Config
		switch LogEncoder {
		case "console":
			config = zap.NewDevelopmentConfig()
			config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		case "json":
			config = zap.NewProductionConfig()
		}
		config.Level.SetLevel(zapcore.Level(-LogLevel))
		if config.Level.Level() < -1 {
			config.Sampling = nil
		}
		zapLog, err := config.Build()
		if err != nil {
			panic(fmt.Sprintf("Error during log initialization: (%v)?", err))
		}
		logf.SetLogger(zapr.NewLogger(zapLog))
	},
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&LogLevel, "log-level", "v", 0, "log level increasing in verbosity. INFO=0,DEBUG=1,TRACE=2")
	rootCmd.PersistentFlags().StringVar(&LogEncoder, "log-encoder", "json", "set the log encoder. can be either `json` or `console`")
}

func main() {
	rootCmd.AddCommand(serveCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
