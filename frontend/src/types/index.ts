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

export interface ExportProgress {
  percentage: number
  fps: number
  speed: string
  eta: string
}

export interface EncoderOption {
  name: string
  value: string
  available: boolean
}

export interface GPUInfo {
  stream_copy: boolean
  encoders: EncoderOption[]
  preferred: string
  gpu_vendor: string
  gpu_available: boolean
}

export interface HistoryEntry {
  id: string
  type: "download" | "export"
  title: string
  source: string
  file_path: string
  file_size: number
  duration: number
  timestamp: string
  status: string
  error?: string
}
