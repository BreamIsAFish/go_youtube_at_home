package requestModel

// type VideoUploadRequest struct {
// 	Title       	string 		`json:"title xml:"title" form:"title" validate:"required"`
// 	// Description 	string 		`json:"description" xml:"description" form:"description"`
// 	Description 	string 		`json:"description" xml:"description" form:"description"`
// 	// VideoData			*multipart.FileHeader		`json:"video_data" xml="video_data" form:"video_data" validate:"required"`
// 	// ThumbnailData	*multipart.FileHeader		`json:"thumbnail_data" xml="thumbnail_data" form:"thumbnail_data" validate:"required"`
// 	// VideoData			string		`json:"video_data" xml="video_data" form:"video_data" validate:"required"`
// 	// ThumbnailData	string		`json:"thumbnail_data" xml="thumbnail_data" form:"thumbnail_data" validate:"required"`
// 	// VideoData			*multipart.Form		`form:"video_data"`
// 	// ThumbnailData	*multipart.Form		`form:"thumbnail_data"`
// 	// URL         	string 		`json:"url"`
// 	Data     multipart.File	`json:"data" xml:"data" form:"data" validate:"required"`
// 	// ThumbnailURL 	string 		`json:"thumbnail_url"`
// }

type VideoUploadRequest struct {
    Title       string                `json:"title" xml:"title" form:"title" validate:"required"`
    Description string                `json:"description" xml:"description" form:"description"`
    // Data        *multipart.FileHeader `xml:"data" form:"data" validate:"required"`
}