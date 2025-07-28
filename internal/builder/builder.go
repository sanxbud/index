package builder

import (
	"os"
  "path/filepath"
	"github.com/sanxbud/index"
)

var monthMap = map[string]string{
	"01": "January",
	"02": "February",
	"03": "March",
	"04": "April",
	"05": "May",
	"06": "June",
	"07": "July",
	"08": "August",
	"09": "September",
	"10": "October",
	"11": "November",
	"12": "December",
}


func Run(archive index.Archive, destDir string) (int, error) {
		count := 0
		for fileType := range archive.Data {
		years := archive.Data[fileType]
    		
		for year := range years {
			months := years[year]
			
			for month := range months {
				currPath := filepath.Join(destDir, fileType,year,monthMap[month])
				err := os.MkdirAll(currPath,0777)
				if err != nil {
					return 0,err
				}
				entries := months[month]
				for _,entry := range entries{
					target := entry.Path
					dest := filepath.Join(currPath, entry.Name)
					err := os.Symlink(target,dest)
					if err != nil {
						return 0,err
					}
					count++
				}
				
      }
    }
	}

 return count,nil

}
