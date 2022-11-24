FROM arm32v7/node

WORKDIR iot-monopoly

COPY . .
RUN yarn install
RUN yarn build

EXPOSE 8080

ENV PORT=8080
CMD ["node", "build/index.js"]
