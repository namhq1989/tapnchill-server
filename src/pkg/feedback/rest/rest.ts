import { IRest } from '@/internal/rest/types'
import { Request, Response, Router } from 'express'
import httpRespond from '@/internal/utils/http-respond'
import { IFeedbackRest } from '@/pkg/feedback/rest/types'
import { IFeedbackApplication } from '@/pkg/feedback/application/types'
import requestIp from 'request-ip'
import { ISendFeedbackRequestDto } from '@/pkg/feedback/dto/send-feedback'

class FeedbackRest implements IFeedbackRest {
  private readonly _rest: IRest
  private readonly _feedbackApplication: IFeedbackApplication

  constructor(rest: IRest, feedbackApplication: IFeedbackApplication) {
    this._rest = rest
    this._feedbackApplication = feedbackApplication
  }

  start(): void {
    const router = Router()
    this._rest.server().use('/api/feedback', router)

    router.post('', async (req: Request, res: Response) => {
      try {
        const ip = requestIp.getClientIp(req)
        const { response, error } =
          await this._feedbackApplication.sendFeedback(
            req.context,
            '',
            ip!,
            req.body as ISendFeedbackRequestDto,
          )
        if (error) {
          return httpRespond.r400(res, {}, error)
        }

        return httpRespond.r200(res, response!)
      } catch (error) {
        console.error('failed to send feedback:', error)
        return httpRespond.r400(res, {}, error as Error)
      }
    })
  }
}

export default FeedbackRest
