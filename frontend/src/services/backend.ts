// Backend service — typed wrappers around generated Wails bindings.
// This is the ONLY file that imports from bindings directly.
// All components use this service instead of window.GoApp.

import * as Bindings from "../../bindings/my-clip/internal/app"
import type { DepResult, VideoMetadata, DownloadResult, Config, EncoderOption, GPUInfo } from "../types"

export const Backend = {
  // Core
  GetVersion:      (): Promise<string>                 => Bindings.App.GetVersion() as any,
  GetDependencies: (): Promise<DepResult>               => Bindings.App.GetDependencies() as any,
  RefreshDependencies: (): Promise<DepResult>           => Bindings.App.RefreshDependencies() as any,
  GetConfig:       (): Promise<Config>                  => Bindings.App.GetConfig() as any,
  SaveConfig:      (cfg: Config): Promise<void>         => Bindings.App.SaveConfig(cfg) as any,

  // Source
  SupportedSources: (): Promise<string[]>               => Bindings.App.SupportedSources() as any,
  ResolveSource:   (url: string): Promise<string>       => Bindings.App.ResolveSource(url) as any,
  GetMetadata:     (url: string): Promise<VideoMetadata> => Bindings.App.GetMetadata(url) as any,
  StartDownload:   (url: string, streamID: string): Promise<DownloadResult> =>
                                                           Bindings.App.StartDownload(url, streamID) as any,

  // Media
  ExportFile:       (inputFile: string, encoder: string, format: string): Promise<void> =>
                                                           Bindings.App.ExportFile(inputFile, encoder, format) as any,
  GetGPUInfo:       (): Promise<GPUInfo>               => Bindings.App.GetGPUInfo() as any,
  GetAvailableEncoders: (): Promise<EncoderOption[]>   => Bindings.App.GetAvailableEncoders() as any,
  ProbeFile:        (path: string): Promise<any>        => Bindings.App.ProbeFile(path) as any,
  CreateClip:       (file: string, start: number, end: number): Promise<void> =>
                                                           Bindings.App.CreateClip(file, start, end) as any,

  // Presets
  GetPresets:       (): Promise<any[]>                  => Bindings.App.GetPresets() as any,
  SavePreset:       (preset: any): Promise<void>         => Bindings.App.SavePreset(preset) as any,
  DeletePreset:     (name: string): Promise<void>        => Bindings.App.DeletePreset(name) as any,

  // History
  GetHistory:       (): Promise<any[]>                  => Bindings.App.GetHistory() as any,
  DeleteHistoryEntry: (index: number): Promise<void>    => Bindings.App.DeleteHistoryEntry(index) as any,
  ClearHistory:     (): Promise<void>                   => Bindings.App.ClearHistory() as any,

  // File
  GetOutputDir:     (): Promise<string>                 => Bindings.App.GetOutputDir() as any,
  GetClipOutputDir: (): Promise<string>                 => Bindings.App.GetClipOutputDir() as any,
  OpenFolder:       (path: string): Promise<void>       => Bindings.App.OpenFolder(path) as any,
  GetFileInfo:      (path: string): Promise<any>         => Bindings.App.GetFileInfo(path) as any,
  CopyPathToClipboard: (path: string): Promise<void>    => Bindings.App.CopyPathToClipboard(path) as any,

  // Misc
  ShowNotification: (title: string, message: string): Promise<void> =>
                                                           Bindings.App.ShowNotification(title, message) as any,
  CleanupTemp:      (dir: string): Promise<void>        => Bindings.App.CleanupTemp(dir) as any,
}

