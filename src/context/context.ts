import { IContext, ILogger } from '@/context/types'
import Logger from '@/context/logger'
import { randomUUID } from '@/utils/uuid'

class Context implements IContext {
  private readonly l: ILogger | null = null
  private readonly requestId: string
  private readonly traceId: string

  constructor() {
    this.requestId = randomUUID()
    this.traceId = randomUUID()
    this.l = new Logger(this.requestId, this.traceId)
  }

  logger(): ILogger {
    return this.l!
  }
}

export default Context
