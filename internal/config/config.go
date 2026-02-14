package config

import (
	"flag"
	"os"
)

type Config struct {
	StoragePath string
}

func Load(args []string) (*Config, []string, error) {

	var file string
	fs := flag.NewFlagSet("global", flag.ContinueOnError)
	fs.StringVar(&file, "file", "", "storage path")

	_ = fs.Parse(args)

	path := resolveStoragePath(file)

	return &Config{
		StoragePath: path,
	}, fs.Args(), nil
}

func resolveStoragePath(file string) string {
	if file != "" {
		return file
	}

	if env := os.Getenv("TODO_STORAGE_PATH"); env != "" {
		return env
	}

	return "tasks.json"
}
