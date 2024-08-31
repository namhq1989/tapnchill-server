import { Context } from '@/internal/context'

export interface IRedisConnectOptions {
  url: string
}

export interface IQueue {
  connect(options: IRedisConnectOptions): Promise<void>
  scheduleJob(
    ctx: Context,
    queueName: string,
    jobName: string,
    cronExpression: string,
    data?: any,
  ): Promise<void>
  processJob(
    ctx: Context,
    queueName: string,
    jobName: string,
    processor: () => Promise<void>,
  ): void
}
