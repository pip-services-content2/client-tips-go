package version1

type AttachmentV1 struct {
	Id   string `json:"id"`
	Uri  string `json:"uri"`
	Name string `json:"name"`
}

func NewAttachmentV1(id, uri, name string) *AttachmentV1 {
	return &AttachmentV1{
		Id:   id,
		Uri:  uri,
		Name: name,
	}
}
