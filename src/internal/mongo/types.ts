import { Db } from 'mongodb'

export interface IMongoConnectOptions {
  url: string
  dbName: string
}

export interface IMongo {
  connect(options: IMongoConnectOptions): Promise<void>
  disconnect(): Promise<void>
  getDb(): Db
}
