FROM alpine:3.21

WORKDIR /app

RUN apk add --no-cache bash

COPY scripts/ /app/scripts/
RUN chmod +x /app/scripts/*.sh

CMD ["/bin/sh"]