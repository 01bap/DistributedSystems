FROM node:20 AS builder

WORKDIR /app
COPY package*.json .
RUN npm ci
COPY . .
RUN npm run prepare
RUN npm run build

EXPOSE 3000
CMD [ "node", "build" ]