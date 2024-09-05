import { Mongo } from '@/internal/mongo'
import { ObjectId } from 'mongodb'

class Feedback {
  readonly _id: ObjectId
  email: string
  feedback: string
  ip: string
  createdAt: Date

  constructor(email: string, feedback: string, ip: string) {
    this._id = Mongo.generateId()
    this.email = email
    this.feedback = feedback
    this.ip = ip
    this.createdAt = new Date()
  }
}

export default Feedback
