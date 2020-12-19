package main

import (
	"os"
	"strings"
)

type Repository struct {
	workTree string
	gitDir string
	conf string
}

var initParams = CmdParams{
	"--bare": makeBareRepo,
	"--separate-git-dir": func() error {},
	"-b": func() error {},
}

func CmdInit(params CmdParams) error {
	workTreePath, getWorkingDirErr := os.Getwd() // NOTE: Might return with / at end
	if getWorkingDirErr != nil {return getWorkingDirErr}

	mkdirErr := os.Mkdir(".git", 0755)
	if mkdirErr != nil {return mkdirErr}
	var gitDirPath strings.Builder
	gitDirPath.WriteString(workTreePath)
	gitDirPath.WriteString("/.git")

	etcConfLocation := ""

	return nil
}

func newRepository(workTree, gitDir, conf string) (Repository, error) {
	_, workTreeErr := os.Stat(workTree)
	_, gitDirErr := os.Stat(gitDir)
	_, confErr := os.Stat(conf)

	if workTreeErr != nil {return _, workTreeErr}
	if gitDirErr != nil {return _, gitDirErr}
	if confErr != nil {return _, confErr}

	return Repository{workTree, gitDir, conf}, _
}

func makeBareRepo() (GitStateComponent, error) {
	gitDir := os.Getenv("GIT_DIR") // NOTE: Check to see if path has / at end
	var repository Repository

	if gitDir == "" {
		workTreePath, getWorkingDirErr := os.Getwd() // NOTE: Might return with / at end
		if getWorkingDirErr != nil {return nil, getWorkingDirErr}

		mkdirErr := os.Mkdir(".git", 0755)
		if mkdirErr != nil {return nil, mkdirErr}
		
		var gitDirPath strings.Builder
		gitDirPath.WriteString(workTreePath)
		gitDirPath.WriteString("/.git")
	} else {
		mkdirErr := os.Mkdir(gitDir, 0755)
		if mkdirErr != nil {return nil, mkdirErr}
	}

	return repository, nil
}

/*
func getWorkingDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {return _, err}

	return dir, err
}
*/