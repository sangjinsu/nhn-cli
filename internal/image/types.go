package image

type Image struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Status     string `json:"status"`
	Visibility string `json:"visibility"`
	MinDisk    int    `json:"min_disk"`
	MinRAM     int    `json:"min_ram"`
	DiskFormat string `json:"disk_format"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type ImageListResponse struct {
	Images []Image `json:"images"`
}
