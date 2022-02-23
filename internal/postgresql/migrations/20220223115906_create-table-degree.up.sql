CREATE TABLE "degree" (
    "id" SERIAL PRIMARY KEY,
    "name" varchar(200) not null,
    "created_at" timestamp(3) DEFAULT CURRENT_TIMESTAMP(3)
);