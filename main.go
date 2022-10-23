package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

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

	// statistic := statistics.NewStatistics()
	fileTree := filetree.NewFileTree()
	fileTree.SetAbsoluteSourcePath(progConfig.SourceDir)
	fileTree.SetAbsoluteLinkDir(progConfig.LinkDir)
	fileTree.UseChecksum = progConfig.UseChecksum
	fileTree.UseHardLink = progConfig.UseHardLink
	fileTree.UseFfmpeg = progConfig.UseFfmpeg
	cl.InfoLogger.Println("absoluteLinkDir: ", fileTree.LinkDir)
	cl.InfoLogger.Println("absoluteLinkDir: ", fileTree.LinkDir)
	cl.InfoLogger.Println("Search in: ", fileTree.SourcePath)

	fileTree.GoThroughCollection()
	fileTree.CreateTagsXmlFile()
	writeStatistic(&fileTree.Statistic)
}

func writeStatistic(statistic *statistics.Statistics) {
	f, err := os.Create("./statistics.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	for key, value := range statistic.UsedTags {
		textline := key + ": " + string(value) + "/n"
		_, err = f.WriteString(textline)
		if err != nil {
			cl.ErrorLogger.Println(err)
			return
		}
	}
	textline := "Count of founded files: " + string(statistic.FoundedFiles) + "/n"
	_, err = f.WriteString(textline)
	if err != nil {
		cl.ErrorLogger.Println(err)
		return
	}
	textline = "Count of founded categories: " + string(len(statistic.UsedTags)) + "/n"
	_, err = f.WriteString(textline)
	if err != nil {
		cl.ErrorLogger.Println(err)
		return
	}

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
