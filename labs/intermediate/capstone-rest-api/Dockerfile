# TODO: replace with a multi-stage, non-root build.
FROM golang:1.26 AS build
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build -o /out/api ./intermediate/capstone-rest-api/cmd/api

FROM gcr.io/distroless/static-debian12
COPY --from=build /out/api /api
USER nonroot:nonroot
ENTRYPOINT ["/api"]
