FROM node:14-alpine as builder
WORKDIR /app
COPY package.json yarn.lock ./
RUN yarn
COPY ./ .
RUN yarn build

FROM nginx:alpine as production
RUN mkdir /app
COPY --from=builder /app/dist /app
COPY ./nginx/nginx.conf /etc/nginx/nginx.conf

