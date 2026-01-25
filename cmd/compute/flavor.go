package compute

import (
	"fmt"

	"nhncli/internal/compute"
	"nhncli/internal/output"

	"github.com/spf13/cobra"
)

var flavorCmd = &cobra.Command{
	Use:   "flavor",
	Short: "인스턴스 타입 조회",
}

var flavorListCmd = &cobra.Command{
	Use:   "list",
	Short: "인스턴스 타입 목록 조회",
	RunE:  runFlavorList,
}

var flavorDescribeCmd = &cobra.Command{
	Use:   "describe <flavor-id>",
	Short: "인스턴스 타입 상세 조회",
	Args:  cobra.ExactArgs(1),
	RunE:  runFlavorDescribe,
}

func init() {
	ComputeCmd.AddCommand(flavorCmd)
	flavorCmd.AddCommand(flavorListCmd)
	flavorCmd.AddCommand(flavorDescribeCmd)
}

func runFlavorList(cmd *cobra.Command, args []string) error {
	client, err := compute.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	flavors, err := client.ListFlavors()
	if err != nil {
		return err
	}

	if GetOutput() == "json" {
		return output.PrintJSON(flavors)
	}

	headers := []string{"ID", "NAME", "VCPUS", "RAM (MB)", "DISK (GB)"}
	rows := make([][]string, len(flavors))
	for i, f := range flavors {
		rows[i] = []string{
			f.ID,
			f.Name,
			fmt.Sprintf("%d", f.VCPUs),
			fmt.Sprintf("%d", f.RAM),
			fmt.Sprintf("%d", f.Disk),
		}
	}

	return output.Print(GetOutput(), &output.TableData{
		Headers: headers,
		Rows:    rows,
	}, flavors)
}

func runFlavorDescribe(cmd *cobra.Command, args []string) error {
	flavorID := args[0]

	client, err := compute.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	f, err := client.GetFlavor(flavorID)
	if err != nil {
		return err
	}

	if GetOutput() == "json" {
		return output.PrintJSON(f)
	}

	fmt.Printf("Flavor ID:   %s\n", f.ID)
	fmt.Printf("Name:        %s\n", f.Name)
	fmt.Printf("VCPUs:       %d\n", f.VCPUs)
	fmt.Printf("RAM:         %d MB\n", f.RAM)
	fmt.Printf("Disk:        %d GB\n", f.Disk)
	fmt.Printf("Ephemeral:   %d GB\n", f.Ephemeral)
	fmt.Printf("Public:      %v\n", f.IsPublic)

	return nil
}
