import { ILogger } from '@/internal/context/types'
import { Logger as PinoLogger, pino } from 'pino'

class Logger implements ILogger {
  private logger: PinoLogger | null = null
  private readonly requestId: string
  private readonly traceId: string

  constructor(requestId: string, traceId: string) {
    this.logger = pino({
      transport:
        process.env.ENV !== 'release'
          ? {
              target: 'pino-pretty',
              options: {
                colorize: true,
              },
            }
          : undefined,
    })

    this.requestId = requestId
    this.traceId = traceId
  }

  debug(message: string, data?: object): void {
    this.logger?.debug(
      {
        requestId: this.requestId,
        traceId: this.traceId,
        data: data || {},
      },
      message,
    )
  }
  info(message: string, data?: object): void {
    this.logger?.info(
      {
        requestId: this.requestId,
        traceId: this.traceId,
        data: data || {},
      },
      message,
    )
  }
  warn(message: string, data?: object): void {
    this.logger?.warn(
      {
        requestId: this.requestId,
        traceId: this.traceId,
        data: data || {},
      },
      message,
    )
  }
  error(message: string, error: Error, data?: object): void {
    this.logger?.error(
      {
        requestId: this.requestId,
        traceId: this.traceId,
        data: data || {},
      },
      message,
    )
  }
}

export default Logger
