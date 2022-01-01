FROM node:16-alpine
WORKDIR /app
COPY ./dist .
CMD ["node", "/app/src/index.js"]