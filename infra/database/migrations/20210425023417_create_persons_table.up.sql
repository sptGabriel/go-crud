BEGIN;
CREATE TABLE "persons" (
	"id"             varchar(255) primary KEY,
	"first_name"     varchar(255) NOT NULL,
	"last_name"      varchar(255) NOT NULL,
	"email"          varchar(255) NOT NULL,
	"password" 			 varchar(255) NOT NULL,
	"created_at" 		 timestamptz NOT NULL DEFAULT (now()),
	"updated_at" 		 timestamptz NOT NULL DEFAULT (now())
);
COMMIT;