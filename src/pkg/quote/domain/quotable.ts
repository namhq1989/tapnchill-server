class Quotable {
  _id: string
  quote: string
  author: string

  constructor(_id: string, content: string, author: string) {
    this._id = _id
    this.quote = content
    this.author = author
  }
}

export default Quotable
