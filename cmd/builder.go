package cmd

import (
	"Builder/derive"
	"Builder/directory"
	"Builder/logger"
	"Builder/utils"
	"Builder/yaml"
	"os"
	"os/exec"
	"path/filepath"
)

func Builder(log logger.Logger) {
	os.Setenv("BUILDER_COMMAND", "true")
	path, _ := os.Getwd()

	builderPath := filepath.Join(path, "builder.yaml")

	if _, err := os.Stat(builderPath); err == nil {
		exec.Command("git", "pull").Run()

		yaml.YamlParser(builderPath)

		directory.MakeDirs()
		log.Info("Directories successfully created.")

		utils.CloneRepo()
		log.Info("Repo cloned successfully.")

		derive.ProjectType()

		utils.Docker()

		utils.MakeHidden()
		log.Info("Hidden Dir is now read-only.")
	} else {
		utils.Help()
		log.Fatal("bulder.yaml file not found. cd into it's location.")
	}
}
