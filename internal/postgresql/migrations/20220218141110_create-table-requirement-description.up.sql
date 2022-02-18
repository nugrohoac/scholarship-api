CREATE TABLE requirement_description (
     "id" SERIAL PRIMARY KEY,
     "scholarship_id" bigint not null,
     "description" text DEFAULT '',
     "created_at" timestamp(3) DEFAULT CURRENT_TIMESTAMP(3)
);