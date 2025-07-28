package indexer

import (

	"io/fs"
	"os"
	"path/filepath"
	"github.com/sanxbud/index"
)


func Run(root string) (index.Archive, error) {

	currIndex := index.NewArchive()
	absRoot, err:= filepath.Abs(root)
	if err!= nil{
		return currIndex, err
	}
	filepath.WalkDir(absRoot, func(currPath string, curr fs.DirEntry, err error) error {
		if( !curr.IsDir()){
			//we need to save path, date, type, in a big map to build
			fileType := index.GetType(currPath)
		  info,err := os.Stat(currPath)
			if(err !=nil){
				return err
			}
			date := info.ModTime()

			newEntry := index.IndexEntry{
				Path: currPath, 
				Name: filepath.Base(currPath),
				ModTime: date, 
				Type:fileType,
				}

			currIndex.Add(newEntry)
		}		
		return nil 
	})

	return currIndex, nil

}


