# use with:
# GOARCH=amd64 GOOS=linux go build -o example example_clean.go
# docker build -t example:latest .
# docker run example:latest

FROM scratch

COPY ./example ./example

CMD ["./example"]
