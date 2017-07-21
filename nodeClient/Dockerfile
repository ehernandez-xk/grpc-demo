FROM node:alpine

WORKDIR /app/nodeClient/

COPY main.js main.js
COPY package.json package.json
COPY init.sh init.sh
COPY service.proto service.proto

RUN apk update ; \
    apk upgrade ; \
    apk add libc6-compat ; \
    npm install

ENTRYPOINT ["node", "main.js"]
