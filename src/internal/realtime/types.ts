export interface IRealtime {
  broadcastAll: (event: string, data: object) => void
}
