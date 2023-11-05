package helpers

import (
	"fmt"
	"os"
	"strings"
)

func ImageDeleter(images []string) error {
	for i := 0; i <= len(images)-1; i++ {
		imageSplit := strings.Split(images[i], "/")
		if err := os.Remove(fmt.Sprintf("./public/uploads/%v", imageSplit[len(imageSplit)-1])); err != nil {
			return err
		}
	}
	return nil
}
