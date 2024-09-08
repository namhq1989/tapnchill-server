export interface ITelegram {
  sendMessage: (message: string) => void
}

export interface ITelegramSendMessageResponse {
  ok: boolean
}
