import { IContext, ILogger } from '@/internal/context/types'
import Logger from '@/internal/context/logger'
import { randomUUID } from '@/internal/utils/uuid'

class Context implements IContext {
  private l: ILogger | null = null
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

  destroy(): void {
    this.l = null
  }
}

export default Context
