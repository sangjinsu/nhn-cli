package vpc

import (
	"fmt"

	"nhncli/internal/vpc"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:               "delete <vpc-id>",
	Short:             "VPC 삭제",
	Args:              cobra.ExactArgs(1),
	RunE:              runDelete,
	ValidArgsFunction: completeVPCIDs,
}

func init() {
	VpcCmd.AddCommand(deleteCmd)
}

func runDelete(cmd *cobra.Command, args []string) error {
	vpcID := args[0]

	client, err := vpc.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	if err := client.DeleteVPC(vpcID); err != nil {
		return err
	}

	fmt.Printf("✅ VPC %s가 삭제되었습니다.\n", vpcID)
	return nil
}
