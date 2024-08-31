import { IContext } from '@/internal/context/types'
import { Config } from '@/internal/config'
import { Rest } from '@/internal/rest'
import App, { IApp, IModule } from '@/app'
import { Mongo } from '@/internal/mongo'
import { Queue } from '@/internal/queue'
import QuoteModule from '@/pkg/quote'

declare global {
  namespace Express {
    interface Request {
      context: IContext
    }
  }
}

const start = async () => {
  const app = await createApp()
  const config = app.getConfig()
  const rest = app.getRest()

  // modules
  const modules: IModule[] = [new QuoteModule()]
  for (const module of modules) {
    await module.start(app)
  }

  rest.server().listen(config.restPort(), () => {
    console.log(`🚀 [server] running at http://localhost:${config.restPort()}`)
  })
}

const createApp = async (): Promise<IApp> => {
  const config = new Config()
  const rest = new Rest()

  const mongo = new Mongo({
    url: config.mongoUrl(),
    dbName: config.mongoDbName(),
  })
  await mongo.connect()

  const queue = new Queue({
    url: config.queueRedisUrl(),
  })

  return new App(config, rest, mongo, queue)
}

start().then()
