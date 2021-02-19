package main

import "encoding/xml"

// Metadata main struct
type Metadata struct {
	XMLName  xml.Name `xml:"metadata"`
	Text     string   `xml:",chardata"`
	Xmlns    string   `xml:"xmlns,attr"`
	Rpm      string   `xml:"rpm,attr"`
	Packages string   `xml:"packages,attr"`
	Package  []struct {
		Text    string `xml:",chardata"`
		Type    string `xml:"type,attr"`
		Name    string `xml:"name"`
		Arch    string `xml:"arch"`
		Version struct {
			Text  string `xml:",chardata"`
			Epoch string `xml:"epoch,attr"`
			Ver   string `xml:"ver,attr"`
			Rel   string `xml:"rel,attr"`
		} `xml:"version"`
		Checksum struct {
			Text  string `xml:",chardata"`
			Type  string `xml:"type,attr"`
			Pkgid string `xml:"pkgid,attr"`
		} `xml:"checksum"`
		Summary     string `xml:"summary"`
		Description string `xml:"description"`
		Packager    string `xml:"packager"`
		URL         string `xml:"url"`
		Time        struct {
			Text  string `xml:",chardata"`
			File  string `xml:"file,attr"`
			Build string `xml:"build,attr"`
		} `xml:"time"`
		Size struct {
			Text      string `xml:",chardata"`
			Package   string `xml:"package,attr"`
			Installed string `xml:"installed,attr"`
			Archive   string `xml:"archive,attr"`
		} `xml:"size"`
		Location struct {
			Text string `xml:",chardata"`
			Href string `xml:"href,attr"`
		} `xml:"location"`
		Format struct {
			Text        string `xml:",chardata"`
			License     string `xml:"license"`
			Vendor      string `xml:"vendor"`
			Group       string `xml:"group"`
			Buildhost   string `xml:"buildhost"`
			Sourcerpm   string `xml:"sourcerpm"`
			HeaderRange struct {
				Text  string `xml:",chardata"`
				Start string `xml:"start,attr"`
				End   string `xml:"end,attr"`
			} `xml:"header-range"`
			Provides struct {
				Text  string `xml:",chardata"`
				Entry []struct {
					Text  string `xml:",chardata"`
					Name  string `xml:"name,attr"`
					Flags string `xml:"flags,attr"`
					Epoch string `xml:"epoch,attr"`
					Ver   string `xml:"ver,attr"`
					Rel   string `xml:"rel,attr"`
				} `xml:"entry"`
			} `xml:"provides"`
			Requires struct {
				Text  string `xml:",chardata"`
				Entry struct {
					Text  string `xml:",chardata"`
					Name  string `xml:"name,attr"`
					Flags string `xml:"flags,attr"`
					Epoch string `xml:"epoch,attr"`
					Ver   string `xml:"ver,attr"`
				} `xml:"entry"`
			} `xml:"requires"`
			File []string `xml:"file"`
		} `xml:"format"`
	} `xml:"package"`
}

// Repomd basic struct
type Repomd struct {
	XMLName  xml.Name `xml:"repomd"`
	Text     string   `xml:",chardata"`
	Xmlns    string   `xml:"xmlns,attr"`
	Rpm      string   `xml:"rpm,attr"`
	Revision string   `xml:"revision"`
	Data     []struct {
		Text     string `xml:",chardata"`
		Type     string `xml:"type,attr"`
		Checksum struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
		} `xml:"checksum"`
		OpenChecksum struct {
			Text string `xml:",chardata"`
			Type string `xml:"type,attr"`
		} `xml:"open-checksum"`
		Location struct {
			Text string `xml:",chardata"`
			Href string `xml:"href,attr"`
		} `xml:"location"`
		Timestamp       string `xml:"timestamp"`
		Size            string `xml:"size"`
		OpenSize        string `xml:"open-size"`
		DatabaseVersion string `xml:"database_version"`
	} `xml:"data"`
}
