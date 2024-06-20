package repository

import (
	"context"
	"go_youtube_at_home/configs"
	"go_youtube_at_home/internal/domain"
	databaseModel "go_youtube_at_home/internal/model/database_model"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// const (
// 	videoStorageDir = "/videos/"
// )

type videoRepository struct {
	collection 		*mongo.Collection
	mongoTimeout 	int
}

func NewVideoRepository(collection *mongo.Collection) domain.VideoRepository {
	return &videoRepository{
		collection: collection,
		mongoTimeout: configs.GetConfig().MongoDB.Timeout,
	}
}

func (r *videoRepository) SaveFileToStorage(name string, data *multipart.FileHeader) (string, error) {
	// Get current directory
	// _, filename, _, ok := runtime.Caller(0)
	// if !ok {
	// 	log.Fatal("Failed to get current frame")
	// }
	// dir := filepath.Dir(filename)
	// if !strings.HasSuffix(dir, "/storage") {
	// 	dir += "/"
	// }
	// fmt.Println(dir)
	// fmt.Println(dir + name)
	// Save video to storage
	// out, err := os.Create(dir + name)
	filePath := filepath.Join("./storage", name)
	// filePath := "/storage/" + name
	out, err := os.Create(filePath)
	if err != nil {
			return "", err
	}
	defer out.Close()

	// Open the file referenced by the FileHeader, which returns an io.ReadCloser
	file, err := data.Open()
	if err != nil {
			return "", err // Handle the error appropriately
	}
	defer file.Close() 

	_, err = io.Copy(out, file)
	return filePath, err
}

// func (r *videoRepository) GetFileFromStorage(fileType string, name string) (*multipart.FileHeader, error) {
// 	// Get current directory
// 	_, filename, _, ok := runtime.Caller(0)
// 	if !ok {
// 		log.Fatal("Failed to get current frame")
// 	}
// 	dir := filepath.Dir(filename)
// 	if !strings.HasSuffix(dir, "/"+fileType+"/") {
// 		dir += "/"
// 	}
// 	// Get video from storage
// 	video, err := os.Open(dir + name)
// 	if err != nil {
// 			return nil, err
// 	}

// 	// Open the file referenced by the FileHeader, which returns an io.ReadCloser
// 	// file, err := data.Open()
// 	// if err != nil {
// 	// 		return err // Handle the error appropriately
// 	// }
// 	// defer file.Close() 

// 	return video, nil
// }

func (r *videoRepository) CreateVideo(video *databaseModel.Video) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(r.mongoTimeout)*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, video)
	return err
}

func (r *videoRepository) GetVideosByUserID(userID string) ([]*databaseModel.Video, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(r.mongoTimeout)*time.Second)
	defer cancel()
	
	var videos []*databaseModel.Video
    cursor, err := r.collection.Find(ctx, bson.M{"uploader_id": userID})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    for cursor.Next(ctx) {
        var video databaseModel.Video
        if err := cursor.Decode(&video); err != nil {
            return nil, err
        }
        videos = append(videos, &video)
    }

    if err := cursor.Err(); err != nil {
        return nil, err
    }

    return videos, nil
}

func (r *videoRepository) GetVideos() ([]*databaseModel.Video, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(r.mongoTimeout)*time.Second)
	defer cancel()
	
	var videos []*databaseModel.Video
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
			return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
			var video databaseModel.Video
			if err := cursor.Decode(&video); err != nil {
					return nil, err
			}
			videos = append(videos, &video)
	}

	if err := cursor.Err(); err != nil {
			return nil, err
	}

	return videos, nil
}