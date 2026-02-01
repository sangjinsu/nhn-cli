package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish|powershell]",
	Short: "셸 자동완성 스크립트 생성",
	Long: `셸 자동완성 스크립트를 생성합니다.

지원 셸: bash, zsh, fish, powershell

설치 방법은 각 서브커맨드의 도움말을 참조하세요:
  nhn completion bash --help
  nhn completion zsh --help
  nhn completion fish --help
  nhn completion powershell --help`,
}

var completionBashCmd = &cobra.Command{
	Use:   "bash",
	Short: "bash 자동완성 스크립트 생성",
	Long: `bash 셸 자동완성 스크립트를 생성합니다.

현재 셸에 적용:
  source <(nhn completion bash)

영구 적용 (Linux):
  nhn completion bash > /etc/bash_completion.d/nhn

영구 적용 (macOS, bash-completion 필요):
  nhn completion bash > $(brew --prefix)/etc/bash_completion.d/nhn`,
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Root().GenBashCompletionV2(os.Stdout, true)
	},
}

var completionZshCmd = &cobra.Command{
	Use:   "zsh",
	Short: "zsh 자동완성 스크립트 생성",
	Long: `zsh 셸 자동완성 스크립트를 생성합니다.

현재 셸에 적용:
  source <(nhn completion zsh)

영구 적용:
  nhn completion zsh > "${fpath[1]}/_nhn"

참고: zsh 자동완성을 처음 사용하는 경우 다음을 ~/.zshrc에 추가하세요:
  autoload -Uz compinit && compinit`,
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Root().GenZshCompletion(os.Stdout)
	},
}

var completionFishCmd = &cobra.Command{
	Use:   "fish",
	Short: "fish 자동완성 스크립트 생성",
	Long: `fish 셸 자동완성 스크립트를 생성합니다.

현재 셸에 적용:
  nhn completion fish | source

영구 적용:
  nhn completion fish > ~/.config/fish/completions/nhn.fish`,
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Root().GenFishCompletion(os.Stdout, true)
	},
}

var completionPowershellCmd = &cobra.Command{
	Use:   "powershell",
	Short: "PowerShell 자동완성 스크립트 생성",
	Long: `PowerShell 자동완성 스크립트를 생성합니다.

현재 셸에 적용:
  nhn completion powershell | Out-String | Invoke-Expression

영구 적용:
  nhn completion powershell > nhn.ps1
  # 위 파일을 $PROFILE에서 소싱`,
	DisableFlagsInUseLine: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
	completionCmd.AddCommand(completionBashCmd)
	completionCmd.AddCommand(completionZshCmd)
	completionCmd.AddCommand(completionFishCmd)
	completionCmd.AddCommand(completionPowershellCmd)
}
