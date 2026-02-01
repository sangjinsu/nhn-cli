package objectstorage

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"nhncli/internal/objectstorage"
	"nhncli/internal/output"

	"github.com/spf13/cobra"
)

var (
	objectContainer string
	objectFile      string
	objectOutput    string
	objectName      string
)

var objectCmd = &cobra.Command{
	Use:   "object",
	Short: "오브젝트 관리",
}

var objectListCmd = &cobra.Command{
	Use:   "list",
	Short: "오브젝트 목록 조회",
	RunE:  runObjectList,
}

var objectUploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "오브젝트 업로드",
	Long: `파일을 Object Storage에 업로드합니다.

예시:
  nhn objectstorage object upload --container my-container --file ./test.txt
  nhn os object upload --container my-container --file ./test.txt --name custom-name`,
	RunE: runObjectUpload,
}

var objectDownloadCmd = &cobra.Command{
	Use:               "download <object-name>",
	Short:             "오브젝트 다운로드",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: completeObjectNames,
	RunE:              runObjectDownload,
}

var objectDeleteCmd = &cobra.Command{
	Use:               "delete <object-name>",
	Short:             "오브젝트 삭제",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: completeObjectNames,
	RunE:              runObjectDelete,
}

var objectDescribeCmd = &cobra.Command{
	Use:               "describe <object-name>",
	Short:             "오브젝트 메타데이터 조회",
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: completeObjectNames,
	RunE:              runObjectDescribe,
}

func init() {
	ObjectstorageCmd.AddCommand(objectCmd)
	objectCmd.AddCommand(objectListCmd)
	objectCmd.AddCommand(objectUploadCmd)
	objectCmd.AddCommand(objectDownloadCmd)
	objectCmd.AddCommand(objectDeleteCmd)
	objectCmd.AddCommand(objectDescribeCmd)

	// --container flag for all object subcommands
	objectListCmd.Flags().StringVar(&objectContainer, "container", "", "컨테이너 이름 (필수)")
	objectListCmd.MarkFlagRequired("container")

	objectUploadCmd.Flags().StringVar(&objectContainer, "container", "", "컨테이너 이름 (필수)")
	objectUploadCmd.Flags().StringVar(&objectFile, "file", "", "업로드할 파일 경로 (필수)")
	objectUploadCmd.Flags().StringVar(&objectName, "name", "", "오브젝트 이름 (기본: 파일명)")
	objectUploadCmd.MarkFlagRequired("container")
	objectUploadCmd.MarkFlagRequired("file")

	objectDownloadCmd.Flags().StringVar(&objectContainer, "container", "", "컨테이너 이름 (필수)")
	objectDownloadCmd.Flags().StringVar(&objectOutput, "output-file", "", "저장 경로 (기본: 현재 디렉토리에 오브젝트 이름)")
	objectDownloadCmd.MarkFlagRequired("container")

	objectDeleteCmd.Flags().StringVar(&objectContainer, "container", "", "컨테이너 이름 (필수)")
	objectDeleteCmd.MarkFlagRequired("container")

	objectDescribeCmd.Flags().StringVar(&objectContainer, "container", "", "컨테이너 이름 (필수)")
	objectDescribeCmd.MarkFlagRequired("container")

	// Register container name completion for --container flags
	for _, c := range []*cobra.Command{objectListCmd, objectUploadCmd, objectDownloadCmd, objectDeleteCmd, objectDescribeCmd} {
		c.RegisterFlagCompletionFunc("container", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			client, err := objectstorage.NewClient(GetProfile(), GetRegion(), GetDebug())
			if err != nil {
				return nil, cobra.ShellCompDirectiveNoFileComp
			}
			containers, err := client.ListContainers()
			if err != nil {
				return nil, cobra.ShellCompDirectiveNoFileComp
			}
			var names []string
			for _, ct := range containers {
				names = append(names, ct.Name)
			}
			return names, cobra.ShellCompDirectiveNoFileComp
		})
	}
}

func runObjectList(cmd *cobra.Command, args []string) error {
	client, err := objectstorage.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	objects, err := client.ListObjects(objectContainer)
	if err != nil {
		return err
	}

	if GetOutput() == "json" {
		return output.PrintJSON(objects)
	}

	headers := []string{"NAME", "BYTES", "CONTENT-TYPE", "LAST-MODIFIED"}
	rows := make([][]string, len(objects))
	for i, o := range objects {
		rows[i] = []string{
			o.Name,
			fmt.Sprintf("%d", o.Bytes),
			o.ContentType,
			o.LastModified,
		}
	}

	return output.Print(GetOutput(), &output.TableData{
		Headers: headers,
		Rows:    rows,
	}, objects)
}

func runObjectUpload(cmd *cobra.Command, args []string) error {
	f, err := os.Open(objectFile)
	if err != nil {
		return fmt.Errorf("파일 열기 실패: %w", err)
	}
	defer f.Close()

	name := objectName
	if name == "" {
		name = filepath.Base(objectFile)
	}

	client, err := objectstorage.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	if err := client.UploadObject(objectContainer, name, f, ""); err != nil {
		return err
	}

	fmt.Printf("오브젝트 %s가 컨테이너 %s에 업로드되었습니다.\n", name, objectContainer)
	return nil
}

func runObjectDownload(cmd *cobra.Command, args []string) error {
	objName := args[0]

	client, err := objectstorage.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	body, err := client.DownloadObject(objectContainer, objName)
	if err != nil {
		return err
	}
	defer body.Close()

	outPath := objectOutput
	if outPath == "" {
		outPath = filepath.Base(objName)
	}

	f, err := os.Create(outPath)
	if err != nil {
		return fmt.Errorf("파일 생성 실패: %w", err)
	}
	defer f.Close()

	written, err := io.Copy(f, body)
	if err != nil {
		return fmt.Errorf("파일 쓰기 실패: %w", err)
	}

	fmt.Printf("오브젝트 %s를 %s에 다운로드했습니다. (%d bytes)\n", objName, outPath, written)
	return nil
}

func runObjectDelete(cmd *cobra.Command, args []string) error {
	objName := args[0]

	client, err := objectstorage.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	if err := client.DeleteObject(objectContainer, objName); err != nil {
		return err
	}

	fmt.Printf("오브젝트 %s가 컨테이너 %s에서 삭제되었습니다.\n", objName, objectContainer)
	return nil
}

func runObjectDescribe(cmd *cobra.Command, args []string) error {
	objName := args[0]

	client, err := objectstorage.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	meta, err := client.GetObjectMetadata(objectContainer, objName)
	if err != nil {
		return err
	}

	if GetOutput() == "json" {
		return output.PrintJSON(meta)
	}

	fmt.Printf("Object:         %s\n", objName)
	fmt.Printf("Container:      %s\n", objectContainer)
	fmt.Printf("Content-Length: %s\n", meta.ContentLength)
	fmt.Printf("Content-Type:   %s\n", meta.ContentType)
	fmt.Printf("ETag:           %s\n", meta.ETag)
	fmt.Printf("Last-Modified:  %s\n", meta.LastModified)

	return nil
}

func completeObjectNames(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	if len(args) != 0 {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
	container, _ := cmd.Flags().GetString("container")
	if container == "" {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
	client, err := objectstorage.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
	objects, err := client.ListObjects(container)
	if err != nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
	var names []string
	for _, o := range objects {
		names = append(names, o.Name)
	}
	return names, cobra.ShellCompDirectiveNoFileComp
}
