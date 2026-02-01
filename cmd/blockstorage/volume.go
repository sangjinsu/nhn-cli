package blockstorage

import (
	"fmt"

	"nhncli/internal/blockstorage"
	"nhncli/internal/output"

	"github.com/spf13/cobra"
)

var (
	volumeName   string
	volumeSize   int
	volumeType   string
	volumeAZ     string
	volumeDesc   string
	volumeSnapID string
)

var volumeCmd = &cobra.Command{
	Use:   "volume",
	Short: "볼륨 관리",
}

var volumeListCmd = &cobra.Command{
	Use:   "list",
	Short: "볼륨 목록 조회",
	RunE:  runVolumeList,
}

var volumeDescribeCmd = &cobra.Command{
	Use:   "describe <volume-id>",
	Short: "볼륨 상세 조회",
	Args:  cobra.ExactArgs(1),
	RunE:  runVolumeDescribe,
}

var volumeCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "볼륨 생성",
	Long: `새로운 볼륨을 생성합니다.

예시:
  nhn blockstorage volume create \
    --name my-volume \
    --size 100 \
    --type SSD`,
	RunE: runVolumeCreate,
}

var volumeDeleteCmd = &cobra.Command{
	Use:   "delete <volume-id>",
	Short: "볼륨 삭제",
	Args:  cobra.ExactArgs(1),
	RunE:  runVolumeDelete,
}

func init() {
	BlockstorageCmd.AddCommand(volumeCmd)
	volumeCmd.AddCommand(volumeListCmd)
	volumeCmd.AddCommand(volumeDescribeCmd)
	volumeCmd.AddCommand(volumeCreateCmd)
	volumeCmd.AddCommand(volumeDeleteCmd)

	volumeCreateCmd.Flags().StringVar(&volumeName, "name", "", "볼륨 이름")
	volumeCreateCmd.Flags().IntVar(&volumeSize, "size", 0, "볼륨 크기 (GB, 필수)")
	volumeCreateCmd.Flags().StringVar(&volumeType, "type", "", "볼륨 타입 (SSD/HDD)")
	volumeCreateCmd.Flags().StringVar(&volumeAZ, "availability-zone", "", "가용성 영역")
	volumeCreateCmd.Flags().StringVar(&volumeDesc, "description", "", "볼륨 설명")
	volumeCreateCmd.Flags().StringVar(&volumeSnapID, "snapshot-id", "", "스냅샷 ID (스냅샷에서 생성)")
	volumeCreateCmd.MarkFlagRequired("size")
}

func runVolumeList(cmd *cobra.Command, args []string) error {
	client, err := blockstorage.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	volumes, err := client.ListVolumes()
	if err != nil {
		return err
	}

	if GetOutput() == "json" {
		return output.PrintJSON(volumes)
	}

	headers := []string{"ID", "NAME", "STATUS", "SIZE(GB)", "TYPE", "AZ"}
	rows := make([][]string, len(volumes))
	for i, v := range volumes {
		rows[i] = []string{
			v.ID,
			v.Name,
			v.Status,
			fmt.Sprintf("%d", v.Size),
			v.VolumeType,
			v.AvailabilityZone,
		}
	}

	return output.Print(GetOutput(), &output.TableData{
		Headers: headers,
		Rows:    rows,
	}, volumes)
}

func runVolumeDescribe(cmd *cobra.Command, args []string) error {
	volumeID := args[0]

	client, err := blockstorage.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	vol, err := client.GetVolume(volumeID)
	if err != nil {
		return err
	}

	if GetOutput() == "json" {
		return output.PrintJSON(vol)
	}

	fmt.Printf("Volume ID:         %s\n", vol.ID)
	fmt.Printf("Name:              %s\n", vol.Name)
	fmt.Printf("Description:       %s\n", vol.Description)
	fmt.Printf("Status:            %s\n", vol.Status)
	fmt.Printf("Size:              %d GB\n", vol.Size)
	fmt.Printf("Volume Type:       %s\n", vol.VolumeType)
	fmt.Printf("Availability Zone: %s\n", vol.AvailabilityZone)
	fmt.Printf("Bootable:          %s\n", vol.Bootable)
	fmt.Printf("Encrypted:         %v\n", vol.Encrypted)
	fmt.Printf("Created:           %s\n", vol.CreatedAt.Format("2006-01-02 15:04:05"))

	if len(vol.Attachments) > 0 {
		fmt.Printf("\nAttachments:\n")
		for _, a := range vol.Attachments {
			fmt.Printf("  - Server: %s, Device: %s\n", a.ServerID, a.Device)
		}
	}

	return nil
}

func runVolumeCreate(cmd *cobra.Command, args []string) error {
	client, err := blockstorage.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	req := &blockstorage.VolumeCreateRequest{
		Volume: blockstorage.VolumeCreateBody{
			Size: volumeSize,
		},
	}

	if volumeName != "" {
		req.Volume.Name = volumeName
	}
	if volumeType != "" {
		req.Volume.VolumeType = volumeType
	}
	if volumeAZ != "" {
		req.Volume.AvailabilityZone = volumeAZ
	}
	if volumeDesc != "" {
		req.Volume.Description = volumeDesc
	}
	if volumeSnapID != "" {
		req.Volume.SnapshotID = volumeSnapID
	}

	vol, err := client.CreateVolume(req)
	if err != nil {
		return err
	}

	if GetOutput() == "json" {
		return output.PrintJSON(vol)
	}

	fmt.Printf("볼륨이 생성되었습니다.\n")
	fmt.Printf("Volume ID:  %s\n", vol.ID)
	fmt.Printf("Name:       %s\n", vol.Name)
	fmt.Printf("Size:       %d GB\n", vol.Size)
	fmt.Printf("Status:     %s\n", vol.Status)

	return nil
}

func runVolumeDelete(cmd *cobra.Command, args []string) error {
	volumeID := args[0]

	client, err := blockstorage.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	if err := client.DeleteVolume(volumeID); err != nil {
		return err
	}

	fmt.Printf("볼륨 %s가 삭제되었습니다.\n", volumeID)
	return nil
}
