FROM node:lts-alpine3.18

WORKDIR /app

COPY package.json .

COPY . .

RUN npm i

RUN npm run build

RUN rm -rf src/ static/

USER node:node

EXPOSE 3000

CMD ["node", "build/index.js"]