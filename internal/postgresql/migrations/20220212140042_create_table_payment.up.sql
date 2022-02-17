CREATE TABLE payment (
    "id" SERIAL PRIMARY KEY,
    "scholarship_id" bigint not null,
    "deadline" timestamp(3) not null,
    "transfer_date" timestamp(3) default null,
    "bank_account_name" varchar(100) NOT NULL default '',
    "image" json DEFAULT NULL,
    "created_at" timestamp(3) DEFAULT CURRENT_TIMESTAMP(3),
    "updated_at" timestamp(3) NULL
);