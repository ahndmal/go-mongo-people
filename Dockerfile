FROM golang:1.18-rc-alpine
COPY . .
RUN make deps
RUN go build
CMD["./go-mongo-people"]
#CMD["go-mongo-people.exe"] # Windows