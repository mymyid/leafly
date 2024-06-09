package face

type FaceDetect struct {
	IDUser    string `json:"iduser,omitempty" bson:"iduser,omitempty"`
	IDFile    string `json:"idfile,omitempty" bson:"idfile,omitempty"`
	Nfaces    int    `json:"nfaces,omitempty" bson:"nfaces,omitempty"`
	Base64Str string `json:"base64str,omitempty" bson:"base64str,omitempty"`
}
