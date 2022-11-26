package mediainformation

import (
	"github.com/OlafRadicke/banwoldang/statistics"
	"github.com/OlafRadicke/banwoldang/config"
	gt "github.com/OlafRadicke/go-gthumb"
)


// MediaInformation Represent information about a single Media with his
// Manifest file and data
type MediaInformation struct {

	// Absolute path for the target link in the link directory from the conten
	// file.
	AbsoluteContentLinkDirPath string

	// The source path of the media as an absolute path
	AbsoluteContentSourcePath string

	// Absolute path of the link directory
	AbsoluteLinkDirPath string

	// Absolute path for the target link in the link directory from the
	// manifest file.
	AbsoluteManifestLinkDirPath string

	// linkSubDir The sub directory of the link directory.
	linkSubDir string

	// The relative directory path on up with the media files.
	OnsUpPath string

	// The name of the media file
	ContentFileName string

	// The hash value of the media file
	hashValue string

	// The extension of the media file
	Extension string

	// The list with the categories of a media file
	// Categories []string

	// This object represents the conten of the XML comment files
	Comments *gt.CommentsFile

	// A helper object for statistic analyses
	Statistics *statistics.Statistics

	// Application configuration
	progConfig *config.YamlConfig
}
