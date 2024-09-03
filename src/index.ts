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
  const modules: IModule[] = [
    new QuoteModule(),
    new WeatherModule(),
  ]
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

  const location = new Location(config.ipInfoToken())

  const weather = new Weather(config.visualCrossingToken())

  return new App(config, rest, mongo, queue, caching, location, weather)
}

start().then()
