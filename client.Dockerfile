FROM node:14.15.2-alpine3.10

RUN yarn global add @angular/cli

WORKDIR /app

COPY ./entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

# for window
RUN cat /entrypoint.sh | tr -d '\r' > /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
EXPOSE 3000