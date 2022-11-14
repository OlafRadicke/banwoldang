package filetree

import (
	"github.com/OlafRadicke/banwoldang/statistics"
	"github.com/OlafRadicke/banwoldang/config"
)

// NewProgArguments create new instance of FileTree and get it back.
func NewFileTree(progConfig *config.YamlConfig, statistics *statistics.Statistics) *FileTree {
	fileTree := FileTree{}

	fileTree.SourcePath = ""
	fileTree.LinkDir = ""
	// fileTree.Findings = 0
	fileTree.UseChecksum = false
	fileTree.UseHardLink = false
	fileTree.progConfig = progConfig
	fileTree.Statistic = statistics
	return &fileTree
}
