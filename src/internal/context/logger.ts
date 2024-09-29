import { ILogger } from '@/internal/context/types'
import { Logger as PinoLogger, pino } from 'pino'

const _logger = pino(
  pino({
    transport:
      process.env.ENV !== 'release'
        ? {
            target: 'pino-pretty',
            options: {
              colorize: true,
            },
          }
        : undefined,
  }),
)

class Logger implements ILogger {
  private logger: PinoLogger | null = null

  constructor(requestId: string, traceId: string) {
    this.logger = _logger.child({
      requestId,
      traceId,
    })
  }

  debug(message: string, data?: object): void {
    this.logger?.debug(data || {}, message)
  }
  info(message: string, data?: object): void {
    this.logger?.info(data || {}, message)
  }
  warn(message: string, data?: object): void {
    this.logger?.warn(data || {}, message)
  }
  error(message: string, error: Error, data?: object): void {
    this.logger?.error({ error: error.message, ...data }, message)
  }
}

export default Logger
