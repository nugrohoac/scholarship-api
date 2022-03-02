CREATE TABLE "user_school" (
     "id" SERIAL PRIMARY KEY,
     "user_id" bigint not null,
     "school_id" bigint not null,
     "degree_id" bigint not null,
     "major_id" bigint not null,
     "enrollment_date" timestamp(3) NULL,
     "graduation_date" timestamp(3) NOT NULL,
     "created_at" timestamp(3) DEFAULT CURRENT_TIMESTAMP(3)
);