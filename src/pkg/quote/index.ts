import { IApp, IModule } from '@/app'
import QuoteRepository from '@/pkg/quote/repository/quote-repository'
import QuoteApplication from '@/pkg/quote/application/application'
import QuoteRest from '@/pkg/quote/rest/rest'

class QuoteModule implements IModule {
  name = () => 'Quote'
  start = async (app: IApp) => {
    // dependencies
    const quoteRepository = new QuoteRepository(app.getMongo())

    // application
    const quoteApplication = new QuoteApplication(quoteRepository)

    // rest
    const quoteRest = new QuoteRest(app.getRest(), quoteApplication)
    quoteRest.start()

    console.log('🚀 [quote] started')
    return null
  }
}

export default QuoteModule
