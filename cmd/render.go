// %N0 CARR13R█

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func renderSauceFunc(cmd *cobra.Command, args []string) {
	file, _ := os.Open(args[0])
	defer file.Close()

	buffer := make([]byte, 1024)

	newoffset, _ := file.Seek(-128, 2)

	n, err := file.Read(buffer)

	if err != nil {
		panic(err)
	}

	// out the buffer content
	fmt.Println(string(buffer[:n]))
	fmt.Println(string(buffer[0:5]))
}
