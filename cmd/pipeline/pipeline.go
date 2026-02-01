package pipeline

import (
	"fmt"

	"nhncli/cmd"
	"nhncli/internal/pipeline"

	"github.com/spf13/cobra"
)

var PipelineCmd = &cobra.Command{
	Use:   "pipeline",
	Short: "Pipeline 서비스 관리",
	Long:  "NHN Cloud Pipeline 서비스를 관리합니다.",
}

var executeCmd = &cobra.Command{
	Use:   "execute [pipeline-name]",
	Short: "파이프라인 수동 실행",
	Args:  cobra.ExactArgs(1),
	RunE:  runExecute,
}

func init() {
	cmd.GetRootCmd().AddCommand(PipelineCmd)
	PipelineCmd.AddCommand(executeCmd)
	PipelineCmd.PersistentFlags().String("app-key", "", "Pipeline AppKey (프로필 설정 오버라이드)")
}

func runExecute(c *cobra.Command, args []string) error {
	pipelineName := args[0]

	appKey, _ := c.Flags().GetString("app-key")
	opts := pipeline.ClientOption{AppKey: appKey}
	pipelineClient, err := pipeline.NewClient(cmd.GetProfile(), cmd.GetDebug(), opts)
	if err != nil {
		return err
	}

	if err := pipelineClient.ExecutePipeline(pipelineName); err != nil {
		return err
	}

	fmt.Printf("파이프라인 '%s' 실행이 요청되었습니다.\n", pipelineName)
	return nil
}
