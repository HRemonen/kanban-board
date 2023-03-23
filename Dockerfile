FROM node:latest

WORKDIR /usr/src/app

# COPY --chown=node:node . .

COPY . .

RUN npm install

# USER node

CMD ["npm", "run", "dev"]