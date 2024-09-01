import { Queue as BullMQQueue, Worker } from 'bullmq'
import IORedis from 'ioredis'
import {
  IQueue,
  IQueueDashboardCredentials,
  IRedisConnectOptions,
  IWorkerOptions,
} from '@/internal/queue/types'
import { Context } from '@/internal/context'
import { ExpressAdapter } from '@bull-board/express'
import { createBullBoard } from '@bull-board/api'
import { BullAdapter } from '@bull-board/api/bullAdapter'
import { Express, NextFunction, Request, Response } from 'express'

process.setMaxListeners(30)

class Queue implements IQueue {
  private readonly _redisConnection: IORedis | null = null
  private readonly _dashboardCreds: IQueueDashboardCredentials | null = null
  private readonly _server: Express | null = null
  private readonly _queues: Map<string, BullMQQueue> = new Map()
  private readonly _workers: Map<string, Worker> = new Map()

  constructor(
    options: IRedisConnectOptions,
    dashboardCreds: IQueueDashboardCredentials,
    server: Express,
  ) {
    this._redisConnection = new IORedis(options.url, {
      maxRetriesPerRequest: null,
    })
    this._dashboardCreds = dashboardCreds

    this._server = server
    this._connect()
    this._startDashboard()
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

  private _setDashboardBasicAuth(
    req: Request,
    res: Response,
    next: NextFunction,
  ) {
    const authHeader = req.headers['authorization']
    if (!authHeader) {
      res.setHeader('WWW-Authenticate', 'Basic realm="tapnchill"')
      return res.status(401).send('Authentication required.')
    }

    const base64Credentials = authHeader.split(' ')[1]
    const credentials = Buffer.from(base64Credentials, 'base64').toString(
      'ascii',
    )
    const [username, password] = credentials.split(':')

    const validUsername = this._dashboardCreds!.username
    const validPassword = this._dashboardCreds!.password

    if (username === validUsername && password === validPassword) {
      return next()
    } else {
      return res.status(403).send('Forbidden: Incorrect username or password')
    }
  }

  private _startDashboard() {
    const serverAdapter = new ExpressAdapter()
    serverAdapter.setBasePath('/api/q')

    createBullBoard({
      queues: [
        new BullAdapter(this._getOrCreateQueue('quotes')),
      ],
      serverAdapter,
    })

    this._server!.use(
      '/api/q',
      this._setDashboardBasicAuth.bind(this),
      serverAdapter.getRouter(),
    )
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
