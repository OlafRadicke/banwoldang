package mediainformation

// Generate directory tree with symlinks of file with categories.
func (mediaInfo *MediaInformation) GenerateLinkDirTree() {

	mediaInfo.GenerateLinkDirTreeOfOldNameParts()
	if mediaInfo.progConfig.UseFfmpeg {
		mediaInfo.GenerateLinkDirTreeOfMovieFacts()
	}
}
