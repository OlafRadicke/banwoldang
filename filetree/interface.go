package filetree

import (
	"os"

	"github.com/OlafRadicke/banwoldang/statistics"
)

// filetree The inteface of the feletree struckt
type filetree interface {
	GoThroughCollection()
	NewFileTree() FileTree
	PrintAll()

	fileHandler(path string, info os.FileInfo, err error) error

	SetAbsoluteSourcePath(string)
	SetAbsoluteLinkDir(string)
	JoinAllUsedCategories(Categories []string)
	CreateTagsXmlFile()
}

// FileTree This struct represent the information about the tree
type FileTree struct {
	// SourcePath The start path for searching media files
	SourcePath string

	// LinkDir Location for the link directory
	LinkDir string

	// UseChecksum Is the value true, than it will be create checksums as file names (for the links)
	UseChecksum bool

	// UseHardLink Is the value true, than it will be try to use hard links (for the link tree directory)
	UseHardLink bool

	// Is the value true, than ffmpeg support try to create links about media facts
	UseFfmpeg bool

	//  A helper obect for statistic analyses
	Statistic statistics.Statistics
}

// NewProgArguments create new instance of FileTree and get it back.
func NewFileTree() FileTree {
	fileTree := FileTree{}

	fileTree.SourcePath = ""
	fileTree.LinkDir = ""
	// fileTree.Findings = 0
	fileTree.UseChecksum = false
	fileTree.UseHardLink = false
	fileTree.Statistic = statistics.NewStatistics()
	return fileTree
}
