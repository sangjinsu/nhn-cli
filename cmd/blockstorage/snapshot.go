package blockstorage

import (
	"fmt"

	"nhncli/internal/blockstorage"
	"nhncli/internal/output"

	"github.com/spf13/cobra"
)

var (
	snapshotVolumeID string
	snapshotName     string
	snapshotDesc     string
	snapshotForce    bool
)

var snapshotCmd = &cobra.Command{
	Use:   "snapshot",
	Short: "스냅샷 관리",
}

var snapshotListCmd = &cobra.Command{
	Use:   "list",
	Short: "스냅샷 목록 조회",
	RunE:  runSnapshotList,
}

var snapshotDescribeCmd = &cobra.Command{
	Use:               "describe <snapshot-id>",
	Short:             "스냅샷 상세 조회",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: completeSnapshotIDs,
	RunE:              runSnapshotDescribe,
}

var snapshotCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "스냅샷 생성",
	Long: `볼륨의 스냅샷을 생성합니다.

예시:
  nhn blockstorage snapshot create \
    --volume-id <volume-id> \
    --name my-snapshot`,
	RunE: runSnapshotCreate,
}

var snapshotDeleteCmd = &cobra.Command{
	Use:               "delete <snapshot-id>",
	Short:             "스냅샷 삭제",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: completeSnapshotIDs,
	RunE:              runSnapshotDelete,
}

func init() {
	BlockstorageCmd.AddCommand(snapshotCmd)
	snapshotCmd.AddCommand(snapshotListCmd)
	snapshotCmd.AddCommand(snapshotDescribeCmd)
	snapshotCmd.AddCommand(snapshotCreateCmd)
	snapshotCmd.AddCommand(snapshotDeleteCmd)

	snapshotCreateCmd.Flags().StringVar(&snapshotVolumeID, "volume-id", "", "볼륨 ID (필수)")
	snapshotCreateCmd.Flags().StringVar(&snapshotName, "name", "", "스냅샷 이름")
	snapshotCreateCmd.Flags().StringVar(&snapshotDesc, "description", "", "스냅샷 설명")
	snapshotCreateCmd.Flags().BoolVar(&snapshotForce, "force", false, "사용 중인 볼륨 강제 스냅샷")
	snapshotCreateCmd.MarkFlagRequired("volume-id")
}

func runSnapshotList(cmd *cobra.Command, args []string) error {
	client, err := blockstorage.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	snapshots, err := client.ListSnapshots()
	if err != nil {
		return err
	}

	if GetOutput() == "json" {
		return output.PrintJSON(snapshots)
	}

	headers := []string{"ID", "NAME", "STATUS", "SIZE(GB)", "VOLUME ID"}
	rows := make([][]string, len(snapshots))
	for i, s := range snapshots {
		rows[i] = []string{
			s.ID,
			s.Name,
			s.Status,
			fmt.Sprintf("%d", s.Size),
			s.VolumeID,
		}
	}

	return output.Print(GetOutput(), &output.TableData{
		Headers: headers,
		Rows:    rows,
	}, snapshots)
}

func runSnapshotDescribe(cmd *cobra.Command, args []string) error {
	snapshotID := args[0]

	client, err := blockstorage.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	snap, err := client.GetSnapshot(snapshotID)
	if err != nil {
		return err
	}

	if GetOutput() == "json" {
		return output.PrintJSON(snap)
	}

	fmt.Printf("Snapshot ID:  %s\n", snap.ID)
	fmt.Printf("Name:         %s\n", snap.Name)
	fmt.Printf("Description:  %s\n", snap.Description)
	fmt.Printf("Status:       %s\n", snap.Status)
	fmt.Printf("Size:         %d GB\n", snap.Size)
	fmt.Printf("Volume ID:    %s\n", snap.VolumeID)
	fmt.Printf("Created:      %s\n", snap.CreatedAt.Format("2006-01-02 15:04:05"))

	return nil
}

func runSnapshotCreate(cmd *cobra.Command, args []string) error {
	client, err := blockstorage.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	req := &blockstorage.SnapshotCreateRequest{
		Snapshot: blockstorage.SnapshotCreateBody{
			VolumeID: snapshotVolumeID,
		},
	}

	if snapshotName != "" {
		req.Snapshot.Name = snapshotName
	}
	if snapshotDesc != "" {
		req.Snapshot.Description = snapshotDesc
	}
	if snapshotForce {
		req.Snapshot.Force = true
	}

	snap, err := client.CreateSnapshot(req)
	if err != nil {
		return err
	}

	if GetOutput() == "json" {
		return output.PrintJSON(snap)
	}

	fmt.Printf("✅ 스냅샷이 생성되었습니다.\n")
	fmt.Printf("Snapshot ID:  %s\n", snap.ID)
	fmt.Printf("Name:         %s\n", snap.Name)
	fmt.Printf("Volume ID:    %s\n", snap.VolumeID)
	fmt.Printf("Status:       %s\n", snap.Status)

	return nil
}

func runSnapshotDelete(cmd *cobra.Command, args []string) error {
	snapshotID := args[0]

	client, err := blockstorage.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	if err := client.DeleteSnapshot(snapshotID); err != nil {
		return err
	}

	fmt.Printf("✅ 스냅샷 %s가 삭제되었습니다.\n", snapshotID)
	return nil
}

func completeSnapshotIDs(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	if len(args) != 0 {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
	client, err := blockstorage.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
	snapshots, err := client.ListSnapshots()
	if err != nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
	var completions []string
	for _, s := range snapshots {
		completions = append(completions, fmt.Sprintf("%s\t%s", s.ID, s.Name))
	}
	return completions, cobra.ShellCompDirectiveNoFileComp
}
