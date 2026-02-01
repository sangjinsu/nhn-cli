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

var (
	artifactID    int
	serverGroupID int
	concurrentNum int
	nextWhenFail  bool
	deployNote    string
	async         bool
)

func init() {
	cmd.GetRootCmd().AddCommand(DeployCmd)
	DeployCmd.AddCommand(executeCmd)

	executeCmd.Flags().IntVar(&artifactID, "artifact-id", 0, "아티팩트 ID (필수)")
	executeCmd.Flags().IntVar(&serverGroupID, "server-group-id", 0, "서버 그룹 ID (필수)")
	executeCmd.Flags().IntVar(&concurrentNum, "concurrent-num", 0, "동시 실행 수")
	executeCmd.Flags().BoolVar(&nextWhenFail, "next-when-fail", false, "실패 시 다음 진행")
	executeCmd.Flags().StringVar(&deployNote, "deploy-note", "", "배포 메모")
	executeCmd.Flags().BoolVar(&async, "async", false, "비동기 실행")

	executeCmd.MarkFlagRequired("artifact-id")
	executeCmd.MarkFlagRequired("server-group-id")
}

func runExecute(c *cobra.Command, args []string) error {
	deployClient, err := deploy.NewClient(cmd.GetProfile(), cmd.GetDebug())
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
