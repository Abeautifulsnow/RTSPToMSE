export interface StreamInfo {
  uuid: string
  name: string
  url: string
  channel: string
}

export interface Streams {
  name: string
  channels: Record<string, Channel>
}

export interface Channel {
  url: string
}

export type infoType = Map<keyof StreamInfo, string>
