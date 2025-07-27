package index

import (
	"strings"
	"path/filepath"
	"time"
)

var ImageTypes = map[string]struct{}{
  ".jpg": {},
  ".jpeg": {},
  ".png": {},
	".bmp": {},
	".gif": {},
	".psd": {},
	".tiff": {},
	".xcf": {},

}

var VideoTypes = map[string]struct{}{
  ".mp4": {}, 
	".mov": {},
	".avi": {}, 
	".mkv": {},
}


func GetType(currPath string) string {
	fileType := strings.ToLower(filepath.Ext(currPath))
	if _, ok := ImageTypes[fileType]; ok {
		return "Images"	
	} else if _, ok := VideoTypes[fileType]; ok {
		return "Videos"	
	}
	return "Other"
}

func GetDateStrings(modTime time.Time) (string,string) {
	year :=modTime.Format("2006")
	month :=modTime.Format("01")

	return year, month
}
