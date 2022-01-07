CREATE TABLE bank (
    "id" SERIAL PRIMARY KEY,
    "name" varchar(100) NOT NULL,
    "code" varchar(30) NOT NULL,
    "created_at" timestamp(3) DEFAULT CURRENT_TIMESTAMP(3),
    "updated_at" timestamp(3) DEFAULT NULL
);