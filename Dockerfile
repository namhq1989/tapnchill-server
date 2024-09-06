# Stage 1: Build the application
FROM node:20-alpine AS builder

# Install pnpm globally
RUN npm install -g pnpm

# Set the working directory inside the container
WORKDIR /usr/src/app

# Copy package.json and pnpm-lock.yaml for dependency installation
COPY package.json pnpm-lock.yaml ./

# Install dependencies, including tsconfig-paths
RUN pnpm install --frozen-lockfile

# Copy the rest of the application source code
COPY . .

# Build the TypeScript code for production
RUN pnpm build

# Stage 2: Create a minimal production image
FROM node:20-alpine

# Set the working directory inside the container
WORKDIR /usr/src/app

# Copy only the production files from the builder stage
COPY --from=builder /usr/src/app/ .

# Use node to run the compiled app with tsconfig-paths to resolve aliases
CMD ["node", "-r", "ts-node/register/transpile-only", "-r", "tsconfig-paths/register", "dist/src/index.js"]
