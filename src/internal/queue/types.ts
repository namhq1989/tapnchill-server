import { Context } from '@/internal/context'

export interface IRedisConnectOptions {
  url: string
}

export interface IWorkerOptions {
  autorun?: boolean
}

export interface IQueue {
  scheduleJob(
    ctx: Context,
    queueName: string,
    jobName: string,
    cronExpression: string,
    data?: any,
  ): Promise<void>

  createJob(
    ctx: Context,
    queueName: string,
    jobName: string,
    data?: any,
  ): Promise<void>
  processJob(
    ctx: Context,
    queueName: string,
    jobName: string,
    options: IWorkerOptions,
    processor: () => Promise<void>,
  ): void
}
