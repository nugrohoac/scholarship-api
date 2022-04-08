create table applicant_eligibility (
    "id" SERIAL PRIMARY KEY,
    "applicant_id" bigint not null,
    "requirement_id" bigint not null,
    "value" boolean not null,
    "created_at" timestamp(3) DEFAULT CURRENT_TIMESTAMP(3)
);