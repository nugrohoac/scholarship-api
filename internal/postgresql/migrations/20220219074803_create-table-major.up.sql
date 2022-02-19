CREATE TABLE major (
    "id" SERIAL PRIMARY KEY,
    "name" varchar(150) not null,
    "created_at" timestamp(3) DEFAULT CURRENT_TIMESTAMP(3)
);