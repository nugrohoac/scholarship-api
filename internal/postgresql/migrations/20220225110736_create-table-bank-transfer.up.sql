CREATE TABLE "bank_transfer" (
    "id" SERIAL PRIMARY KEY,
    "name" varchar(150) not null,
    "account_name" varchar(150) not null,
    "account_no" varchar(150) not null,
    "image" json DEFAULT NULL,
    "created_at" timestamp(3) DEFAULT CURRENT_TIMESTAMP(3)
);