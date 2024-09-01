export interface IConfig {
  restPort: () => number

  queueRedisUrl: () => string
  queueDashboardUsername: () => string
  queueDashboardPassword: () => string

  mongoUrl: () => string
  mongoDbName: () => string
}
