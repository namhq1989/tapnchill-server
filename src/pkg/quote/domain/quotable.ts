class Quotable {
  id: number
  quote: string
  author: string

  constructor(id: number, content: string, author: string) {
    this.id = id
    this.quote = content
    this.author = author
  }
}

export default Quotable
