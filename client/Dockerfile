FROM node:16.9.1-alpine

WORKDIR /client

ENV PATH /react-auth/node_modules/.bin:$PATH

COPY package* ./

COPY yarn.lock ./

RUN yarn install 

RUN yarn global add react-scripts@3.4.1

COPY . /client/

CMD ["yarn", "start"]