package blockstorage

import (
	"nhncli/internal/blockstorage"
	"nhncli/internal/output"

	"github.com/spf13/cobra"
)

var typeCmd = &cobra.Command{
	Use:   "type",
	Short: "볼륨 타입 관리",
}

var typeListCmd = &cobra.Command{
	Use:   "list",
	Short: "볼륨 타입 목록 조회",
	RunE:  runTypeList,
}

func init() {
	BlockstorageCmd.AddCommand(typeCmd)
	typeCmd.AddCommand(typeListCmd)
}

func runTypeList(cmd *cobra.Command, args []string) error {
	client, err := blockstorage.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	types, err := client.ListVolumeTypes()
	if err != nil {
		return err
	}

	if GetOutput() == "json" {
		return output.PrintJSON(types)
	}

	headers := []string{"ID", "NAME"}
	rows := make([][]string, len(types))
	for i, t := range types {
		rows[i] = []string{t.ID, t.Name}
	}

	return output.Print(GetOutput(), &output.TableData{
		Headers: headers,
		Rows:    rows,
	}, types)
}
