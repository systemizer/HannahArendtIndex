clean:
	rm -rf build

.PHONY: builddir
builddir:
	mkdir -p build

.PHONY: build-backend
build-backend: builddir
	cd backend && CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build --tags "fts5"  -o ../build/arendt ./

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

psql:
	docker-compose exec postgres psql -U postgres -d arendt

dump:
	PGPASSWORD=postgres pg_dump -Fc --no-acl --no-owner -h localhost -p 5432 -U postgres arendt > mydb.dump

dump-data:
	PGPASSWORD=postgres pg_dump -h localhost -U postgres --data-only --inserts arendt > dump.sql
	sed -i -e 's/^INSERT INTO public\./INSERT INTO /' dump.sql
