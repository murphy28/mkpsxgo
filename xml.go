package mkpsxgo

import (
	"encoding/xml"
)

// IsoProject represents the root <iso_project> element.
type IsoProject struct {
	XMLName   xml.Name `xml:"iso_project"`
	ImageName string   `xml:"image_name,attr"`
	CueSheet  string   `xml:"cue_sheet,attr"`
	NoXa      int      `xml:"no_xa,attr"`
	Tracks    []Track  `xml:"track"`
}

// Track represents a <track> element.
type Track struct {
	XMLName           xml.Name           `xml:"track"`
	Type              string             `xml:"type,attr"`
	XaEdc             *bool              `xml:"xa_edc,attr"`
	NewType           *bool              `xml:"new_type,attr"`
	Source            string             `xml:"source,attr"`
	TrackID           string             `xml:"trackid,attr"`
	Identifiers       *Identifiers       `xml:"identifiers"`
	License           *License           `xml:"license"`
	DefaultAttributes *DefaultAttributes `xml:"default_attributes"`
	DirectoryTree     *DirectoryTree     `xml:"directory_tree"`
	Pregap            *Pregap            `xml:"pregap"`
}

// Identifiers represents the <identifiers> element.
type Identifiers struct {
	XMLName      xml.Name `xml:"identifiers"`
	System       string   `xml:"system,attr"`
	Application  string   `xml:"application,attr"`
	Volume       string   `xml:"volume,attr"`
	VolumeSet    string   `xml:"volume_set,attr"`
	Publisher    string   `xml:"publisher,attr"`
	DataPreparer string   `xml:"data_preparer,attr"`
	Copyright    string   `xml:"copyright,attr"`
	CreationDate string   `xml:"creation_date,attr"`
}

// License represents the <license> element.
type License struct {
	XMLName xml.Name `xml:"license"`
	File    string   `xml:"file,attr"`
}

// DefaultAttributes represents the <default_attributes> element.
type DefaultAttributes struct {
	XMLName  xml.Name `xml:"default_attributes"`
	GmtOffs  int      `xml:"gmt_offs,attr"`
	XaAttrib int      `xml:"xa_attrib,attr"`
	XaPerm   int      `xml:"xa_perm,attr"`
	XaGid    int      `xml:"xa_gid,attr"`
	XaUid    int      `xml:"xa_uid,attr"`
}

// DirectoryTree represents the <directory_tree> element.
// It can contain both <dir> and <file> elements directly.
type DirectoryTree struct {
	XMLName xml.Name `xml:"directory_tree"`
	SrcDir  string   `xml:"srcdir,attr"`
	Dirs    []Dir    `xml:"dir"`
	Files   []File   `xml:"file"`
	Dummy   *Dummy   `xml:"dummy"`
}

// Dir represents a <dir> element within a directory_tree.
type Dir struct {
	XMLName xml.Name `xml:"dir"`
	Name    string   `xml:"name,attr"`
	Source  string   `xml:"source,attr"`
	SrcDir  string   `xml:"srcdir,attr"`
	Files   []File   `xml:"file"`
	Dirs    []Dir    `xml:"dir"`
}

// File represents a <file> element within a directory_tree or a dir.
type File struct {
	XMLName xml.Name `xml:"file"`
	Name    string   `xml:"name,attr"`
	Source  string   `xml:"source,attr"`
	Type    string   `xml:"type,attr"`
	TrackID string   `xml:"trackid,attr"`
}

// Dummy represents a <dummy> element.
type Dummy struct {
	XMLName xml.Name `xml:"dummy"`
	Sectors int      `xml:"sectors,attr"`
	Type    int      `xml:"type,attr"`
}

// Pregap represents the <pregap> element for audio tracks.
type Pregap struct {
	XMLName  xml.Name `xml:"pregap"`
	Duration string   `xml:"duration,attr"`
}
