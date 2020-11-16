package main

// VideoInfo ...
type VideoInfo struct {
	VideoID      string        `json:"id"`
	Title        string        `json:"title"`
	Uploader     string        `json:"uploader"`
	UploaderURL  string        `json:"uploader_url"`
	UploaderID   string        `json:"uploader_id"`
	UploadDate   string        `json:"upload_date"`
	Description  string        `json:"description"`
	Categories   []string      `json:"categories"`
	Tags         []string      `json:"tags"`
	Views        int           `json:"view_count"`
	Thumnails    []Thumbnail   `json:"thumbnails"`
	VideoFormats []VideoFormat `json:"formats"`
	Duration     int           `json:"duration"`
}

// Thumbnail ...
type Thumbnail struct {
	ID     string `json:"id"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Res    string `json:"resolution"`
}

// VideoFormat ...
type VideoFormat struct {
	ID         string `json:"format_id"`
	URL        string `json:"url"`
	Ext        string `json:"ext"`
	Resolution string `json:"format_note"`
	FileSize   int    `json:"filesize"`
	Width      int    `json:"width"`
	Height     int    `json:"height"`
	Acodec     string `json:"acodec"`
}
