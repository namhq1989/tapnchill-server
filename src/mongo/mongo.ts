import { IMongo } from '@/mongo/types'
import { Db, MongoClient } from 'mongodb'

class MongoDB implements IMongo {
  private client: MongoClient | null = null
  private db: Db | null = null

  async connect(): Promise<void> {
    if (this.db) {
      return
    }
  }

  async disconnect(): Promise<void> {
    if (!this.db) {
      return
    }
  }

  async getDb(): Promise<Db> {
    return this.db!
  }
}

export default MongoDB
