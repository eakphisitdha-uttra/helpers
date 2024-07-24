package helpers

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"os"
	"strings"
)

func UploadBase64(base64Code, fileName, filePath string) (string, string, int64, error) {

	base64Data := strings.Split(base64Code, ",")
	if len(base64Data) > 1 {
		base64Code = base64Data[1]
	}

	imgBytes, err := base64.StdEncoding.DecodeString(strings.TrimSpace(base64Code))
	if err != nil {
		return "", "", 0, err
	}
	_, imageType, err := image.Decode(bytes.NewReader(imgBytes))
	if err != nil {
		return "", "", 0, err
	}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		err := os.MkdirAll(filePath, 0755)
		if err != nil {
			return "", "", 0, err
		}
	}
	imgPath := fmt.Sprintf("%s/%s.%s", filePath, fileName, imageType)

	// Create the image file.
	imgFile, err := os.Create(imgPath)
	if err != nil {
		return "", "", 0, err
	}
	defer imgFile.Close()

	// Write the bytes into the image file.
	_, err = imgFile.Write(imgBytes)
	if err != nil {
		return "", "", 0, err
	}
	// Get the file size
	fileInfo, err := imgFile.Stat()
	if err != nil {
		return "", "", 0, err
	}
	fileSize := fileInfo.Size()
	return fmt.Sprintf("%s.%s", fileName, imageType), imageType, fileSize, nil
}
