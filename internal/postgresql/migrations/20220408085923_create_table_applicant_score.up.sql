create table applicant_score (
    "id" SERIAL PRIMARY KEY,
    "applicant_id" bigint not null,
    "name" varchar(150) not null,
    "value" int not null,
    "created_at" timestamp(3) DEFAULT CURRENT_TIMESTAMP(3)
);