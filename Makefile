ENTRY_PATH = cmd/main.go
TERRAFORM_ROOT = infra
ARTIFACT_DIR = $(TERRAFORM_ROOT)/artifacts

# Start only the React Frontend
dev-ui:
	cd frontend && npm run dev

# Start only the Go Backend
dev-api:
	cd backend && go run $(ENTRY_PATH)

dev:
# Run both in parallel
	(cd backend && go run $(ENTRY_PATH)) & (cd frontend && npm run dev)

lint:
	golangci-lint run

kill:
	killall -9 node

tidy:
	cd backend && go mod tidy

build:
# Clean and create build directory
	rm -rf $(ARTIFACT_DIR) && mkdir -p $(ARTIFACT_DIR)
# Build Go binary
	cd backend && GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o ../$(ARTIFACT_DIR)/bootstrap ./$(ENTRY_PATH)
# Zip the binary into a package Terraform can find
	cd $(ARTIFACT_DIR) && zip -r berpadel.zip bootstrap
# Remove configs and bootstrap
	rm -rf $(ARTIFACT_DIR)/bootstrap 

validate:
	terraform fmt && terraform validate

init:
	terraform -chdir=$(TERRAFORM_ROOT) init 

plan:
	terraform -chdir=$(TERRAFORM_ROOT) plan 

apply:
	terraform -chdir=$(TERRAFORM_ROOT) apply -auto-approve 