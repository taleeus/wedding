package static

import "embed"

//go:embed css/* scripts/* img/* fonts/*
var Assets embed.FS
