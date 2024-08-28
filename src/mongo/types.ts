import { Db } from 'mongodb'

export interface IMongo {
  connect(): Promise<void>
  disconnect(): Promise<void>
  getDb(): Promise<Db>
}
