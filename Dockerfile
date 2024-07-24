FROM golang:1.19.2 AS build
WORKDIR /build_dir
# Copy dependencies list
COPY go.mod go.sum ./
# Build with optional lambda.norpc tag
COPY . .
RUN go build -tags lambda.norpc -o main main.go
# Copy artifacts to a clean image
FROM public.ecr.aws/lambda/provided:al2023
COPY --from=build /build_dir/main ./main
ENTRYPOINT [ "./main" ]
