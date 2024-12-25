package web

import "embed"

//go:embed build/*
var Fs embed.FS
