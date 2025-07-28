package cleaner

import (
	"io/fs"
	"fmt"
	"os"
	"path/filepath"
)


func Run(root string) (int,error) {
	count:=0
	absRoot, err:= filepath.Abs(root)
	if err!= nil{
		return 0, err
	}
	filepath.WalkDir(absRoot, func(currPath string, curr fs.DirEntry, err error) error {
		if( !curr.IsDir()){
			info, err := os.Lstat(currPath)
			if err != nil {
				return err
			}

			if info.Mode()&os.ModeSymlink != 0 {
				target,err := os.Readlink(currPath)	
				if err != nil{
					return err
				}
	
				if !filepath.IsAbs(target) {
					target = filepath.Join(filepath.Dir(currPath), target)
				}
				if _, err := os.Stat(target); os.IsNotExist(err) {
            fmt.Printf("Removing broken symlink: %s -> %s\n", currPath, target)
            os.Remove(currPath)
						count++
        }


			}

		}		

		return nil 
	})

	return count, nil
}
