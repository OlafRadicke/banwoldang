package linkdirectories

import (
	"github.com/OlafRadicke/banwoldang/config"
	"github.com/OlafRadicke/banwoldang/mediainformation"
)

// Generate directory tree with symlinks of file with categories.
func GenerateLinkDirTree(progConfig *config.YamlConfig, mediaInfo *mediainformation.MediaInformation) {

	GenerateLinkDirTreeOfCategories(mediaInfo)
	GenerateLinkDirTreeOfCategoryCount(mediaInfo)
	GenerateLinkDirTreeOfChecksum(mediaInfo)
	if progConfig.UseFfmpeg {
		GenerateLinkDirTreeOfMovieFacts(mediaInfo)
	}
	GenerateLinkDirTreeOfOldNameParts(mediaInfo)
	GenerateLinkDirTreeWithoutManifests(mediaInfo)
}
