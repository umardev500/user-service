FROM golang as dev

WORKDIR /app

COPY . .

EXPOSE 5013

CMD air
