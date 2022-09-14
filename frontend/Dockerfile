FROM node:14.16.0-alpine3.13
WORKDIR /app
COPY package*.json ./
RUN npm install
RUN npm install -g @angular/cli@10.2.0
COPY . .
EXPOSE 4200
CMD ng serve --host 0.0.0.0