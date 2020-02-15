default: fmt build propagate

fmt:
	go fmt .

build:
	go build -o terraform-provider-harvest

propagate: build
	cp terraform-provider-harvest ~/.terraform.d/plugins/terraform-provider-harvest
