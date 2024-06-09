package model

type Login struct {
	Login string `json:"login,omitempty" bson:"login,omitempty" query:"login" url:"login,omitempty" reqHeader:"login"`
}
