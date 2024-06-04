FROM node:20-alpine as build

WORKDIR /usr/local/app

COPY package*.json ./

RUN npm install -g @angular/cli

RUN npm install

COPY . .

RUN ng build

FROM nginx:alpine

COPY --from=build /usr/local/app/dist/test/browser /usr/share/nginx/html

EXPOSE 4200

CMD ["nginx", "-g", "daemon off;"]

