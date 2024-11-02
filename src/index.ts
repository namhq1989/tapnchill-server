import { IContext } from '@/internal/context/types'
import { Config } from '@/internal/config'
import { Rest } from '@/internal/rest'
import App, { IApp, IModule } from '@/app'
import { Mongo } from '@/internal/mongo'
import { Queue } from '@/internal/queue'
import QuoteModule from '@/pkg/quote'
import { Location } from '@/internal/location'
import Weather from '@/internal/weather/weather'
import Caching from '@/internal/caching/caching'
import WeatherModule from '@/pkg/weather'
import { Realtime } from '@/internal/realtime'
import WebhookModule from '@/pkg/webhook'
import FeedbackModule from '@/pkg/feedback'
import { Telegram } from '@/internal/telegram'

declare global {
  namespace Express {
    interface Request {
      context: IContext
    }
  }
}

// process.on('uncaughtException', (err) => {
//   console.error('Uncaught Exception:', err)
// })
//
// process.on('unhandledRejection', (reason, promise) => {
//   console.error('Unhandled Rejection at:', promise, 'reason:', reason)
// })

// setInterval(() => {
//   const memoryUsage = process.memoryUsage()
//   console.log('Memory Usage:', {
//     rss: memoryUsage.rss / (1024 * 1024) + ' MB',
//     heapTotal: memoryUsage.heapTotal / (1024 * 1024) + ' MB',
//     heapUsed: memoryUsage.heapUsed / (1024 * 1024) + ' MB',
//     external: memoryUsage.external / (1024 * 1024) + ' MB',
//   })
// }, 60000)

const start = async () => {
  const app = await createApp()
  const config = app.getConfig()
  const rest = app.getRest()

  // modules
  const modules: IModule[] = [
    new QuoteModule(),
    new FeedbackModule(),
    new WeatherModule(),
    new WebhookModule(),
  ]
  for (const module of modules) {
    await module.start(app)
  }

  rest.http().listen(config.restPort(), () => {
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

  const queue = new Queue(
    {
      url: config.queueRedisUrl(),
    },
    {
      username: config.queueDashboardUsername(),
      password: config.queueDashboardPassword(),
    },
    rest.server(),
  )

  const caching = new Caching({
    url: config.cachingRedisUrl(),
  })

  const realtime = new Realtime(rest.http())

  const location = new Location(config.ipInfoToken())

  const weather = new Weather(config.visualCrossingToken())

  const telegram = new Telegram(
    config.telegramBotToken(),
    config.telegramChannelId(),
  )

  return new App(
    config,
    rest,
    mongo,
    queue,
    caching,
    realtime,
    location,
    weather,
    telegram,
  )
}

start().then()
