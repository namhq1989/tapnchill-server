export interface IContext {
  logger: () => ILogger
}

export interface ILogger {
  debug: (message: string, data?: object) => void
  info: (message: string, data?: object) => void
  warn: (message: string, data?: object) => void
  error: (message: string, error: Error, data?: object) => void
}
