package deploy

import (
	"fmt"

	"nhncli/cmd"
	"nhncli/internal/deploy"

	"github.com/spf13/cobra"
)

var DeployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy 서비스 관리",
	Long:  "NHN Cloud Deploy 서비스를 관리합니다.",
}

var executeCmd = &cobra.Command{
	Use:   "execute",
	Short: "배포 실행",
	RunE:  runExecute,
}

var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "바이너리 업로드",
	Long:  "아티팩트에 바이너리 파일을 업로드합니다.",
	RunE:  runUpload,
}

var (
	artifactID    int
	serverGroupID int
	concurrentNum int
	nextWhenFail  bool
	deployNote    string
	async         bool

	// upload flags
	uploadArtifactID    int
	uploadBinaryGroupKey int
	uploadType          string
	uploadFile          string
	uploadVersion       string
	uploadDescription   string
	uploadOsType        string
	uploadMetaFile      string
	uploadFix           bool
)

func init() {
	cmd.GetRootCmd().AddCommand(DeployCmd)
	DeployCmd.AddCommand(executeCmd)
	DeployCmd.AddCommand(uploadCmd)
	DeployCmd.PersistentFlags().String("app-key", "", "Deploy AppKey (프로필 설정 오버라이드)")

	executeCmd.Flags().IntVar(&artifactID, "artifact-id", 0, "아티팩트 ID (필수)")
	executeCmd.Flags().IntVar(&serverGroupID, "server-group-id", 0, "서버 그룹 ID (필수)")
	executeCmd.Flags().IntVar(&concurrentNum, "concurrent-num", 0, "동시 실행 수")
	executeCmd.Flags().BoolVar(&nextWhenFail, "next-when-fail", false, "실패 시 다음 진행")
	executeCmd.Flags().StringVar(&deployNote, "deploy-note", "", "배포 메모")
	executeCmd.Flags().BoolVar(&async, "async", false, "비동기 실행")

	executeCmd.MarkFlagRequired("artifact-id")
	executeCmd.MarkFlagRequired("server-group-id")

	uploadCmd.Flags().IntVar(&uploadArtifactID, "artifact-id", 0, "아티팩트 ID (필수)")
	uploadCmd.Flags().IntVar(&uploadBinaryGroupKey, "binary-group-key", 0, "바이너리 그룹 키 (필수)")
	uploadCmd.Flags().StringVar(&uploadType, "type", "", "애플리케이션 타입: client 또는 server (필수)")
	uploadCmd.Flags().StringVar(&uploadFile, "file", "", "업로드할 파일 경로 (필수)")
	uploadCmd.Flags().StringVar(&uploadVersion, "version", "", "바이너리 버전 (최대 100자)")
	uploadCmd.Flags().StringVar(&uploadDescription, "description", "", "설명")
	uploadCmd.Flags().StringVar(&uploadOsType, "os-type", "", "OS 타입: iOS, Android, etc (client 타입일 때)")
	uploadCmd.Flags().StringVar(&uploadMetaFile, "meta-file", "", "iOS plist 파일 경로")
	uploadCmd.Flags().BoolVar(&uploadFix, "fix", false, "client 바이너리 fix 플래그")

	uploadCmd.MarkFlagRequired("artifact-id")
	uploadCmd.MarkFlagRequired("binary-group-key")
	uploadCmd.MarkFlagRequired("type")
	uploadCmd.MarkFlagRequired("file")
}

func runExecute(c *cobra.Command, args []string) error {
	appKey, _ := c.Flags().GetString("app-key")
	opts := deploy.ClientOption{AppKey: appKey}
	deployClient, err := deploy.NewClient(cmd.GetProfile(), cmd.GetDebug(), opts)
	if err != nil {
		return err
	}

	req := &deploy.DeployExecuteRequest{
		ArtifactID:    artifactID,
		ServerGroupID: serverGroupID,
		ConcurrentNum: concurrentNum,
		NextWhenFail:  nextWhenFail,
		DeployNote:    deployNote,
	}

	result, err := deployClient.ExecuteDeploy(req)
	if err != nil {
		return err
	}

	if result != nil {
		fmt.Printf("배포가 실행되었습니다. (ID: %d, 상태: %s)\n", result.DeploymentID, result.Status)
	} else {
		fmt.Println("배포가 요청되었습니다.")
	}

	return nil
}

func runUpload(c *cobra.Command, args []string) error {
	appKey, _ := c.Flags().GetString("app-key")
	opts := deploy.ClientOption{AppKey: appKey}
	deployClient, err := deploy.NewClient(cmd.GetProfile(), cmd.GetDebug(), opts)
	if err != nil {
		return err
	}

	req := &deploy.BinaryUploadRequest{
		ArtifactID:      uploadArtifactID,
		BinaryGroupKey:  uploadBinaryGroupKey,
		ApplicationType: uploadType,
		BinaryFile:      uploadFile,
		Version:         uploadVersion,
		Description:     uploadDescription,
		OsType:          uploadOsType,
		MetaFile:        uploadMetaFile,
		Fix:             uploadFix,
	}

	result, err := deployClient.UploadBinary(req)
	if err != nil {
		return err
	}

	if result != nil {
		fmt.Printf("바이너리가 업로드되었습니다.\n")
		fmt.Printf("  Binary Key:   %s\n", result.BinaryKey)
		fmt.Printf("  Download URL: %s\n", result.DownloadUrl)
	} else {
		fmt.Println("바이너리 업로드가 완료되었습니다.")
	}

	return nil
}
