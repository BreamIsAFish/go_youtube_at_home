package databaseModel

type User struct {
	ID       string    `json:"_id" bson:"_id"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}