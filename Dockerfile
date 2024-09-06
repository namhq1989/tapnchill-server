FROM node:20-alpine AS builder

RUN npm install -g pnpm pkg

WORKDIR /usr/src/app

COPY package.json pnpm-lock.yaml ./

RUN pnpm install --frozen-lockfile

COPY . .

RUN pnpm build

RUN pkg dist/index.js --targets node18-linux-x64 --output /usr/src/app/server

FROM alpine:3.20

COPY --from=builder /usr/src/app/server /usr/local/bin/server

CMD ["server"]
