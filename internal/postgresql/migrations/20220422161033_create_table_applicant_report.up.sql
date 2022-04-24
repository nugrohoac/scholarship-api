create table applicant_report (
    "id" SERIAL PRIMARY KEY,
    "applicant_id" bigint not null,
    "report" json default null,
    "created_at" timestamp(3) DEFAULT CURRENT_TIMESTAMP(3)
);