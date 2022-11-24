FROM node:16

WORKDIR iot-monopoly

COPY . .
RUN yarn install
RUN yarn build

EXPOSE 8080

ENV PORT=8080
CMD ["node", "build/server.js"]
