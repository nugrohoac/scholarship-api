CREATE TABLE "requirement" (
    "id" SERIAL PRIMARY KEY,
    "scholarship_id" bigint NOT NULL,
    "type" varchar(255) NOT NULL,
    "name" varchar(255) NOT NULL,
    "value" varchar(255) NOT NULL,
    "created_at" timestamp(3) DEFAULT CURRENT_TIMESTAMP(3),
    "updated_at" timestamp(3) NULL
);