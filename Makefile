.PHONY: app backend frontend

app:
	docker-compose up --build backend frontend

backend:
	docker-compose run backend go build

frontend:
	docker-compose run frontend bash -c "npm install && npm run build"

backend/xml_update/xml_update: backend/xml_update/main.go
	docker-compose run backend bash -c "cd xml_update && go build"