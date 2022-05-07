package model

type Audit struct{
	Err string `bson:"err" json:"err"`
	UserId string  `bson:"userId" json:"userId"`
	Time date.Time `bson:"time" json:"time"`
}
