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

type LogInfo struct {
	PhoneNumber string `json:"phonenumber,omitempty" bson:"phonenumber,omitempty"`
	Alias       string `json:"alias,omitempty" bson:"alias,omitempty"`
	RepoOrg     string `json:"repoorg,omitempty" bson:"repoorg,omitempty"`
	RepoName    string `json:"reponame,omitempty" bson:"reponame,omitempty"`
	Commit      string `json:"commit,omitempty" bson:"commit,omitempty"`
	Remaining   int    `json:"remaining,omitempty" bson:"remaining,omitempty"`
	FileName    string `json:"filename,omitempty" bson:"filename,omitempty"`
	Base64Str   string `json:"base64str,omitempty" bson:"base64str,omitempty"`
	FileHash    string `json:"filehash,omitempty" bson:"filehash,omitempty"`
	Error       string `json:"error,omitempty" bson:"error,omitempty"`
}
