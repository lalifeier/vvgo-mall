package types

type Part struct {
	Url  string
	Size int64
	Ext  string
}

type Stream struct {
	Quality string
	Parts   []*Part
	Size    int64
	Ext     string
}

type Data struct {
	Url     string
	Site    string
	Title   string
	Type    string
	Streams map[string]*Stream
}
