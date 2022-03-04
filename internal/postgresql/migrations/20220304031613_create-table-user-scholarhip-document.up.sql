CREATE TABLE "user_scholarship_document" (
    "id" SERIAL PRIMARY KEY,
    "user_scholarship_id" bigint not null,
    "document" json DEFAULT NULL,
    "created_at" timestamp(3) DEFAULT CURRENT_TIMESTAMP(3)
);