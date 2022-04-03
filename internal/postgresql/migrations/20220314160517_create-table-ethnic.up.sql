CREATE TABLE "ethnic" (
    "id" SERIAL PRIMARY KEY,
    "name" varchar(100) not null default '',
    "created_at" timestamp(3) DEFAULT CURRENT_TIMESTAMP(3)
);