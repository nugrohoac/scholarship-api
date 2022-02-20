CREATE TABLE "school" (
    "id" SERIAL PRIMARY KEY,
    "name" varchar(200) DEFAULT '',
    "type" varchar(100) DEFAULT '',
    "address" text not null default '',
    "status" smallint DEFAULT 1,
    "created_at" timestamp(3) DEFAULT CURRENT_TIMESTAMP(3),
    "updated_at" timestamp(3) NULL
);