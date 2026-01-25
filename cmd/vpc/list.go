package vpc

import (
	"nhncli/internal/output"
	"nhncli/internal/vpc"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "VPC 목록 조회",
	RunE:  runList,
}

func init() {
	VpcCmd.AddCommand(listCmd)
}

func runList(cmd *cobra.Command, args []string) error {
	client, err := vpc.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	vpcs, err := client.ListVPCs()
	if err != nil {
		return err
	}

	if GetOutput() == "json" {
		return output.PrintJSON(vpcs)
	}

	headers := []string{"ID", "NAME", "CIDR", "STATE"}
	rows := make([][]string, len(vpcs))
	for i, v := range vpcs {
		rows[i] = []string{v.ID, v.Name, v.CIDRv4, v.State}
	}

	return output.Print(GetOutput(), &output.TableData{
		Headers: headers,
		Rows:    rows,
	}, vpcs)
}
