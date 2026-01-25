package compute

import (
	"nhncli/internal/compute"
	"nhncli/internal/output"

	"github.com/spf13/cobra"
)

var azCmd = &cobra.Command{
	Use:   "az",
	Short: "가용성 영역 조회",
}

var azListCmd = &cobra.Command{
	Use:   "list",
	Short: "가용성 영역 목록 조회",
	RunE:  runAZList,
}

func init() {
	ComputeCmd.AddCommand(azCmd)
	azCmd.AddCommand(azListCmd)
}

func runAZList(cmd *cobra.Command, args []string) error {
	client, err := compute.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	azs, err := client.ListAvailabilityZones()
	if err != nil {
		return err
	}

	if GetOutput() == "json" {
		return output.PrintJSON(azs)
	}

	headers := []string{"NAME", "AVAILABLE"}
	rows := make([][]string, len(azs))
	for i, az := range azs {
		available := "No"
		if az.ZoneState.Available {
			available = "Yes"
		}
		rows[i] = []string{az.ZoneName, available}
	}

	return output.Print(GetOutput(), &output.TableData{
		Headers: headers,
		Rows:    rows,
	}, azs)
}
