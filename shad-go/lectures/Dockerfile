FROM golang:1.20-alpine

RUN CGO_ENABLED=0 go install golang.org/x/tools/cmd/present@v0.6.0

COPY . /lectures

WORKDIR /lectures

ENTRYPOINT ["/go/bin/present", "-base", "/go/pkg/mod/golang.org/x/tools@v0.6.0/cmd/present", "-http", ":8080", "-orighost", "p.go.manytask.org", "-use_playground", "-play=0"]
