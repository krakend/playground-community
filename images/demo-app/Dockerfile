FROM node:23-alpine

WORKDIR /home/app

ADD package.json .
RUN npm install
COPY . .
RUN npm run build

CMD ["npm", "run", "prod"]

EXPOSE 3000
