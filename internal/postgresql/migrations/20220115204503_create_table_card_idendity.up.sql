CREATE TABLE card_identity (
    "id" SERIAL PRIMARY KEY,
    "type" varchar(30) NOT NULL,
    "no" varchar(50) NOT NULL,
    "image" json DEFAULT NULL,
    "user_id" bigint DEFAULT 0,
    "created_at" timestamp(3) DEFAULT CURRENT_TIMESTAMP(3)
);