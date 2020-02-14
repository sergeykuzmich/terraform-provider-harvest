default: fmt build copy

fmt:
	go fmt .

build: fmt
	go build -o terraform-provider-harvest

copy: build
	cp terraform-provider-harvest ~/.terraform.d/plugins/terraform-provider-harvest
