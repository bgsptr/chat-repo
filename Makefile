request:
	createAcc:
curl -X POST http://localhost:3000/api/v1/user -H "Content-Type: application/json" -d '{"username": "wawa", "password": "wawa"}'

	migrate -database "mysql://root:@tcp(127.0.0.2:8111)/chat_app" -path internal/database/mysql/migration up

	migrate -database "mysql://root:@tcp(127.0.0.2:8111)/chat_app" -path internal/database/mysql/migration down