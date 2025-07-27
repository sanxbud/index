package index

import (
	"fmt"
)

type Archive struct {
		
	//map[ File Type ] -> map[ Year ] -> map[ Month ]-[]files
	Data map[string]map[string]map[string][]IndexEntry
}

func NewArchive() Archive {
	return Archive{
		Data:make(map[string]map[string]map[string][]IndexEntry),
	}
}

func (archive *Archive) Add(entry IndexEntry) {
	year := entry.ModTime.Format("2006")
  month := entry.ModTime.Format("01")
  fileType := entry.Type

  if _, exists := archive.Data[fileType]; !exists {
      archive.Data[fileType] = make(map[string]map[string][]IndexEntry)
  }
  if _, exists := archive.Data[fileType][year]; !exists {
      archive.Data[fileType][year] = make(map[string][]IndexEntry)
  }

  archive.Data[fileType][year][month] = append(archive.Data[fileType][year][month], entry)
}

func (archive *Archive) PrintSummary(){
	for fileType := range archive.Data {
		years := archive.Data[fileType]
    		
		for year := range years {
			months := years[year]
			
			for month := range months {
				entries := months[month]
        fmt.Printf("%s/%s/%s → %d files\n", fileType, year, month, len(entries))
      }
    }
	}
}
