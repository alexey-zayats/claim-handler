package cmd

import (
	"fmt"
	"github.com/alexey-zayats/claim-handler/internal/util"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "test",
	Long:  "test",
	Run:   testMain,
}

func init() {
	rootCmd.AddCommand(testCmd)
}

func testMain(cmd *cobra.Command, args []string) {

	inn := []int64{
		233302501109,
		6140006405,
		232303039460,
		23110720778,
		233610495883,
		7728881149,
	}

	ogrn := []int64{
		320237500010156,
		1176196008328,
		317237500237245,
		306231114200083,
		316237000055971,
		1147746651898,
	}

	for _, d := range inn {
		fmt.Printf("INN(%d) is Valid(%v)\n", d, util.CheckINN(d))
	}

	for _, d := range ogrn {
		fmt.Printf("OGRN(%d) is Valid(%v)\n", d, util.CheckOGRN(d))
	}
}
