FROM golang:1.13.0 As goimage
ENV GO111MODULE=on
WORKDIR /go/src/github.com/kukkar/tigerhall-kittens
COPY . /go/src/github.com/kukkar/tigerhall-kittens
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o tigerhall-kittens main.go 

FROM golang:1.13.0
ENV ENV_FILE_PATH=/etc/kukkar/production.properties
#ENV ELASTIC_APM_SERVER_URL=http://apm-prod-prod-utils-apm-server.prod-utils:8200
#ENV ELASTIC_APM_SERVICE_NAME=MERCHANT-SERVICE

#ENV ELASTIC_APM_SERVER_URL=http://apm.kukkar.in
#ENV ELASTIC_APM_SERVICE_NAME=tigerhall-kittens
#ENV ELASTIC_APM_SECRET_TOKEN=gNSvBzGYqoxh


RUN go get -u github.com/go-sql-driver/mysql
COPY --from=goimage /go/src/github.com/kukkar/tigerhall-kittens/tigerhall-kittens .
COPY --from=goimage /go/src/github.com/kukkar/tigerhall-kittens/conf/ conf/
CMD ["./tigerhall-kittens"]
