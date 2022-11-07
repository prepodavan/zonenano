package cmd

import (
	"fmt"
	"os"

	"github.com/prepodavan/zonenano/appbuild"
	"github.com/prepodavan/zonenano/usecase"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   appbuild.AppName,
	Short: "Prints given timestamp in unix nanoseconds (epoch)",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		calc := usecase.NewUnixTimeCalculator()
		if dateStr == nil || len(*dateStr) == 0 {
			fmt.Println(calc.CalcNanos(cmd.Context(), *zone, *offset))
			return
		}
		nanos, err := calc.CalcNanosForInstant(cmd.Context(), *dateStr, *layout, *zone, *offset)
		if err != nil {
			fmt.Println("error: ", err.Error())
			return
		}
		fmt.Println(nanos)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var (
	offset  *int
	zone    *string
	dateStr *string
	layout  *string
)

func init() {
	offset = rootCmd.Flags().IntP(
		"offset",
		"o",
		0,
		"zone offset in seconds (seconds east of UTC)",
	)
	zone = rootCmd.Flags().StringP(
		"zone",
		"z",
		"UTC",
		"zone name",
	)
	dateStr = rootCmd.Flags().StringP(
		"date",
		"d",
		"",
		"date to be set. now as default",
	)
	layout = rootCmd.Flags().StringP(
		"layout",
		"l",
		"02-01-2006T15:04:05",
		"time layout if date specified",
	)
}
