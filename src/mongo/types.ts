import { Db, ObjectId } from 'mongodb'

export interface IMongoConnectOptions {
  url: string
  dbName: string
}

export interface IMongo {
  connect(options: IMongoConnectOptions): Promise<Db>
  disconnect(): Promise<void>
  getDb(): Promise<Db>
  generateObjectId(): ObjectId
  validateObjectId(id: string): boolean
}
