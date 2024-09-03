import { IWeatherRest } from '@/pkg/weather/rest/types'
import { IRest } from '@/internal/rest/types'
import { Request, Response, Router } from 'express'
import httpRespond from '@/internal/utils/http-respond'
import { IRealtime } from '@/internal/realtime/types'

class WebhookRest implements IWeatherRest {
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
