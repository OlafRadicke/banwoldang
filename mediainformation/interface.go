package mediainformation

type mediainformation interface {
	CreateMediaFileHash()
	GenerateSingleCatFileTree(string) // TODO
	GenerateFileTree()
	GenerateNonCatFileTree()
	ReadingManifestFile()
	ReconstructContenSourceFile()
	ReconstructManifestFile()
	SetAbsoluteContentLinkDirPath(string)
	SetAbsoluteLinkDirPath(string)
}

type MediaInformation struct {

	// Relative path of manifest xml file.
	ManifestFilePath string
	// The relative directory path on up with the media files.
	OnsUpPath string
	// The name of the media file
	ContentFileName string
	// The source path of the media as an absolute path
	AbsoluteContentSourcePath string
	// Absolute path of the link directory
	AbsoluteLinkDirPath string
	// Absolute path for the target link in the link directory rom the conten file.
	AbsoluteContentLinkDirPath string
	// The hash value of the media file
	HashValue string
	// The extension of the media file
	Extension string
	// The list with the categories of a media file
	Categories []string
}
