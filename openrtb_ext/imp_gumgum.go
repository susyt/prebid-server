package openrtb_ext

// ExtImpGumGum defines the contract for bidrequest.imp[i].ext.gumgum
// Either Zone or PubId must be present, others are optional parameters
type ExtImpGumGum struct {
	Zone   string  `json:"zone,omitempty"`
	PubId  float64 `json:"pubId,omitempty"`
	Irisid string  `json:"irisid,omitempty"`
}

// ExtImpGumGumVideo defines the contract for bidresponse.seatbid.bid[i].ext.gumgum.video
type ExtImpGumGumVideo struct {
	Irisid string `json:"irisid,omitempty"`
}
