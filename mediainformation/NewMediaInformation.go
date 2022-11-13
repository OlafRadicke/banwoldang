package mediainformation

import (
	"github.com/OlafRadicke/banwoldang/statistics"
	"github.com/OlafRadicke/banwoldang/config"
	cl "github.com/OlafRadicke/banwoldang/customlogger"
	gt "github.com/OlafRadicke/go-gthumb"
)



// NewProgArguments create new instance of MediaInformation and get it back.
// @path Path of media file.
func NewMediaInformation(progConfig *config.YamlConfig, statistics *statistics.Statistics, path string) *MediaInformation {
	var err error
	mediaInfo := MediaInformation{}
	mediaInfo.progConfig = progConfig
	mediaInfo.Statistics = statistics
	mediaInfo.Comments, err = gt.NewCommentsFile(path)
	if err != nil {
		cl.InfoLogger.Println("error to init object: ", err)
	}
	mediaInfo.SetAbsoluteLinkDirPath(progConfig.LinkDir)
	mediaInfo.SetAbsoluteContentSourcePath(path)
	mediaInfo.ReconstructManifestFile()
	mediaInfo.CreateMediaFileHash()
	return &mediaInfo
}

