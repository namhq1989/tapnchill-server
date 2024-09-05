import { IContext } from '@/internal/context/types'
import { IFeedbackCommandSendFeedback } from '@/pkg/feedback/application/types'
import { IFeedbackRepository } from '@/pkg/feedback/repository/types'
import {
  ISendFeedbackRequestDto,
  ISendFeedbackResponseDto,
} from '@/pkg/feedback/dto/send-feedback'
import Feedback from '@/pkg/feedback/domain/feedback'

class FeedbackCommandSendFeedback implements IFeedbackCommandSendFeedback {
  private readonly _feedbackRepository: IFeedbackRepository

  constructor(feedbackRepository: IFeedbackRepository) {
    this._feedbackRepository = feedbackRepository
  }

  async sendFeedback(
    ctx: IContext,
    performerId: string,
    ip: string,
    req: ISendFeedbackRequestDto,
  ): Promise<{
    response: ISendFeedbackResponseDto | null
    error: Error | null
  }> {
    ctx
      .logger()
      .info('[api] new send feedback request', { performerId, ip, req })

    ctx.logger().info('new feedback model')
    const feedback = new Feedback(req.email, req.feedback, ip)

    ctx.logger().info('persist feedback in db')
    const error = await this._feedbackRepository.create(ctx, feedback)
    if (error) {
      ctx.logger().error('failed to persist feedback in db', error)
      return {
        response: null,
        error,
      }
    }

    return {
      response: {
        ok: true,
      },
      error: null,
    }
  }
}

export default FeedbackCommandSendFeedback
