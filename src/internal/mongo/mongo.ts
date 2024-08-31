import { IMongo, IMongoConnectOptions } from '@/internal/mongo/types'
import { Db, MongoClient, ObjectId } from 'mongodb'

class Mongo implements IMongo {
  private _client: MongoClient | null = null
  private _db: Db | null = null
  private _connectOptions: IMongoConnectOptions | null = null

  constructor(options: IMongoConnectOptions) {
    this._connectOptions = options
  }

  async connect(): Promise<void> {
    if (!this._client) {
      try {
        this._client = await MongoClient.connect(this._connectOptions!.url)
        await this._client.connect()
        await this._client.db().command({ ping: 1 })
      } catch (pingError) {
        await this._client?.close()
        this._client = null
        console.error('Failed to ping MongoDB server', pingError)
        throw new Error('Unable to connect to MongoDB server: ping failed!')
      }
    }

    if (!this._db) {
      this._db = this._client.db(this._connectOptions!.dbName)
    }

    console.log('🚀 [mongodb] connected')
  }

  async disconnect(): Promise<void> {
    if (this._client) {
      await this._client.close()
      this._client = null
      this._db = null
    }
  }

  getDb(): Db {
    if (!this._db) {
      throw new Error(
        'Database connection is not established. Please call connect() first!',
      )
    }

    return this._db
  }

  static generateId(): ObjectId {
    return new ObjectId()
  }

  static validateObjectId(id: string): boolean {
    return ObjectId.isValid(id)
  }
}

export default Mongo
