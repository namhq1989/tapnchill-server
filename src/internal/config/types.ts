export interface IConfig {
  restPort: () => number

  queueRedisUrl: () => string

  mongoUrl: () => string
  mongoDbName: () => string
}
