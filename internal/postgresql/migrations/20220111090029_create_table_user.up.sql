CREATE TABLE "user" (
    "id" SERIAL PRIMARY KEY,
    "name" varchar(100) DEFAULT '',
    "type" varchar(30) NOT NULL,
    "email" varchar(100) NOT NULL,
    "phone_no" varchar(20) NOT NULL,
    "photo" json DEFAULT NULL,
    "company_name" varchar(255) DEFAULT '',
    "password" varchar(255) NOT NULL,
    "status" smallint DEFAULT 0,
    "country_id" bigint DEFAULT 0,
    "postal_code" varchar(50) DEFAULT '',
    "address" varchar(255) DEFAULT '',
    "gender" varchar(15) DEFAULT '',
    "ethnic" varchar(30) DEFAULT '',
    "bank_id" bigint DEFAULT 0,
    "bank_account_no" varchar(50) DEFAULT '',
    "bank_account_name" varchar(100) DEFAULT '',
    "created_at" timestamp(3) DEFAULT CURRENT_TIMESTAMP(3),
    "updated_at" timestamp(3) NULL
);