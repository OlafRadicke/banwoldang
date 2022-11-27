package mediainformation

import (
	"github.com/OlafRadicke/banwoldang/statistics"
	"github.com/OlafRadicke/banwoldang/config"
	cl "github.com/OlafRadicke/banwoldang/customlogger"
	gt "github.com/OlafRadicke/go-gthumb"
)



// NewProgArguments create new instance of MediaInformation and get it back.
// @path Path of comment file.
func NewMediaInformationByManifest(progConfig *config.YamlConfig, statistics *statistics.Statistics, path string) (*MediaInformation, error) {
	var err error
	mediaInfo := MediaInformation{}
	mediaInfo.progConfig = progConfig
	mediaInfo.Statistics = statistics
	mediaInfo.Comments, err = gt.NewCommentsFile(path)
	if err != nil {
		cl.InfoLogger.Println("error to init object: ", err)
	}
	mediaInfo.hashValue = ""
	mediaInfo.SetAbsoluteLinkDirPath(progConfig.LinkDir)
	mediaInfo.SetAbsoluteManifestSourcePath(path)
	err = mediaInfo.ReconstructContenSourceFile()
	if err != nil {
		cl.ErrorLogger.Println("ReconstructContenSourceFile error: ", err)
		return nil, err
	}
	_, err = mediaInfo.GetHashValue()
	if err != nil {
		cl.ErrorLogger.Println("Hash sum error: ", err)
		return nil, err
	}

	// mediaInfo.CreateMediaFileHash()
	return &mediaInfo, nil
}

