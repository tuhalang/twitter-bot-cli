package cmd

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/joho/godotenv"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/tuhalang/twitter-bot/cmd/rules"
)

var (
	// EnvFile Used for flags.
	envFile string

	rootCmd = &cobra.Command{
		Use:   "twitter-bot",
		Short: "Twitter-Bot is a bot interaction with twitter account",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to Twitter-Bot")
		},
	}
)

// Execute cmd
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&envFile, "env", "", "enviroment file (default is $HOME/.env)")
	rootCmd.AddCommand(streamCmd)
	rules.RegisterRootCmd(rootCmd)
}

func initConfig() {
	if envFile == "" {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		envFile = path.Join(home, ".env")
	}

	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatal("Cannot locate file enviroment !")
	}

}
