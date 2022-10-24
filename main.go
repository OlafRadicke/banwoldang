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

	// cl.InfoLogger.Println("PROGRAMM START")
	// cl.ErrorLogger.Println("PROGRAMM START", 1, "test")
	// cl.ErrorLogger.Fatal("der neue Error-Logger mit hartem Ende...")

	progConfig := checkArguments()
	fmt.Println("SourceDir: ", progConfig.SourceDir)

	statistic := statistics.NewStatistics()
	fileTree := filetree.NewFileTree(statistic)
	fileTree.SetAbsoluteSourcePath(progConfig.SourceDir)
	fileTree.SetAbsoluteLinkDir(progConfig.LinkDir)
	fileTree.UseChecksum = progConfig.UseChecksum
	fileTree.UseHardLink = progConfig.UseHardLink
	fileTree.UseFfmpeg = progConfig.UseFfmpeg

	cl.InfoLogger.Println("absoluteLinkDir: ", fileTree.LinkDir)
	cl.InfoLogger.Println("Search in: ", fileTree.SourcePath)
	fileTree.GoThroughCollection()
	fileTree.CreateTagsXmlFile()
	fileTree.Statistic.WriteStatistic(progConfig.LinkDir)
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
