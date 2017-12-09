install:
	@go build
test:
	@go test `glide nv` -cover
