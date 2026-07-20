export interface DepStatus {
  name: string
  found: boolean
  version: string
  path: string
}

export interface DepResult {
  ffmpeg: DepStatus
  ffprobe: DepStatus
  yt_dlp: DepStatus
  amf: DepStatus
}

export interface Config {
  output_dir: string
  theme: string
  preferred_encoder: string
}
