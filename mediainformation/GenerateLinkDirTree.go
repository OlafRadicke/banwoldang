package mediainformation

// Generate directory tree with symlinks of file with categories.
func (mediaInfo *MediaInformation) GenerateLinkDirTree() {

	mediaInfo.GenerateLinkDirTreeOfOldNameParts()
	mediaInfo.GenerateLinkDirTreeOfCategories()
	mediaInfo.GenerateLinkDirTreeOfChecksum()
	if mediaInfo.UseFfmpeg {
		mediaInfo.GenerateLinkDirTreeOfMovieFacts()
	}
}
