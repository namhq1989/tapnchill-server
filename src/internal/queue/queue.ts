import { Queue as BullMQQueue, Worker } from 'bullmq'
import IORedis from 'ioredis'
import { IQueue, IRedisConnectOptions } from '@/internal/queue/types'
import { Context } from '@/internal/context'

class Queue implements IQueue {
  private _redisConnection: IORedis | null = null

  async connect(options: IRedisConnectOptions): Promise<void> {
    try {
      this._redisConnection = new IORedis(options.url)

      this._redisConnection.on('error', (error) => {
        console.error('Redis connection error:', error)
      })

      console.log('🚀 [queue] connected')
    } catch (error) {
      console.error('Failed to establish Redis connection:', error)
      throw new Error('Unable to connect to Redis')
    }
  }

  async scheduleJob(
    ctx: Context,
    queueName: string,
    jobName: string,
    cronExpression: string,
    data: any = {},
  ): Promise<void> {
    const queue = new BullMQQueue(queueName, {
      connection: this._redisConnection!,
    })

    await queue.add(jobName, data, {
      repeat: {
        pattern: cronExpression,
      },
    })

    ctx
      .logger()
      .info(
        `Job '${jobName}' scheduled in queue '${queueName}' with cron expression '${cronExpression}'`,
      )

    await queue.close()
  }

  processJob(
    ctx: Context,
    queueName: string,
    jobName: string,
    processor: () => Promise<void>,
  ): void {
    const worker = new Worker(
      queueName,
      async () => {
        ctx.logger().info(`Processing job '${jobName}' in queue '${queueName}'`)
        await processor()
      },
      { connection: this._redisConnection! },
    )

    // graceful shutdown
    process.on('SIGTERM', async () => {
      await worker.close()
      await this._redisConnection!.quit()
      ctx.logger().info(`Worker for queue '${queueName}' shut down gracefully`)
    })
  }
}

export default Queue
