
package linkdirectories

import (
	"github.com/OlafRadicke/banwoldang/mediainformation"
	gt "github.com/OlafRadicke/go-gthumb"
)


type Linkdirectories struct {
	mediaInfo *mediainformation.MediaInformation
}

func NewLinkdirectories(mediaInfo *mediainformation.MediaInformation) *Linkdirectories {
	linkdirectories := Linkdirectories{}
	linkdirectories.mediaInfo = mediaInfo
	return &linkdirectories
}


func (linkdirectories *Linkdirectories) getCategories() ([]gt.XmlCategory){
	return linkdirectories.mediaInfo.Comments.GetCategories()
}

