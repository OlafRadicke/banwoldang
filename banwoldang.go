package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"

	"github.com/OlafRadicke/banwoldang/config"
	cl "github.com/OlafRadicke/banwoldang/customlogger"
	"github.com/OlafRadicke/banwoldang/filetree"
	"github.com/OlafRadicke/banwoldang/statistics"
	gthumb "github.com/OlafRadicke/go-gthumb"
	"github.com/OlafRadicke/banwoldang/linkdirectories"
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

	WalkThroughCollection(progConfig, statistic, fileTree)
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
		fmt.Println("error! check the log files for more information")
		cl.ErrorLogger.Fatal(err)
	}

	err = yaml.Unmarshal([]byte(configContent), &yamlConf)
	if err != nil {
		fmt.Println("error! check the log files for more information")
		cl.ErrorLogger.Fatal(err)
	}
	return &yamlConf
}

func WalkThroughCollection(progConfig *config.YamlConfig, statistic *statistics.Statistics, fileTree  *filetree.FileTree){
	// Read xml...

	fmt.Println("Start walk in ", progConfig.SourceDir)

	gtFileTree := gthumb.NewFileTree(progConfig.SourceDir)
	gtFileTree.GoThroughCollection()

	for index, path := range gtFileTree.ListOfMediaFiles {
		if progConfig.ShowProgress {
			fmt.Printf("\rMedia files: %d/%d", index, len(gtFileTree.ListOfMediaFiles))
		}
		mediainformation.NewMediaInformation(progConfig, statistic, path)
	}
	fmt.Printf("\n")
	for index, path := range gtFileTree.ListOfCommentFiles {
		if progConfig.ShowProgress {
			fmt.Printf("\rComment files: %d/%d", index, len(gtFileTree.ListOfCommentFiles))
		}
		mediaInfo, err := mediainformation.NewMediaInformationByManifest(progConfig, statistic, path)
		if err != nil {
			cl.ErrorLogger.Println(err)
			cl.ErrorLogger.Println("jump over ", path)
			continue
		}
		linkdirectories.GenerateLinkDirTree(progConfig, statistic, mediaInfo)
		fileTree.JoinAllUsedCategories(mediaInfo.Comments.GetCategories())
	}
	fmt.Printf("\n")

}

