import { IRest } from '@/internal/rest/types'
import { Request, Response, Router } from 'express'
import { IRealtime } from '@/internal/realtime/types'
import { IWebhookRest } from '@/pkg/webhook/rest/types'
import httpRespond from '@/internal/utils/http-respond'

class WebhookRest implements IWebhookRest {
  private readonly _rest: IRest
  private readonly _realtime: IRealtime

  constructor(rest: IRest, realtime: IRealtime) {
    this._rest = rest
    this._realtime = realtime
  }

  start(): void {
    const router = Router()
    this._rest.server().use('/api/webhook', router)

    router.post('/channel-stats', async (req: Request, res: Response) => {
      this._realtime.broadcastAll('channel-stats', req.body)
      return httpRespond.r200(res, {})
    })
  }
}

export default WebhookRest
