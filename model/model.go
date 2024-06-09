package model

type Secret struct {
	Secret string `json:"secret,omitempty" bson:"secret,omitempty" query:"secret" url:"secret,omitempty" reqHeader:"secret"`
}

type FaceInfo struct {
	PhoneNumber string `phonenumber:"secret,omitempty" bson:"phonenumber,omitempty"`
	Commit      string `json:"commit,omitempty" bson:"commit,omitempty"`
	Remaining   int    `json:"remaining,omitempty" bson:"remaining,omitempty"`
	FileHash    string `json:"filehash,omitempty" bson:"filehash,omitempty"`
	Error       string `json:"error,omitempty" bson:"error,omitempty"`
}
