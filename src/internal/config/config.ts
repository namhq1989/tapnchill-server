import { IConfig } from '@/internal/config/types'
import * as process from 'node:process'

class Config implements IConfig {
  private readonly _restPort: number

  private readonly _queueRedisUrl?: string
  private readonly _queueDashboardUsername?: string
  private readonly _queueDashboardPassword?: string

  private readonly _mongoUrl?: string
  private readonly _mongoDbName?: string

  private readonly _cachingRedisUrl?: string

  private readonly _ipInfoToken?: string

  private readonly _visualCrossingToken?: string

  constructor() {
    this._restPort = process.env.REST_PORT
      ? Number(process.env.REST_PORT)
      : 3000

    this._queueRedisUrl = process.env.QUEUE_REDIS_URL
    this._queueDashboardUsername = process.env.QUEUE_DASHBOARD_USERNAME
    this._queueDashboardPassword = process.env.QUEUE_DASHBOARD_PASSWORD

    this._mongoUrl = process.env.MONGO_URL
    this._mongoDbName = process.env.MONGO_DB_NAME

    this._cachingRedisUrl = process.env.CACHING_REDIS_URL

    this._ipInfoToken = process.env.IP_INFO_TOKEN

    this._visualCrossingToken = process.env.VISUAL_CROSSING_TOKEN

    const error = this.validate()
    if (error) {
      console.log('config validation error:', error)
      process.exit(1)
    }
  }

  validate(): Error | null {
    if (!this._queueRedisUrl) {
      return Error('Missing required QUEUE_REDIS_URL environment variable')
    }
    if (!this._mongoUrl) {
      return Error('Missing required MONGO_URL environment variable')
    }
    if (!this._mongoDbName) {
      return Error('Missing required MONGO_DB_NAME environment variable')
    }
    if (!this._cachingRedisUrl) {
      return Error('Missing required CACHING_REDIS_URL environment variable')
    }
    if (!this._ipInfoToken) {
      return Error('Missing required IP_INFO_TOKEN environment variable')
    }
    if (!this._visualCrossingToken) {
      return Error(
        'Missing required VISUAL_CROSSING_TOKEN environment variable',
      )
    }

    return null
  }

  restPort(): number {
    return this._restPort
  }

  queueRedisUrl(): string {
    return this._queueRedisUrl!
  }

  queueDashboardUsername(): string {
    return this._queueDashboardUsername!
  }

  queueDashboardPassword(): string {
    return this._queueDashboardPassword!
  }

  mongoUrl(): string {
    return this._mongoUrl!
  }

  mongoDbName(): string {
    return this._mongoDbName!
  }

  cachingRedisUrl(): string {
    return this._cachingRedisUrl!
  }

  ipInfoToken(): string {
    return this._ipInfoToken!
  }

  visualCrossingToken(): string {
    return this._visualCrossingToken!
  }
}

export default Config
