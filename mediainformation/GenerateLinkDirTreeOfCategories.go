package mediainformation

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"gopkg.in/vansante/go-ffprobe.v2"

	cl "github.com/OlafRadicke/banwoldang/customlogger"
)

// GenerateLinkDirTreeOfCategories Creates a directory tree with links based on the
// categories in the xml manifest files.
func (mediaInfo *MediaInformation) GenerateLinkDirTreeOfCategories() {

	if len(mediaInfo.Categories) > 0 {
		cl.InfoLogger.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
		cl.InfoLogger.Println("+++++++++++++++++++ create links for ", len(mediaInfo.Categories), " categories ++++++++++++++++++++")
		cl.InfoLogger.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	}

	for i := 0; i < len(mediaInfo.Categories); i++ {
		if len(mediaInfo.Categories) < 6 {
			mediaInfo.GenerateLinkDirTreeOfCategoryCount()
		}

		mediaInfo.SetAbsoluteContentLinkDirPath("categories/" + mediaInfo.Categories[i])
		mediaInfo.SetAbsoluteManifestLinkDirPath("categories/" + mediaInfo.Categories[i])
		mediaInfo.CreateLinkDirSubDir("categories/" + mediaInfo.Categories[i])
		mediaInfo.CreateContentLink()
		mediaInfo.CreateNewEmptyManifestFile()
		mediaInfo.CreateManifestLink()

	}
	mediFacs(mediaInfo.AbsoluteContentSourcePath)
}

func mediFacs(mediaPath string) {

	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	// fmt.Println("Try to open ", mediaPath)
	reader, err := os.Open(mediaPath)
	if err != nil {
		fmt.Println(err)
	}
	defer reader.Close()

	data, err := ffprobe.ProbeReader(ctx, reader)
	if err != nil {
		fmt.Println("Error getting data: ", err)
		fmt.Println("no media file")
		return
	}

	fmt.Println("Height ", data.Streams[0].Height)
	fmt.Println("Width ", data.Streams[0].Width)
	// fmt.Println("Duration ", data.Streams[0].Duration)
	seconds := data.Streams[0].Duration
	int_seconds, err := strconv.ParseFloat(seconds, 64)
	if err != nil {
		// ... handle error
		fmt.Println(err)
	}
	minutes := int(int_seconds / 60)
	fmt.Println("minutes ", minutes)
	fmt.Println("-----------------------------------------")

}
