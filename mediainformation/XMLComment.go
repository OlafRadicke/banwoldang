package mediainformation

import (
	"encoding/xml"
)

// XML strctur of manifest file:
//
// <?xml version="1.0" encoding="UTF-8"?>
// <comment version="3.0">
//   <caption/>
//   <note/>
//   <place/>
//   <categories>
//     <category value="c-anal"/>
//     <category value="f-athletic"/>
//     <category value="f-brunette"/>
//     <category value="l-living_room"/>
//     <category value="n-german"/>
//     <category value="q-90s"/>
//     <category value="q-aaa"/>
//     <category value="q-aaaa"/>
//     <category value="q-self_made"/>
//     <category value="q-video_hd"/>
//     <category value="s-bottomless"/>
//     <category value="s-chav"/>
//     <category value="s-socks"/>
//     <category value="s-tattoo"/>
//   </categories>
// </comment>

type XMLComment struct {
	XMLName    xml.Name      `xml:"comment"`
	Caption    string        `xml:"caption"`
	Note       string        `xml:"note"`
	Place      string        `xml:"place"`
	Categories XMLCategories `xml:"categories"`
}

type XMLCategories struct {
	CategoryList []XMLCategory `xml:"category"`
	Value        string        `xml:"value,attr"`
}

type XMLCategory struct {
	Value string `xml:"value,attr"`
}
