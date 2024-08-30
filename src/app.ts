import { IMongo } from '@/internal/mongo/types'
import { IQueue } from '@/internal/queue/types'

export interface App {
  getMongo: () => IMongo
  getQueue: () => IQueue
}
