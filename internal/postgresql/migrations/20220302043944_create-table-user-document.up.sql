CREATE TABLE "user_document" (
     "id" SERIAL PRIMARY KEY,
     "user_id" bigint not null,
     "document" json DEFAULT NULL,
     "created_at" timestamp(3) DEFAULT CURRENT_TIMESTAMP(3)
);