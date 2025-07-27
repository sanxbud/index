package index

import "time"

type IndexEntry struct {
  Path string
  Name string
	ModTime time.Time
	Type string
}
