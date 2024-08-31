class Quotable {
  _id: string
  content: string
  author: string

  constructor(_id: string, content: string, author: string) {
    this._id = _id
    this.content = content
    this.author = author
  }
}

export default Quotable
