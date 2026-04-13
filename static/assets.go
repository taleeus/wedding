package static

import "embed"

//go:embed css/* scripts/* img/*
var Assets embed.FS
