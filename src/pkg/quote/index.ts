import { IApp, IModule } from '@/app'
import QuoteRepository from '@/pkg/quote/repository/quote-repository'

class QuoteModule implements IModule {
  name = () => 'Quote'
  start = async (app: IApp) => {
    const quoteRepository = new QuoteRepository(app.getMongo())

    console.log('🚀 [quote] started')
    return null
  }
}

export default QuoteModule
