import { IContext } from '@/internal/context/types'
import {
  IFeedbackApplication,
  IFeedbackCommandSendFeedback,
} from '@/pkg/feedback/application/types'
import { IFeedbackRepository } from '@/pkg/feedback/repository/types'
import FeedbackCommandSendFeedback from '@/pkg/feedback/application/send-feedback'
import { ISendFeedbackRequestDto } from '@/pkg/feedback/dto/send-feedback'

class FeedbackApplication implements IFeedbackApplication {
  private readonly _sendFeedbackHandler: IFeedbackCommandSendFeedback

  constructor(feedbackRepository: IFeedbackRepository) {
    this._sendFeedbackHandler = new FeedbackCommandSendFeedback(
      feedbackRepository,
    )
  }

  sendFeedback(
    ctx: IContext,
    performerId: string,
    ip: string,
    req: ISendFeedbackRequestDto,
  ) {
    return this._sendFeedbackHandler.sendFeedback(ctx, performerId, ip, req)
  }
}

export default FeedbackApplication
