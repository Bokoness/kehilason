FROM node

WORKDIR /app

ARG PORT

COPY package.json .

RUN npm install

COPY . .

EXPOSE $PORT

CMD [ "npm", "start" ]