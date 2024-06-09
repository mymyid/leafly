package ghupload

type GHCreds struct {
	GitHubAccessToken string `json:"githubaccesstoken,omitempty" bson:"githubaccesstoken,omitempty"`
	GitHubAuthorName  string `json:"githubauthorname,omitempty" bson:"githubauthorname,omitempty"`
	GitHubAuthorEmail string `json:"githubauthoremail,omitempty" bson:"githubauthoremail,omitempty"`
}
