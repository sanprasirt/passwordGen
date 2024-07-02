/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"math/rand"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "A brief description of your command",
	Long: `Generate random password with cutomizablr options.
	usage of using your command. For example:

passwordGen generate -l 12 -d -s`,
	Run: func(cmd *cobra.Command, args []string) {
		length, _ := cmd.Flags().GetInt("length")
		isDigits, _ := cmd.Flags().GetBool("digits")
		isSymbols, _ := cmd.Flags().GetBool("symbols")
		
		charSet := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
		if isDigits {
			charSet += "0123456789"
		}
		if isSymbols {
			charSet += "!@#$%^&*()_+,./<>?;':\"[]{}"
		}

		password := make([]byte, length)
		for i := range password {
			password[i] = charSet[rand.Intn(len(charSet))]
		}
		cmd.Println(string(password))

	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	generateCmd.Flags().IntP("length", "l", 8, "Length of the password")
	generateCmd.Flags().BoolP("digits", "d", false, "Include digits in the password")
	generateCmd.Flags().BoolP("symbols", "s", false, "Include symbols in the password")
}
