FROM golang
WORKDIR /go/src/app
COPY backend.go /go/src/app
RUN go build /go/src/app/backend.go
RUN chmod ugo+x /go/src/app/backend
#CMD[/go/src/app/backend]
CMD /go/src/app/backend
