import { IContext } from '@/internal/context/types'
import { Config } from '@/internal/config'
import { Rest } from '@/internal/rest'
import App, { IApp } from '@/app'
import { Mongo } from '@/internal/mongo'
import { Queue } from '@/internal/queue'

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

  rest.server().listen(config.restPort(), () => {
    console.log(`🚀 [server] running at http://localhost:${config.restPort()}`)
  })
}

const createApp = async (): Promise<IApp> => {
  const config = new Config()
  const rest = new Rest()

  const mongo = new Mongo()
  await mongo.connect({
    url: config.mongoUrl(),
    dbName: config.mongoDbName(),
  })

  const queue = new Queue()
  await queue.connect({
    url: config.queueRedisUrl(),
  })

  return new App(config, rest, mongo, queue)
}

start().then()
