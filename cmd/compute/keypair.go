package compute

import (
	"fmt"

	"nhncli/internal/compute"
	"nhncli/internal/output"

	"github.com/spf13/cobra"
)

var (
	keypairName      string
	keypairPublicKey string
)

var keypairCmd = &cobra.Command{
	Use:   "keypair",
	Short: "키페어 관리",
}

var keypairListCmd = &cobra.Command{
	Use:   "list",
	Short: "키페어 목록 조회",
	RunE:  runKeypairList,
}

var keypairCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "키페어 생성",
	Long: `새로운 키페어를 생성합니다.

예시:
  # 새 키 생성 (개인키 출력)
  nhn compute keypair create --name my-keypair

  # 기존 공개키 등록
  nhn compute keypair create --name my-keypair --public-key "ssh-rsa AAAA..."`,
	RunE: runKeypairCreate,
}

var keypairDeleteCmd = &cobra.Command{
	Use:   "delete <keypair-name>",
	Short: "키페어 삭제",
	Args:  cobra.ExactArgs(1),
	RunE:  runKeypairDelete,
}

func init() {
	ComputeCmd.AddCommand(keypairCmd)
	keypairCmd.AddCommand(keypairListCmd)
	keypairCmd.AddCommand(keypairCreateCmd)
	keypairCmd.AddCommand(keypairDeleteCmd)

	keypairCreateCmd.Flags().StringVar(&keypairName, "name", "", "키페어 이름 (필수)")
	keypairCreateCmd.Flags().StringVar(&keypairPublicKey, "public-key", "", "공개키 (선택, 없으면 새로 생성)")
	keypairCreateCmd.MarkFlagRequired("name")
}

func runKeypairList(cmd *cobra.Command, args []string) error {
	client, err := compute.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	keypairs, err := client.ListKeypairs()
	if err != nil {
		return err
	}

	if GetOutput() == "json" {
		return output.PrintJSON(keypairs)
	}

	headers := []string{"NAME", "FINGERPRINT"}
	rows := make([][]string, len(keypairs))
	for i, kp := range keypairs {
		rows[i] = []string{kp.Name, kp.Fingerprint}
	}

	return output.Print(GetOutput(), &output.TableData{
		Headers: headers,
		Rows:    rows,
	}, keypairs)
}

func runKeypairCreate(cmd *cobra.Command, args []string) error {
	client, err := compute.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	kp, err := client.CreateKeypair(keypairName, keypairPublicKey)
	if err != nil {
		return err
	}

	if GetOutput() == "json" {
		return output.PrintJSON(kp)
	}

	fmt.Printf("✅ 키페어가 생성되었습니다.\n")
	fmt.Printf("Name:        %s\n", kp.Name)
	fmt.Printf("Fingerprint: %s\n", kp.Fingerprint)

	if kp.PrivateKey != "" {
		fmt.Printf("\n⚠️  아래 개인키를 안전하게 저장하세요. 다시 확인할 수 없습니다.\n")
		fmt.Printf("\n%s\n", kp.PrivateKey)
	}

	return nil
}

func runKeypairDelete(cmd *cobra.Command, args []string) error {
	name := args[0]

	client, err := compute.NewClient(GetProfile(), GetRegion(), GetDebug())
	if err != nil {
		return err
	}

	if err := client.DeleteKeypair(name); err != nil {
		return err
	}

	fmt.Printf("✅ 키페어 %s가 삭제되었습니다.\n", name)
	return nil
}
