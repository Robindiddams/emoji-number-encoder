package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Robindiddams/emoji-number-encoder/emojinumber"
	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "ene",
	Short: "encode a phone number as emojis",
	Run: func(cmd *cobra.Command, args []string) {
		if decode, _ := cmd.Flags().GetBool("decode"); decode {
			fmt.Println(emojinumber.Decode(args[0]))
		} else {
			i, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("bad number entered")
				os.Exit(1)
			}
			emojinum := emojinumber.Encode(i)
			fmt.Println(emojinum)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.aaa.yaml)")

	rootCmd.Flags().BoolP("decode", "d", false, "to decode some emojis")
}
