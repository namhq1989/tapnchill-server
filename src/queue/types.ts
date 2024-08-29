import { Context } from '@/context'

export interface IRedisConfig {
  url: string
}

export interface IQueue {
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
