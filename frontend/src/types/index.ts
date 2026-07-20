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

export interface StreamInfo {
  id: string
  quality: string
  resolution: string
  format: string
  size: number
  has_audio: boolean
  has_video: boolean
  bitrate: number
  codec: string
}

export interface VideoMetadata {
  source: string
  title: string
  author: string
  duration: number
  thumbnail: string
  url: string
  streams: StreamInfo[]
}

export interface DownloadProgress {
  percentage: number
  speed: string
  eta: string
  bytes_loaded: number
  total_bytes: number
}

export interface DownloadResult {
  file_path: string
  size: number
  duration: number
  format: string
}
