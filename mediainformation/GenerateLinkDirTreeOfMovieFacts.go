package mediainformation

import (
	"context"
	"os"
	"strconv"
	"time"

	cl "github.com/OlafRadicke/banwoldang/customlogger"
	"gopkg.in/vansante/go-ffprobe.v2"
)

// GenerateLinkDirTreeOfMovieFacts Creates a directory tree with links based on the
// the facts of the found at movie files:
// - Links for duration in minutes
func (mediaInfo *MediaInformation) GenerateLinkDirTreeOfMovieFacts() {
	var minutes_duration string
	var resolution_height string
	var resolution_width string

	cl.InfoLogger.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	cl.InfoLogger.Println("++++++++++++++++++++++ create movie facts category +++++++++++++++++++")
	cl.InfoLogger.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")

	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	reader, err := os.Open(mediaInfo.AbsoluteContentSourcePath)
	if err != nil {
		cl.InfoLogger.Println(err)
		cl.InfoLogger.Println("File: ", mediaInfo.AbsoluteContentSourcePath)
		cl.InfoLogger.Println("no media file")
		return
	}
	defer reader.Close()

	data, err := ffprobe.ProbeReader(ctx, reader)
	if err != nil {
		cl.InfoLogger.Println("Error getting data: ", err)
		cl.InfoLogger.Println("File: ", mediaInfo.AbsoluteContentSourcePath)
		cl.InfoLogger.Println("no media file")
		return
	}
	if len(data.Streams) == 0 {
		cl.InfoLogger.Println("File has no streams: ", mediaInfo.AbsoluteContentSourcePath)
		return
	}
	cl.InfoLogger.Println("CodecType ", data.Streams[0].CodecType)
	cl.InfoLogger.Println("Height ", data.Streams[0].Height)
	resolution_height = strconv.Itoa(data.Streams[0].Height)
	cl.InfoLogger.Println("Width ", data.Streams[0].Width)
	resolution_width = strconv.Itoa(data.Streams[0].Width)
	seconds := data.Streams[0].Duration
	int_seconds, err := strconv.ParseFloat(seconds, 64)
	if err != nil {
		cl.InfoLogger.Println(err)
		cl.InfoLogger.Println("Duration: ", data.Streams[0].Duration)
		cl.InfoLogger.Println("File: ", mediaInfo.AbsoluteContentSourcePath)
		return
	}
	minutes := int(int_seconds / 60)
	cl.InfoLogger.Println("minutes ", minutes)
	if minutes == 0 {
		cl.ErrorLogger.Println("duration is 0 minutes")
		cl.ErrorLogger.Println("File: ", mediaInfo.AbsoluteContentSourcePath)
		return
	}
	minutes_duration = strconv.Itoa(minutes)

	// Links for duration in minutes
	mediaInfo.SetAbsoluteContentLinkDirPath("duration/" + minutes_duration)
	mediaInfo.SetAbsoluteManifestLinkDirPath("duration/" + minutes_duration)
	mediaInfo.CreateLinkDirSubDir("duration/" + minutes_duration)
	mediaInfo.CreateContentLink()
	// mediaInfo.CreateNewEmptyManifestFile()
	mediaInfo.CreateManifestLink()
	mediaInfo.Statistics.ContDuration(minutes)

	// Links for height
	mediaInfo.SetAbsoluteContentLinkDirPath("resolution_height/" + resolution_height)
	mediaInfo.SetAbsoluteManifestLinkDirPath("resolution_height/" + resolution_height)
	mediaInfo.CreateLinkDirSubDir("resolution_height/" + resolution_height)
	mediaInfo.CreateContentLink()
	// mediaInfo.CreateNewEmptyManifestFile()
	mediaInfo.CreateManifestLink()
	mediaInfo.Statistics.ContResolutionHeight(data.Streams[0].Height)

	// Links for width
	mediaInfo.SetAbsoluteContentLinkDirPath("resolution_width/" + resolution_width)
	mediaInfo.SetAbsoluteManifestLinkDirPath("resolution_width/" + resolution_width)
	mediaInfo.CreateLinkDirSubDir("resolution_width/" + resolution_width)
	mediaInfo.CreateContentLink()
	// mediaInfo.CreateNewEmptyManifestFile()
	mediaInfo.CreateManifestLink()
	mediaInfo.Statistics.ContResolutionWidth(data.Streams[0].Width)
}
