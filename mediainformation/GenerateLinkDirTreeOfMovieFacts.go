package mediainformation

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"gopkg.in/vansante/go-ffprobe.v2"
)

// GenerateLinkDirTreeOfMovieFacts Creates a directory tree with links based on the
// the facts of the found at movie files:
// - Links for duration in minutes
func (mediaInfo *MediaInformation) GenerateLinkDirTreeOfMovieFacts() {
	var minutes_duration string
	var resolution_height string
	var resolution_width string

	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	reader, err := os.Open(mediaInfo.AbsoluteContentSourcePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer reader.Close()

	data, err := ffprobe.ProbeReader(ctx, reader)
	if err != nil {
		fmt.Println("Error getting data: ", err)
		fmt.Println("no media file")
		return
	}
	if len(data.Streams) == 0 {
		return
	}
	fmt.Println("CodecType ", data.Streams[0].CodecType)
	fmt.Println("Height ", data.Streams[0].Height)
	resolution_height = strconv.Itoa(data.Streams[0].Height)
	fmt.Println("Width ", data.Streams[0].Width)
	resolution_width = strconv.Itoa(data.Streams[0].Width)
	// fmt.Println("Duration ", data.Streams[0].Duration)
	seconds := data.Streams[0].Duration
	int_seconds, err := strconv.ParseFloat(seconds, 64)
	if err != nil {
		// ... handle error
		fmt.Println(err)
		return
	}
	minutes := int(int_seconds / 60)
	fmt.Println("minutes ", minutes)
	if minutes == 0 {
		return
	}
	minutes_duration = strconv.Itoa(minutes)
	fmt.Println("-----------------------------------------")

	// Links for duration in minutes
	mediaInfo.SetAbsoluteContentLinkDirPath("duration/" + minutes_duration)
	mediaInfo.SetAbsoluteManifestLinkDirPath("duration/" + minutes_duration)
	mediaInfo.CreateLinkDirSubDir("duration/" + minutes_duration)
	mediaInfo.CreateContentLink()
	mediaInfo.CreateNewEmptyManifestFile()
	mediaInfo.CreateManifestLink()

	// Links for height
	mediaInfo.SetAbsoluteContentLinkDirPath("resolution_height/" + resolution_height)
	mediaInfo.SetAbsoluteManifestLinkDirPath("resolution_height/" + resolution_height)
	mediaInfo.CreateLinkDirSubDir("resolution_height/" + resolution_height)
	mediaInfo.CreateContentLink()
	mediaInfo.CreateNewEmptyManifestFile()
	mediaInfo.CreateManifestLink()

	// Links for height
	mediaInfo.SetAbsoluteContentLinkDirPath("resolution_width/" + resolution_width)
	mediaInfo.SetAbsoluteManifestLinkDirPath("resolution_width/" + resolution_width)
	mediaInfo.CreateLinkDirSubDir("resolution_width/" + resolution_width)
	mediaInfo.CreateContentLink()
	mediaInfo.CreateNewEmptyManifestFile()
	mediaInfo.CreateManifestLink()

}
