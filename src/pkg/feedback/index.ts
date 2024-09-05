import { IApp, IModule } from '@/app'
import FeedbackRepository from '@/pkg/feedback/repository/feedback-repository'
import FeedbackApplication from '@/pkg/feedback/application/application'
import FeedbackRest from '@/pkg/feedback/rest/rest'

class FeedbackModule implements IModule {
  name = () => 'Feedback'
  start = async (app: IApp) => {
    // dependencies
    const feedbackRepository = new FeedbackRepository(app.getMongo())

    // application
    const feedbackApplication = new FeedbackApplication(feedbackRepository)

    // rest
    const feedbackRest = new FeedbackRest(app.getRest(), feedbackApplication)
    feedbackRest.start()

    console.log('🚀 [pkg feedback] started')
    return null
  }
}

export default FeedbackModule
