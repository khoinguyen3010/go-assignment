### Overview
- Backend assignment for Manabie project for todo service.
- This repo has resolved issues asked by Manabie:
	- Change Database from SQLLite to Postgres
	- unit test for 'services' layer (todo)
	- unit test for 'storage' layer (todo)
	- split 'services' into 'use case' and 'transport' layer (todo)
- How to run:
	- docker-compose up -d
	- go run main.go
	- Import Postman collection from `docs` to check example

### DB Schema
```sql
-- users definition

CREATE TABLE users (
	id TEXT NOT NULL,
	password TEXT NOT NULL,
	max_todo INTEGER DEFAULT 5 NOT NULL,
	CONSTRAINT users_PK PRIMARY KEY (id)
);

INSERT INTO users (id, password, max_todo) VALUES('firstUser', 'example', 5);

-- tasks definition

CREATE TABLE tasks (
	id TEXT NOT NULL,
	content TEXT NOT NULL,
	user_id TEXT NOT NULL,
    created_date TEXT NOT NULL,
	CONSTRAINT tasks_PK PRIMARY KEY (id),
	CONSTRAINT tasks_FK FOREIGN KEY (user_id) REFERENCES users(id)
);
```

### Sequence diagram
![auth and create tasks request](https://github.com/manabie-com/togo/blob/master/docs/sequence.svg)
