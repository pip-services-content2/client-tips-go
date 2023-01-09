package version1

import "time"

type TipV1 struct {
	// Identification
	Id     string   `json:"id"`
	Topics []string `json:"topics"`

	// Automatically managed fields
	Creator    *PartyReferenceV1 `json:"creator"`
	CreateTime time.Time         `json:"create_time"`

	// Content
	Title   map[string]string `json:"title"`
	Content map[string]string `json:"content"`
	MoreUrl string            `json:"more_url"`
	Pics    []*AttachmentV1   `json:"pics"`
	Docs    []*AttachmentV1   `json:"docs"`

	// Search
	Tags    []string `json:"tags"`
	AllTags []string `json:"all_tags"`

	// Status
	Status string `json:"status"`

	// Custom fields
	CustomHdr any `json:"custom_hdr"`
	CustomDat any `json:"custom_dat"`
}
