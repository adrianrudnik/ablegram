package push3api

type FileAllocation struct {
	table map[*FileAllocation]bool
}

type FileAllocationEntry struct {
	Path  string `json:"path"`
	IsDir bool   `json:"is_dir"`
	Size  int    `json:"size"`
}

func NewFileAllocation() *FileAllocation {
	return &FileAllocation{
		table: make(map[*FileAllocation]bool),
	}
}
