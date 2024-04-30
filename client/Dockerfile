FROM node:22.0.0-alpine3.19

WORKDIR /espeather

COPY . .

RUN npm install

EXPOSE 5173

CMD [ "npm", "run", "dev", "--", "--host" ]
