package filetree

import (
	"github.com/OlafRadicke/banwoldang/statistics"
	"github.com/OlafRadicke/banwoldang/config"
)

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

	//  A helper object for statistic analyses
	Statistic *statistics.Statistics

	// Application configuration
	progConfig *config.YamlConfig
}

