package ktp

type KTPProps struct {
	IDUser    string `json:"iduser,omitempty" bson:"iduser,omitempty"`
	IDFile    string `json:"idfile,omitempty" bson:"idfile,omitempty"`
	Ncard     int    `json:"ncard,omitempty" bson:"ncard,omitempty"`
	Base64Str string `json:"base64str,omitempty" bson:"base64str,omitempty"`
}
