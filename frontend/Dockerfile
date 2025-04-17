FROM node:latest AS builder

WORKDIR /app
COPY package*.json .
RUN npm ci
COPY . .
# RUN npx prisma generate
RUN npm run prepare
RUN npm run build
# RUN npm prune --production


# WORKDIR /app
# COPY --from=builder /app/build build/
# COPY --from=builder /app/node_modules node_modules/
# COPY --from=builder /app/prisma prisma/
# COPY package.json .
EXPOSE 3000
# ENV NODE_ENV=production
CMD [ "node", "build" ]