import slugify from 'slugify'

const toSlug = (s: string): string => {
  return slugify(s, { lower: true })
}

export { toSlug }
