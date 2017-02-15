package fileserver

import (
	"io"

	"github.com/juju/errors"
	minio "github.com/minio/minio-go"
	"github.com/spf13/viper"
)

// SaveFile saves the file
func SaveFile(file io.Reader, userID string) (string, error) {
	endpoint := viper.GetString("minio-endpoint")
	accessKeyID := viper.GetString("minio-access")
	secretAccessKey := viper.GetString("minio-secret")

	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, false)
	if err != nil {
		return "", errors.Annotate(err, "could not create client")
	}

	objName := userID + ".go"

	_, err = minioClient.PutObject(
		"submissions",
		objName,
		file,
		"application/octet-stream",
	)

	if err != nil {
		return "", errors.Trace(err)
	}

	return objName, nil
}
