import { Queue as BullMQQueue, Worker } from 'bullmq'
import IORedis from 'ioredis'
import {
  IQueue,
  IRedisConnectOptions,
  IWorkerOptions,
} from '@/internal/queue/types'
import { Context } from '@/internal/context'

class Queue implements IQueue {
  private readonly _redisConnection: IORedis | null = null
  private readonly _queues: Map<string, BullMQQueue> = new Map()
  private readonly _workers: Map<string, Worker> = new Map()

  constructor(options: IRedisConnectOptions) {
    this._redisConnection = new IORedis(options.url, {
      maxRetriesPerRequest: null,
    })
    this._connect()
  }

  private _connect(): void {
    try {
      this._redisConnection!.on('error', (error) => {
        console.error('redis connection error:', error)
      })

      console.log('🚀 [queue] connected')
    } catch (error) {
      console.error('failed to establish Redis connection:', error)
      throw new Error('unable to connect to Redis')
    }
  }

  private _getOrCreateQueue(queueName: string): BullMQQueue {
    if (!this._queues.has(queueName)) {
      const queue = new BullMQQueue(queueName, {
        connection: this._redisConnection!,
      })
      this._queues.set(queueName, queue)
    }
    return this._queues.get(queueName)!
  }

  private _getOrCreateWorker(
    ctx: Context,
    queueName: string,
    jobName: string,
    options: IWorkerOptions,
    processor: () => Promise<void>,
  ): Worker {
    if (!this._workers.has(queueName)) {
      const worker = new Worker(
        queueName,
        async () => {
          ctx
            .logger()
            .info(`processing job '${jobName}' in queue '${queueName}'`)
          await processor()
        },
        { connection: this._redisConnection!, ...options },
      )

      // Add worker to the map
      this._workers.set(queueName, worker)

      // Set up graceful shutdown for the worker
      process.on('SIGTERM', async () => {
        await worker.close()
        ctx
          .logger()
          .info(`worker for queue '${queueName}' shut down gracefully`)
      })
    }

    return this._workers.get(queueName)!
  }

  async scheduleJob(
    ctx: Context,
    queueName: string,
    jobName: string,
    cronExpression: string,
    data: any = {},
  ): Promise<void> {
    const queue = this._getOrCreateQueue(queueName)
    await queue.add(jobName, data, {
      repeat: {
        pattern: cronExpression,
        utc: true,
      },
      attempts: 3,
    })

    ctx
      .logger()
      .info(
        `job '${jobName}' scheduled in queue '${queueName}' with cron expression '${cronExpression}'`,
      )

    await queue.close()
  }

  async createJob(
    ctx: Context,
    queueName: string,
    jobName: string,
    data: any = {},
  ): Promise<void> {
    const queue = this._getOrCreateQueue(queueName)
    await queue.add(jobName, data)

    ctx
      .logger()
      .info(
        `job '${jobName}' added to queue '${queueName}' with data ${JSON.stringify(data)}`,
      )

    await queue.close()
  }

  processJob(
    ctx: Context,
    queueName: string,
    jobName: string,
    options: IWorkerOptions,
    processor: () => Promise<void>,
  ): void {
    this._getOrCreateWorker(ctx, queueName, jobName, options, processor)
  }
}

export default Queue
