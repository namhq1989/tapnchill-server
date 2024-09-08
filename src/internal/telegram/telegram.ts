import {
  ITelegram,
  ITelegramSendMessageResponse,
} from '@/internal/telegram/types'
import axios from 'axios'

class Telegram implements ITelegram {
  private readonly _botToken: string | null = null
  private readonly _channelId: string | null = null
  private readonly _sendMessageApi = `https://api.telegram.org/bot$botToken/sendMessage`

  constructor(botToken: string, channelId: string) {
    this._botToken = botToken
    this._channelId = channelId
  }

  async sendMessage(message: string): Promise<void> {
    if (!this._botToken || !this._channelId) {
      throw new Error('[telegram] no bot token or channel id provided')
    }

    const url = this._sendMessageApi.replace('$botToken', this._botToken)
    const response = await axios.post<ITelegramSendMessageResponse>(url, {
      chat_id: this._channelId,
      text: `[tapnchill] ${message}`,
    })

    console.log('[telegram] sent message response:', response.data)
  }
}

export default Telegram
