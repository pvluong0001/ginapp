FROM node:14.15.2-alpine3.10

RUN yarn global add @angular/cli

WORKDIR /app

COPY ./entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
EXPOSE 3000