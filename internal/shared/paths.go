package shared

// FFmpegSearchPaths lists directories to search for ffmpeg/ffprobe binaries
// on Windows when they are not found via PATH lookup.
var FFmpegSearchPaths = []string{
	`C:\ffmpeg-8.0.1\ffmpeg-2026-01-26-git-fe0813d6e2-full_build\bin`,
	`C:\ffmpeg\bin`,
	`C:\Program Files\ffmpeg\bin`,
	`C:\Program Files (x86)\ffmpeg\bin`,
	`C:\ProgramData\chocolatey\bin`,
}
