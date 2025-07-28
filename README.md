## Index
### _a symlink-based virtual archiver for your media files._

Point it at a folder full of images from over the years, in however many subfolders they happen to be in.
It'll run through, analyze them, and build an output directory with all your media files organized by type and date.
Each file in the index is a symlink, meaning you can instantly have an organized copy of your files with minimal additional disk usage, and without disturbing your original folder setup

### Quick Start

- **Download the latest binary** (Linux only for now):  
  [Releases page](https://github.com/sanxbud/index/releases)

- **Or build from source**:  
  git clone https://github.com/sanxbud/index.git  
  cd index  
  go build -o index ./cmd/index  
  ./index --help


### Commands

- `build`  
  Build a symlinked archive from your source folder to a new destination.  
  Example: `index build --source ~/Pictures --dest ~/PicturesIndex`



- `clean`  
  Prune a directory of broken symlinks.  
  Example: `index clean ~/PicturesIndex`


### Roadmap

- Tighten command flags & error handling
- Add a minimal TUI for previewing archive structure
- Explore optional GUI for cross-platform use


### Contributing

This is an early build - feedback, bug reports, and feature ideas are welcome.


_Shout-out to the [boot.dev](https://boot.dev) hackathon for finally getting me started on this!_
