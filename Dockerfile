FROM node:latest

WORKDIR /usr/src/app

# COPY --chown=node:node . .

COPY package.json package-lock.json ./

RUN npm install

COPY . .

# USER node

CMD ["npm", "run", "dev"]