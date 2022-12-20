package commit

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
)

var Revert = NewIntegrationTest(NewIntegrationTestArgs{
	Description:  "Reverts a commit",
	ExtraCmdArgs: "",
	Skip:         false,
	SetupConfig:  func(config *config.AppConfig) {},
	SetupRepo: func(shell *Shell) {
		shell.CreateFile("myfile", "myfile content")
		shell.GitAddAll()
		shell.Commit("first commit")
	},
	Run: func(shell *Shell, input *Input, assert *Assert, keys config.KeybindingConfig) {
		assert.CommitCount(1)

		input.SwitchToCommitsWindow()

		input.PressKeys(keys.Commits.RevertCommit)
		assert.InConfirm()
		assert.MatchCurrentViewTitle(Equals("Revert commit"))
		assert.MatchCurrentViewContent(MatchesRegexp("Are you sure you want to revert \\w+?"))
		input.Confirm()

		assert.CommitCount(2)
		assert.MatchHeadCommitMessage(Contains("Revert \"first commit\""))
		input.PreviousItem()
		assert.MatchMainViewContent(Contains("-myfile content"))
		assert.FileSystemPathNotPresent("myfile")

		input.Wait(10)
	},
})