/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
  "math/rand"
  "time"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

var u, n, e bool

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates a new password",
	Long: `This command generates a new password and copies it to the clipboard.`,
	Run: func(cmd *cobra.Command, args []string) {
		generatePassword()
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().BoolVarP(&u, "upper", "u", false, "Add uppercase letters to the password")
	generateCmd.Flags().BoolVarP(&n, "number", "n", false, "Add numbers to the password")
	generateCmd.Flags().BoolVarP(&e, "especial", "e", false, "Add especial characters to the password")
}

func generatePassword() {
	upperCase := []string{
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N",
    "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	}

	numbers := []string{
		"1", "2", "3", "4", "5", "6", "7", "8", "9",
	}

	especialCharacters := []string{
		"!", "@", "#", "$", "%", "&", "^", "*", "-", "=", "+",
	}

	characters := []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n",
    "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",	
	}

	if (u) {
		characters = append(characters, upperCase...)
	}
	
	if (n) {
		characters = append(characters, numbers...)
	}
		
	if (e) {
		characters = append(characters, especialCharacters...)
	}

  stringLength := len(characters)
	rand.Seed(time.Now().UnixNano())

	password := ""
	
	for i := 0; i < 16; i++ {
		randomNumber := rand.Intn(stringLength)

		password += characters[randomNumber]
	}

	clipboard.WriteAll(password)
	fmt.Println("Sua nova senha e: ", password)	
	fmt.Println("E ela ja foi copiada para sua area de transferencia")
}
