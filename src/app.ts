import { IMongo } from '@/internal/mongo/types'
import { IQueue } from '@/internal/queue/types'
import { IConfig } from '@/internal/config/types'
import { IRest } from '@/internal/rest/types'

export interface IModule {
  name: () => string
  start: (app: IApp) => Promise<Error | null>
}

export interface IApp {
  getConfig: () => IConfig
  getRest: () => IRest
  getMongo: () => IMongo
  getQueue: () => IQueue
}

class App implements IApp {
  private readonly _config: IConfig
  private readonly _rest: IRest
  private readonly _mongo: IMongo
  private readonly _queue: IQueue

  constructor(config: IConfig, rest: IRest, mongo: IMongo, queue: IQueue) {
    this._config = config
    this._rest = rest
    this._mongo = mongo
    this._queue = queue
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
}

export default App
