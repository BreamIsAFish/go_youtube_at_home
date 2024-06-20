package databaseModel

type Video struct {
	ID          	string    `json:"_id" bson:"_id"`
	Title       	string 		`json:"title" bson:"title"`
	Description 	string 		`json:"description" bson:"description"`
	URL         	string 		`json:"url" bson:"url"`
	ThumbnailURL 	string 		`json:"thumbnail_url" bson:"thumbnail_url"`
	// ThumbnailID 	string 		`json:"thumbnail_id" bson:"thumbnail_id"`
	UploaderID		string    `json:"uploader_id" bson:"uploader_id"`
}