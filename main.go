package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"

	"github.com/OlafRadicke/banwoldang/config"
	cl "github.com/OlafRadicke/banwoldang/customlogger"
	filetree "github.com/OlafRadicke/banwoldang/filetree"
	"github.com/OlafRadicke/banwoldang/statistics"
	gt "github.com/OlafRadicke/go-gthumb"
	ld "github.com/OlafRadicke/banwoldang/linkdirectories"
	"github.com/OlafRadicke/banwoldang/mediainformation"

)

// ProgArguments struct with command line arguments
type ProgArguments struct {
	// Location with the media files
	SourceDir string
	// Location for the link directory
	LinkDir string
	// The minimum of needed params
	MinimumArguments int
	// Is the value true, than it will be create checksums as file names (for the links)
	UseChecksum bool
	// Is the value true, than it will be try to create hard links.
	UseHardLink bool
}

func main() {

	cl.InfoLogger.Println("================================= PROGRAMM START ==============================")
	cl.ErrorLogger.Println("================================ PROGRAMM START ==============================")

	progConfig := checkArguments()
	fmt.Println("SourceDir: ", progConfig.SourceDir)

	statistic := statistics.NewStatistics(progConfig.LinkDir)
	fileTree := filetree.NewFileTree(progConfig, statistic)
	fileTree.SetAbsoluteSourcePath(progConfig.SourceDir)
	fileTree.SetAbsoluteLinkDir(progConfig.LinkDir)

	cl.InfoLogger.Println("absoluteLinkDir: ", fileTree.LinkDir)
	cl.InfoLogger.Println("Search in: ", fileTree.SourcePath)

	useNewLib(progConfig, statistic, fileTree)
	// fileTree.GoThroughCollection()
	fileTree.CreateTagsXmlFile()
	fileTree.Statistic.WriteStatistic()
}

func checkArguments() *config.YamlConfig {
	configPath := flag.String("f", "banwoldang.yaml", "configuration file")
	flag.Parse()

	yamlConf := readConfig(*configPath)
	return yamlConf
}

func readConfig(configPath string) *config.YamlConfig {
	var yamlConf config.YamlConfig

	configContent, err := ioutil.ReadFile(configPath)
	if err != nil {
		cl.ErrorLogger.Fatal(err)
	}

	err = yaml.Unmarshal([]byte(configContent), &yamlConf)
	if err != nil {
		cl.ErrorLogger.Fatal(err)
	}
	return &yamlConf
}

func useNewLib(progConfig *config.YamlConfig, statistic *statistics.Statistics, fileTree  *filetree.FileTree){
	// Read xml...

	fmt.Println("Start walk in ", progConfig.SourceDir)

	gtFileTree := gt.NewFileTree(progConfig.SourceDir)
	gtFileTree.GoThroughCollection()


	for index, path := range gtFileTree.ListOfMediaFiles {
		if progConfig.ShowProgress {
			fmt.Printf("\rMedia files: %d/%d", index, len(gtFileTree.ListOfMediaFiles))
		}
		mediaInfo := mediainformation.NewMediaInformation(progConfig, statistic, path)
		ld.GenerateLinkDirTreeOfChecksum(mediaInfo)
	}
	fmt.Printf("\n")
	for index, path := range gtFileTree.ListOfCommentFiles {
		if progConfig.ShowProgress {
			fmt.Printf("\rComment files: %d/%d", index, len(gtFileTree.ListOfCommentFiles))
		}
		mediaInfo := mediainformation.NewMediaInformationByManifest(progConfig, statistic, path)
		ld.GenerateLinkDirTreeOfCategories(mediaInfo)

		fileTree.JoinAllUsedCategories(mediaInfo.Comments.GetCategories())
		// mediaInfo.GenerateLinkDirTreeWithoutManifests()
	}
	fmt.Printf("\n")

}

