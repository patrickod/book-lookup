package main

type BookResults map[string]BookInfo

type BookInfo struct {
	Authors       []Author
	Cover         CoverInfo
	Key           string
	NumberOfPages int
	PublishDate   string
	Subjects      []Subject
	Title         string
	Url           string
}

type Author struct {
	Name string
	Url  string
}

type CoverInfo struct {
	Small  string
	Medium string
	Large  string
}

type Subject struct {
	Name string
	Url  string
}
