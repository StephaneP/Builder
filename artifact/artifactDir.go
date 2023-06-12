package artifact

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func ArtifactDir() {
	var dirPath string

	dirPath = os.Getenv("BUILDER_PARENT_DIR")
	dirName := os.Getenv("BUILDER_DIR_NAME")

	currentTime := time.Now().Unix()
	artifactStamp := dirName + "_artifact_" + strconv.FormatInt(currentTime, 10)
	os.Setenv("BUILDER_ARTIFACT_STAMP", artifactStamp)
	artifactDir := dirPath + "/" + artifactStamp

	err := os.Mkdir(artifactDir, 0755)
	if err != nil {
		log.Fatal(err)
	}

	val, present := os.LookupEnv("BUILDER_ARTIFACT_DIR")
	if !present {
		os.Setenv("BUILDER_ARTIFACT_DIR", artifactDir)
	} else {
		fmt.Println("BUILDER_ARTIFACT_DIR", val)
	}
}
