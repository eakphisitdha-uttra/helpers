package helpers

import (
	"fmt"
	"io"

	"github.com/nwaples/rardecode"
)

// Get filename(s) from within the Archive
func GetRarContents(RARfileName string) (string, error) {

	rr, err := rardecode.OpenReader(RARfileName, "")

	if err != nil {
		return "", fmt.Errorf("read: failed to create reader: %v", err)
	}

	header, err := rr.Next()
	if err == io.EOF {
		return "", fmt.Errorf("archive is empty: %v", err)
	}
	return header.Name, nil
}
