package indexer

import (

	"io/fs"
	"os"
	"time"
	"strings"
	"path/filepath"
	"github.com/sanxbud/index"
)


var imageTypes = map[string]struct{}{
  ".jpg": {},
  ".jpeg": {},
  ".png": {},
	".bmp": {},
	".gif": {},
	".psd": {},
	".tiff": {},
	".xcf": {},

}

var videoTypes = map[string]struct{}{
  ".mp4": {}, 
	".mov": {},
	".avi": {}, 
	".mkv": {},
}


func Run(root string) (map[string]map[string]map[string][]index.IndexEntry, error) {
									//map[Images/Videos]map[Year]map[Month]-file
	currIndex := make(map[string]map[string]map[string][]index.IndexEntry)
	filepath.WalkDir(root, func(currPath string, curr fs.DirEntry, err error) error {
		if( !curr.IsDir()){
			//we need to save path, date, type, in a big map to build
			fileType := getType(currPath)
		  info,err := os.Stat(currPath)
			if(err !=nil){
				return err
			}
			date := info.ModTime()
			year,month := getDateStrings(date)

			
			if _, exists := currIndex[fileType]; !exists {
					currIndex[fileType] = make(map[string]map[string][]index.IndexEntry)
			}
			if _, exists := currIndex[fileType][year]; !exists {
					currIndex[fileType][year] = make(map[string][]index.IndexEntry)
			}
		
			newEntry := index.IndexEntry{
				Path: currPath, 
				Name: filepath.Base(currPath),
				ModTime: date, 
				Type:fileType,
				}

			currIndex[fileType][year][month] = append(currIndex[fileType][year][month],newEntry)

		}		
		return nil 
	})

	return currIndex, nil

}

func getType(currPath string) string {
	fileType := strings.ToLower(filepath.Ext(currPath))
	if _, ok := imageTypes[fileType]; ok {
		return "Images"	
	} else if _, ok := videoTypes[fileType]; ok {
		return "Videos"	
	}
	return "Other"
}

func getDateStrings(modTime time.Time) (string,string) {
	year :=modTime.Format("2006")
	month :=modTime.Format("01")

	return year, month
}
