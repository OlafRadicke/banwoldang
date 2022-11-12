package config

type YamlConfig struct {
	// Location with the media files
	SourceDir string `yaml:"source_dir"`
	// Location for the link directory
	LinkDir string `yaml:"link_dir"`
	// Is the value true, than it will be create checksums as file names (for the links)
	UseChecksum bool `yaml:"real_checksum"`
	// Is the value true, than it will be try to create hard links.
	UseHardLink bool `yaml:"hard_links"`
	// Is the value true, than ffmpeg support try to create links about media facts
	UseFfmpeg bool `yaml:"ffmpeg_support"`
	// Is the value true, than it will show the working progress
	ShowProgress bool `yaml:"show_progress"`
}
