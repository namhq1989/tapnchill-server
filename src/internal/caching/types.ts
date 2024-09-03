export interface IRedisConnectOptions {
  url: string
}

export interface ICaching {
  generateKey: (key: string) => string
  get: (key: string) => Promise<string>
  set: (key: string, value: any, ttl?: number) => Promise<void>
}
