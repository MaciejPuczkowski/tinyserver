FROM golang:latest as base
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ts .

FROM scratch as release
WORKDIR /www
COPY --from=base /app/ts /ts
ENTRYPOINT ["/ts"]
