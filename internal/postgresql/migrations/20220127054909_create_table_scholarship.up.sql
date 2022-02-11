CREATE TABLE "scholarship" (
    "id" SERIAL PRIMARY KEY,
    "sponsor_id" bigint NOT NULL,
    "name" varchar(200) DEFAULT '',
    "amount" integer NOT NULL,
    "status" smallint DEFAULT 0,
    "image" json DEFAULT NULL,
    "awardee" integer NOT NULL,
    "current_applicant" integer DEFAULT 0,
    "application_end" timestamp(3) NOT NULL,
    "eligibility_description" text DEFAULT '',
    "subsidy_description" text DEFAULT '',
    "requirement_descriptions" text DEFAULT '',
    "funding_start" timestamp(3) NOT NULL,
    "funding_end" timestamp(3) NOT NULL,
    "created_at" timestamp(3) DEFAULT CURRENT_TIMESTAMP(3),
    "updated_at" timestamp(3) NULL
);