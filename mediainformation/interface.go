package mediainformation

type mediainformation interface {
	CreateMediaFileHash()

	// Create an empty manifest file
	CreateEmptyManifestFile()

	GenerateSingleCatFileTree(string) // TODO
	GenerateFileTree()
	GenerateNonCatFileTree()

	ReadingManifestFile()

	ReconstructContenSourceFile()
	ReconstructManifestFile()

	SetAbsoluteContentLinkDirPath(string)
	SetAbsoluteLinkDirPath(string)
	SetAbsoluteManifestLinkDirPath(string)
}

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

	// Absolute path of manifest xml file.
	AbsoluteManifestSourcePath string

	// The relative directory path on up with the media files.
	OnsUpPath string

	// The name of the media file
	ContentFileName string

	// The hash value of the media file
	HashValue string

	// The extension of the media file
	Extension string

	// The list with the categories of a media file
	Categories []string
}
