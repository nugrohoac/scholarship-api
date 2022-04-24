CREATE TABLE "user_scholarship" (
    "id" SERIAL PRIMARY KEY,
    "scholarship_id" bigint not null,
    "user_id" bigint not null,
    "essay" text default '',
    "recommendation_letter" json default null,
    "rating" int not null default 0,
    "status" smallint DEFAULT 0,
    "created_at" timestamp(3) DEFAULT CURRENT_TIMESTAMP(3)
);