package app

import "github.com/cglLaLaLa/SoVITS_web/preprocess/app/command"

type Application struct {
	Commands Commands
}

type Commands struct {
	StartSliceCommand command.StartSliceHandler
}
