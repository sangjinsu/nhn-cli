package objectstorage

import (
	"fmt"

	"nhncli/internal/objectstorage"
	"nhncli/internal/output"

	"github.com/spf13/cobra"
)

var containerCmd = &cobra.Command{
	Use:   "container",
	Short: "컨테이너 관리",
}

var containerListCmd = &cobra.Command{
	Use:   "list",
	Short: "컨테이너 목록 조회",
	RunE:  runContainerList,
}

var containerDescribeCmd = &cobra.Command{
	Use:               "describe <container-name>",
	Short:             "컨테이너 메타데이터 조회",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: completeContainerNames,
	RunE:              runContainerDescribe,
}

var containerCreateCmd = &cobra.Command{
	Use:   "create <container-name>",
	Short: "컨테이너 생성",
	Args:  cobra.ExactArgs(1),
	RunE:  runContainerCreate,
}

var containerDeleteCmd = &cobra.Command{
	Use:               "delete <container-name>",
	Short:             "컨테이너 삭제",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: completeContainerNames,
	RunE:              runContainerDelete,
}

func init() {
	ObjectstorageCmd.AddCommand(containerCmd)
	containerCmd.AddCommand(containerListCmd)
	containerCmd.AddCommand(containerDescribeCmd)
	containerCmd.AddCommand(containerCreateCmd)
	containerCmd.AddCommand(containerDeleteCmd)
}

func runContainerList(cmd *cobra.Command, args []string) error {
	client, err := objectstorage.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	containers, err := client.ListContainers()
	if err != nil {
		return err
	}

	if GetOutput() == "json" {
		return output.PrintJSON(containers)
	}

	headers := []string{"NAME", "OBJECTS", "BYTES"}
	rows := make([][]string, len(containers))
	for i, c := range containers {
		rows[i] = []string{
			c.Name,
			fmt.Sprintf("%d", c.Count),
			fmt.Sprintf("%d", c.Bytes),
		}
	}

	return output.Print(GetOutput(), &output.TableData{
		Headers: headers,
		Rows:    rows,
	}, containers)
}

func runContainerDescribe(cmd *cobra.Command, args []string) error {
	name := args[0]

	client, err := objectstorage.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	meta, err := client.GetContainerMetadata(name)
	if err != nil {
		return err
	}

	if GetOutput() == "json" {
		return output.PrintJSON(meta)
	}

	fmt.Printf("Container:    %s\n", name)
	fmt.Printf("Object Count: %s\n", meta.ObjectCount)
	fmt.Printf("Bytes Used:   %s\n", meta.BytesUsed)
	if meta.ReadACL != "" {
		fmt.Printf("Read ACL:     %s\n", meta.ReadACL)
	}
	if meta.WriteACL != "" {
		fmt.Printf("Write ACL:    %s\n", meta.WriteACL)
	}

	return nil
}

func runContainerCreate(cmd *cobra.Command, args []string) error {
	name := args[0]

	client, err := objectstorage.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	if err := client.CreateContainer(name); err != nil {
		return err
	}

	fmt.Printf("컨테이너 %s가 생성되었습니다.\n", name)
	return nil
}

func runContainerDelete(cmd *cobra.Command, args []string) error {
	name := args[0]

	client, err := objectstorage.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	if err := client.DeleteContainer(name); err != nil {
		return err
	}

	fmt.Printf("컨테이너 %s가 삭제되었습니다.\n", name)
	return nil
}

func completeContainerNames(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	if len(args) != 0 {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
	client, err := objectstorage.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
	containers, err := client.ListContainers()
	if err != nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
	var completions []string
	for _, c := range containers {
		completions = append(completions, c.Name)
	}
	return completions, cobra.ShellCompDirectiveNoFileComp
}
