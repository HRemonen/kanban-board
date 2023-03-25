FROM node:19.7.0-bullseye-slim

WORKDIR /usr/src/app

ADD https://www.google.com /time.now

COPY . .

RUN npm install

CMD ["npm", "run", "dev"]