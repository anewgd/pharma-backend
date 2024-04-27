-- SQL dump generated using DBML (dbml-lang.org)
-- Database: PostgreSQL
-- Generated at: 2024-04-27T07:36:18.623Z

CREATE TABLE "pharmacies" (
  "pharmacy_id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "pharmacy_name" varchar UNIQUE NOT NULL,
  "user_id" uuid UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "pharmacy_branches" (
  "pharmacy_branch_id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "pharmacy_id" uuid NOT NULL,
  "pharmacy_branch_name" varchar UNIQUE NOT NULL,
  "city" varchar NOT NULL,
  "sub_city" varchar NOT NULL,
  "special_location_name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "users" (
  "user_id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "username" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "role" varchar NOT NULL DEFAULT 'ADMIN',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "drugs" (
  "drug_id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "pharmacy_branch_id" uuid NOT NULL,
  "brand_name" varchar UNIQUE NOT NULL,
  "generic_name" varchar NOT NULL,
  "quantity" bigint NOT NULL,
  "expiration_date" timestamptz NOT NULL,
  "manufacturing_date" timestamptz NOT NULL,
  "pharmacist_id" uuid NOT NULL,
  "added_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "pharmacists" (
  "pharmacist_id" uuid PRIMARY KEY DEFAULT (gen_random_uuid()),
  "pharmacy_branch_id" uuid NOT NULL,
  "username" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "role" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "user_sessions" (
  "session_id" uuid PRIMARY KEY NOT NULL,
  "user_id" uuid NOT NULL,
  "refresh_token" varchar NOT NULL,
  "is_blocked" bool NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "pharmacist_sessions" (
  "session_id" uuid PRIMARY KEY NOT NULL,
  "pharmacist_id" uuid NOT NULL,
  "refresh_token" varchar NOT NULL,
  "is_blocked" bool NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "pharmacies" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "pharmacy_branches" ADD FOREIGN KEY ("pharmacy_id") REFERENCES "pharmacies" ("pharmacy_id");

ALTER TABLE "drugs" ADD FOREIGN KEY ("pharmacy_branch_id") REFERENCES "pharmacy_branches" ("pharmacy_branch_id");

ALTER TABLE "drugs" ADD FOREIGN KEY ("pharmacist_id") REFERENCES "pharmacists" ("pharmacist_id");

ALTER TABLE "pharmacists" ADD FOREIGN KEY ("pharmacy_branch_id") REFERENCES "pharmacy_branches" ("pharmacy_branch_id");

ALTER TABLE "user_sessions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "pharmacist_sessions" ADD FOREIGN KEY ("pharmacist_id") REFERENCES "pharmacists" ("pharmacist_id");
