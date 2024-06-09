package model

type Secret struct {
	Secret string `json:"secret,omitempty" bson:"secret,omitempty" query:"secret" url:"secret,omitempty" reqHeader:"secret"`
}
