package main

import (
	"archive/zip"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/hashicorp/go-multierror"
)

type dcpProvider interface {
	process(tmpdir, name string, dcp *DCP) error
}
type upnpdotorg struct {
	DocURL     string // Optional - URL for further documentation about the DCP.
	XMLSpecURL string // Where to download the XML spec from.
	// Any special-case functions to run against the DCP before writing it out.
	Hacks []DCPHackFn
}

func (u upnpdotorg) process(tmpdir, name string, dcp *DCP) error {
	dcp.DocURL = u.DocURL
	specFilename := filepath.Join(tmpdir, name+".zip")
	err := acquireFile(specFilename, u.XMLSpecURL)
	if err != nil {
		return fmt.Errorf("could not acquire spec for %s: %v", name, err)
	}
	archive, err := zip.OpenReader(specFilename)
	if err != nil {
		return fmt.Errorf("error reading zip file %q: %v", specFilename, err)
	}
	defer archive.Close()
	if err := dcp.processZipFile(archive.File); err != nil {
		return fmt.Errorf("error processing spec file %q: %v", specFilename, err)
	}
	for i, hack := range u.Hacks {
		if err := hack(dcp); err != nil {
			return fmt.Errorf("error with Hack[%d] for %s: %v", i, name, err)
		}
	}
	return nil
}

const allSpecsURL = "https://openconnectivity.org/upnp-specs/upnpresources.zip"

type openconnectivitydotorg struct {
	DocPath        string // Optional - Glob to the related documentation about the DCP.
	SpecsURL       string // The HTTP location of the zip archive containing all XML spec.
	XMLSpecZipPath string // Glob to the zip XML spec file.
	XMLSpecPath    string // Glob to the devices and services XMl files.
	// Any special-case functions to run against the DCP before writing it out.
	Hacks []DCPHackFn
}

func (o openconnectivitydotorg) process(tmpdir, name string, dcp *DCP) error {
	allSpecsFilename := filepath.Join(tmpdir, "openconnectivitydotorg_"+name+".zip")
	err := acquireFile(allSpecsFilename, o.SpecsURL)
	if err != nil {
		return fmt.Errorf("could not acquire specs for %s: %v", name, err)
	}
	allSpecsArchive, err := zip.OpenReader(allSpecsFilename)
	if err != nil {
		return fmt.Errorf("error reading zip file %q: %v", allSpecsFilename, err)
	}
	for _, specArchive := range globFiles(o.XMLSpecZipPath, allSpecsArchive.File) {
		f, err := specArchive.Open()
		if err != nil {
			return fmt.Errorf("error reading zip file %q: %v", specArchive.Name, err)
		}
		defer f.Close()
		archive, err := zip.NewReader(newUnbufferedReaderAt(f), specArchive.FileInfo().Size())
		if err != nil {
			return fmt.Errorf("error reading zip file %q: %v", specArchive.Name, err)
		}
		if err := dcp.processZipFile(archive.File); err != nil {
			return fmt.Errorf("error processing spec file %q: %v", specArchive.Name, err)
		}
		for i, hack := range o.Hacks {
			if err := hack(dcp); err != nil {
				return fmt.Errorf("error with Hack[%d] for %s: %v", i, name, err)
			}
		}
	}

	for _, d := range globFiles(o.DocPath, allSpecsArchive.File) {
		dcp.DocURL += d.Name + ", "
	}
	dcp.DocURL = strings.TrimSuffix(dcp.DocURL, ", ")
	return nil
}

type multiProvider []dcpProvider

func (m multiProvider) process(tmpdir, name string, dcp *DCP) (err error) {
	var f int
	for _, p := range m {
		if e := p.process(tmpdir, name, dcp); e != nil {
			f++
			dcp.Reset()
			err = multierror.Append(err, fmt.Errorf("%T: %v", p, e))
			continue
		}
		break
	}
	if f < len(m) {
		return nil
	}
	return
}
