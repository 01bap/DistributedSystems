FROM node:20 AS builder

WORKDIR /app
COPY package*.json .
RUN npm ci
COPY . .
RUN npm run prepare
RUN npm run build

# CMD [ "node", "build" ]

# --- Production stage ---
FROM node:20 AS runner
WORKDIR /app

# Copy only production files
COPY --from=builder /app/build build/
COPY --from=builder /app/node_modules node_modules/
COPY .env .env

ENV NODE_ENV=production

# Use Node.js 20.6+'s built-in loader or dotenv fallback
# Option 1: Using Node 20.6+ `.env` loader
# CMD ["node", "--env-file=.env", "build"]

# Option 2: If <20.6, fallback to dotenv:
CMD ["node", "-r", "dotenv/config", "build"]