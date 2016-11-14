package messages

import "io"

type LoadItemRequest struct {
	correlationID string // required // NOTE: loads shouln't need a correlationID!
	TransactionID string // optional
	Key           string
	Revision      string // optional (provided by KeyMapProjection)
	ETag          string // optional
}
type LoadItemResponse struct {
	correlationID string
	TransactionID string // optional
	ContentLength uint64
	ContentType   string
	Key           string
	ETag          string
	Metadata      map[string]string
	Payload       io.Reader
}
