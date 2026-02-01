package gamebase

import (
	"encoding/json"
	"fmt"
	"os"

	"nhncli/cmd"
	gb "nhncli/internal/gamebase"

	"github.com/spf13/cobra"
)

var launchingCmd = &cobra.Command{
	Use:   "launching",
	Short: "론칭 상태 조회",
	RunE:  runLaunching,
}

func init() {
	GamebaseCmd.AddCommand(launchingCmd)
}

func runLaunching(c *cobra.Command, args []string) error {
	appKey, _ := c.Flags().GetString("app-key")
	secretKey, _ := c.Flags().GetString("secret-key")
	opts := gb.ClientOption{AppKey: appKey, SecretKey: secretKey}
	gbClient, err := gb.NewClient(cmd.GetProfile(), cmd.GetDebug(), opts)
	if err != nil {
		return err
	}

	info, err := gbClient.GetLaunching()
	if err != nil {
		return err
	}

	if cmd.GetOutput() == "json" {
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		return enc.Encode(info)
	}

	fmt.Printf("론칭 상태: %s (코드: %d)\n", info.Status.Name, info.Status.Code)
	return nil
}
