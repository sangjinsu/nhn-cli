package compute

import (
	"fmt"

	"nhncli/internal/compute"
	"nhncli/internal/output"

	"github.com/spf13/cobra"
)

var imageCmd = &cobra.Command{
	Use:   "image",
	Short: "이미지 조회",
}

var imageListCmd = &cobra.Command{
	Use:   "list",
	Short: "이미지 목록 조회",
	RunE:  runImageList,
}

var imageDescribeCmd = &cobra.Command{
	Use:   "describe <image-id>",
	Short: "이미지 상세 조회",
	Args:  cobra.ExactArgs(1),
	RunE:  runImageDescribe,
}

func init() {
	ComputeCmd.AddCommand(imageCmd)
	imageCmd.AddCommand(imageListCmd)
	imageCmd.AddCommand(imageDescribeCmd)
}

func runImageList(cmd *cobra.Command, args []string) error {
	client, err := compute.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	images, err := client.ListImages()
	if err != nil {
		return err
	}

	if GetOutput() == "json" {
		return output.PrintJSON(images)
	}

	headers := []string{"ID", "NAME", "STATUS", "MIN DISK", "MIN RAM"}
	rows := make([][]string, len(images))
	for i, img := range images {
		rows[i] = []string{
			img.ID,
			img.Name,
			img.Status,
			fmt.Sprintf("%d GB", img.MinDisk),
			fmt.Sprintf("%d MB", img.MinRAM),
		}
	}

	return output.Print(GetOutput(), &output.TableData{
		Headers: headers,
		Rows:    rows,
	}, images)
}

func runImageDescribe(cmd *cobra.Command, args []string) error {
	imageID := args[0]

	client, err := compute.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	img, err := client.GetImage(imageID)
	if err != nil {
		return err
	}

	if GetOutput() == "json" {
		return output.PrintJSON(img)
	}

	fmt.Printf("Image ID:    %s\n", img.ID)
	fmt.Printf("Name:        %s\n", img.Name)
	fmt.Printf("Status:      %s\n", img.Status)
	fmt.Printf("Min Disk:    %d GB\n", img.MinDisk)
	fmt.Printf("Min RAM:     %d MB\n", img.MinRAM)
	fmt.Printf("Created:     %s\n", img.Created.Format("2006-01-02 15:04:05"))
	fmt.Printf("Updated:     %s\n", img.Updated.Format("2006-01-02 15:04:05"))

	if len(img.Metadata) > 0 {
		fmt.Printf("\nMetadata:\n")
		for k, v := range img.Metadata {
			fmt.Printf("  %s: %s\n", k, v)
		}
	}

	return nil
}
