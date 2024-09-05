import { IContext } from '@/internal/context/types'
import Feedback from '@/pkg/feedback/domain/feedback'

export interface IFeedbackRepository {
  create: (ctx: IContext, feedback: Feedback) => Promise<Error | null>
}
