import { IMongo, IMongoConnectOptions } from '@/internal/mongo/types'
import { Db, MongoClient, ObjectId } from 'mongodb'

class Mongo implements IMongo {
  private client: MongoClient | null = null
  private db: Db | null = null
  private connectOptions: IMongoConnectOptions | null = null

  async connect(options: IMongoConnectOptions): Promise<void> {
    this.connectOptions = options

    if (!this.client) {
      try {
        this.client = await MongoClient.connect(options.url)
        await this.client.connect()
        await this.client.db().command({ ping: 1 })
      } catch (pingError) {
        await this.client?.close()
        this.client = null
        console.error('Failed to ping MongoDB server', pingError)
        throw new Error('Unable to connect to MongoDB server: ping failed!')
      }
    }

    if (!this.db) {
      this.db = this.client.db(options.dbName)
    }

    console.log('🚀 [mongodb] connected')
  }

  async disconnect(): Promise<void> {
    if (this.client) {
      await this.client.close()
      this.client = null
      this.db = null
    }
  }

  getDb(): Db {
    if (!this.db) {
      throw new Error(
        'Database connection is not established. Please call connect() first!',
      )
    }

    return this.db
  }

  static generateId(): string {
    return new ObjectId().toHexString()
  }

  static validateObjectId(id: string): boolean {
    return ObjectId.isValid(id)
  }
}

export default Mongo
