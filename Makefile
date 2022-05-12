clean:
	rm -rf build

.PHONY: builddir
builddir:
	mkdir -p build

.PHONY: build-backend
build-backend: builddir
	cd backend && CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build --tags "fts5"  -o ../build/arendt ./

install:
	cd frontend && npm install

build-frontend: builddir
	cd frontend && npm run build && cp -r ./build ../build/frontend

build: build-backend build-frontend

docker:
	docker buildx build --platform linux/amd64 -t web .

deploy-staging:
	heroku container:release web --app dry-falls-76518

deploy-prod:
	heroku container:release web --app arendtarchives

run-frontend:
	cd frontend && npm start

run-backend:
	cd backend && go run --tags "fts5" main.go server

run-sync:
	cd backend && go run main.go sync

run-test:
	cd backend && go run main.go test

run-ocr:
	cd backend && go run main.go ocr
