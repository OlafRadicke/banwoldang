package linkdirectories

import (
	"github.com/OlafRadicke/banwoldang/config"
	"github.com/OlafRadicke/banwoldang/mediainformation"
	"github.com/OlafRadicke/banwoldang/statistics"
	cl "github.com/OlafRadicke/banwoldang/customlogger"
)

// Generate directory tree with symlinks of file with categories.
func GenerateLinkDirTree(progConfig *config.YamlConfig, statistics *statistics.Statistics, mediaInfo *mediainformation.MediaInformation) {
	var err error
	GenerateLinkDirTreeOfCategoryCount(mediaInfo)
	err = GenerateLinkDirTreeOfChecksum(mediaInfo, statistics)
	if err != nil {
		cl.ErrorLogger.Println(err)
		return
	}
	GenerateLinkDirTreeOfCategories(mediaInfo)
	if progConfig.UseFfmpeg {
		GenerateLinkDirTreeOfMovieFacts(mediaInfo)
	}
	GenerateLinkDirTreeOfOldNameParts(mediaInfo)
	GenerateLinkDirTreeWithoutManifests(mediaInfo)
}
