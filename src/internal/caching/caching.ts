import { ICaching, IRedisConnectOptions } from '@/internal/caching/types'
import IORedis from 'ioredis'

class Caching implements ICaching {
  private readonly _redis: IORedis | null = null

  constructor(options: IRedisConnectOptions) {
    this._redis = new IORedis(options.url, {
      maxRetriesPerRequest: null,
    })
    this._connect()
  }

  private _connect(): void {
    try {
      this._redis!.on('error', (error) => {
        console.error('redis connection error:', error)
      })

      console.log('🚀 [caching] connected')
    } catch (error) {
      console.error('failed to establish Redis connection:', error)
      throw new Error('unable to connect to Redis')
    }
  }

  generateKey(key: string): string {
    return `tapnchill:caching:${key}`
  }

  async get(key: string): Promise<string> {
    const value = await this._redis!.get(key)
    return value!
  }

  async set(key: string, value: any, ttl?: number): Promise<void> {
    await this._redis!.set(key, value, 'EX', ttl || 3600)
  }
}

export default Caching
