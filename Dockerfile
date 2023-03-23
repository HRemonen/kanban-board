FROM node:latest

WORKDIR /usr/src/app

# COPY --chown=node:node . .

COPY . .

RUN npm i

# USER node

CMD ["npm", "run", "dev"]