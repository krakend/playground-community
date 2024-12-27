FROM node:23-alpine

WORKDIR /home/app

ADD package.json /home/app
ADD package-lock.json /home/app
RUN npm ci
ADD . /home/app

CMD ["npm", "run", "prod"]

EXPOSE 3000
