import { IMongo, IMongoConnectOptions } from '@/mongo/types'
import { Db, MongoClient, ObjectId } from 'mongodb'

class MongoDB implements IMongo {
  private client: MongoClient | null = null
  private db: Db | null = null
  private connectOptions: IMongoConnectOptions | null = null

  async connect(options: IMongoConnectOptions): Promise<Db> {
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

    return this.db
  }

  async disconnect(): Promise<void> {
    if (this.client) {
      await this.client.close()
      this.client = null
      this.db = null
    }
  }

  async getDb(): Promise<Db> {
    if (!this.db && this.connectOptions) {
      try {
        console.log(
          '[mongo] database connection not found, attempting to reconnect...',
        )
        await this.connect(this.connectOptions)
      } catch (error) {
        console.error('Failed to reconnect to the database', error)
        throw new Error('Failed to connect to the database')
      }
    }

    if (!this.db) {
      throw new Error(
        'Database connection is not established. Please call connect() first!',
      )
    }

    return this.db
  }

  generateObjectId(): ObjectId {
    return new ObjectId()
  }

  validateObjectId(id: string): boolean {
    return ObjectId.isValid(id)
  }
}

export default MongoDB
