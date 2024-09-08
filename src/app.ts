import { IMongo } from '@/internal/mongo/types'
import { IQueue } from '@/internal/queue/types'
import { IConfig } from '@/internal/config/types'
import { IRest } from '@/internal/rest/types'
import { ILocation } from '@/internal/location/types'
import { IWeather } from '@/internal/weather/types'
import { ICaching } from '@/internal/caching/types'
import { IRealtime } from '@/internal/realtime/types'
import { ITelegram } from '@/internal/telegram/types'

export interface IModule {
  name: () => string
  start: (app: IApp) => Promise<Error | null>
}

export interface IApp {
  getConfig: () => IConfig
  getRest: () => IRest
  getMongo: () => IMongo
  getQueue: () => IQueue
  getCaching: () => ICaching
  getRealtime: () => IRealtime
  getLocation: () => ILocation
  getWeather: () => IWeather
  getTelegram: () => ITelegram
}

class App implements IApp {
  private readonly _config: IConfig
  private readonly _rest: IRest
  private readonly _mongo: IMongo
  private readonly _queue: IQueue
  private readonly _caching: ICaching
  private readonly _realtime: IRealtime
  private readonly _location: ILocation
  private readonly _weather: IWeather
  private readonly _telegram: ITelegram

  constructor(
    config: IConfig,
    rest: IRest,
    mongo: IMongo,
    queue: IQueue,
    caching: ICaching,
    realtime: IRealtime,
    location: ILocation,
    weather: IWeather,
    telegram: ITelegram,
  ) {
    this._config = config
    this._rest = rest
    this._mongo = mongo
    this._queue = queue
    this._caching = caching
    this._realtime = realtime
    this._location = location
    this._weather = weather
    this._telegram = telegram
  }

  getConfig(): IConfig {
    return this._config
  }

  getRest(): IRest {
    return this._rest
  }

  getMongo(): IMongo {
    return this._mongo
  }

  getQueue(): IQueue {
    return this._queue
  }

  getCaching(): ICaching {
    return this._caching
  }

  getRealtime(): IRealtime {
    return this._realtime
  }

  getLocation(): ILocation {
    return this._location
  }

  getWeather(): IWeather {
    return this._weather
  }

  getTelegram(): ITelegram {
    return this._telegram
  }
}

export default App
