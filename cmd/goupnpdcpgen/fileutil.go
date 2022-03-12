package main

import (
	"archive/zip"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"regexp"
)

func acquireFile(specFilename string, xmlSpecURL string) error {
	if filexists(specFilename) {
		return nil
	}

	tmpFilename := specFilename + ".download"
	if err := downloadFile(tmpFilename, xmlSpecURL); err != nil {
		return err
	}

	return copyFile(specFilename, tmpFilename)
}

func downloadFile(filename, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("could not download %q from %q: %s",
			filename, url, resp.Status)
	}

	w, err := os.Create(filename)
	if err != nil {
		return err
	}

	if _, err := io.Copy(w, resp.Body); err != nil {
		w.Close()
		return err
	}

	return w.Close()
}

func globFiles(pattern string, archive []*zip.File) []*zip.File {
	var files []*zip.File
	for _, f := range archive {
		if matched, err := path.Match(pattern, f.Name); err != nil {
			// This shouldn't happen - all patterns are hard-coded, errors in them
			// are a programming error.
			panic(err)
		} else if matched {
			files = append(files, f)
		}
	}
	return files
}

func unmarshalXmlFile(file *zip.File, data interface{}) error {
	r, err := file.Open()
	if err != nil {
		return err
	}
	decoder := xml.NewDecoder(r)
	defer r.Close()
	return decoder.Decode(data)
}

var scpdFilenameRe = regexp.MustCompile(
	`.*/([a-zA-Z0-9]+)([0-9]+)\.xml`)

func urnPartsFromSCPDFilename(filename string) (*URNParts, error) {
	parts := scpdFilenameRe.FindStringSubmatch(filename)
	if len(parts) != 3 {
		return nil, fmt.Errorf("SCPD filename %q does not have expected number of parts", filename)
	}
	name, version := parts[1], parts[2]
	return &URNParts{
		URN:     serviceURNPrefix + name + ":" + version,
		Name:    name,
		Version: version,
	}, nil
}

func copyFile(dst string, src string) error {
	f, err := os.Open(src)
	if err != nil {
		return err
	}
	return writeFile(dst, f)
}
func writeFile(dst string, r io.ReadCloser) error {
	defer r.Close()
	f, err := os.Create(dst)
	if err != nil {
		return err
	}
	_, err = io.Copy(f, r)
	return err
}

func filexists(p string) bool {
	f, err := os.Open(p)
	if err != nil {
		return !os.IsNotExist(err)
	}
	f.Close()
	return true
}

type unbufferedReaderAt struct {
	R io.Reader
	N int64
}

func newUnbufferedReaderAt(r io.Reader) io.ReaderAt {
	return &unbufferedReaderAt{R: r}
}

func (u *unbufferedReaderAt) ReadAt(p []byte, off int64) (n int, err error) {
	if off < u.N {
		return 0, errors.New("invalid offset")
	}
	diff := off - u.N
	written, err := io.CopyN(ioutil.Discard, u.R, diff)
	u.N += written
	if err != nil {
		return int(written), err
	}

	n, err = u.R.Read(p)
	u.N += int64(n)
	return
}
