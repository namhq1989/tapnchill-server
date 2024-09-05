import { IContext } from '@/internal/context/types'
import {
  ISendFeedbackRequestDto,
  ISendFeedbackResponseDto,
} from '@/pkg/feedback/dto/send-feedback'

export interface IFeedbackApplication extends IFeedbackCommandSendFeedback {}

export interface IFeedbackCommandSendFeedback {
  sendFeedback: (
    ctx: IContext,
    performerId: string,
    ip: string,
    req: ISendFeedbackRequestDto,
  ) => Promise<{
    response: ISendFeedbackResponseDto | null
    error: Error | null
  }>
}
