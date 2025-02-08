CREATE TABLE "Users" (
  "user_id" INT PRIMARY KEY,
  "email" VARCHAR,
  "password_hash" VARCHAR,
  "created_at" DATETIME
);

CREATE TABLE "Secrets" (
  "secret_id" INT PRIMARY KEY,
  "user_id" INT,
  "hash" VARCHAR,
  "data" TEXT,
  "expires_at" DATETIME,
  "created_at" DATETIME
);

CREATE TABLE "Files" (
  "file_id" INT PRIMARY KEY,
  "user_id" INT,
  "hash" VARCHAR,
  "file_path" VARCHAR,
  "expires_at" DATETIME,
  "created_at" DATETIME
);

CREATE TABLE "Access_Logs" (
  "log_id" INT PRIMARY KEY,
  "secret_or_file_id" INT,
  "accessed_at" DATETIME,
  "ip_address" VARCHAR
);

CREATE TABLE "Revoked_Tokens" (
  "token_id" INT PRIMARY KEY,
  "jti" VARCHAR UNIQUE,
  "revoked_at" DATETIME
);

ALTER TABLE "Secrets" ADD FOREIGN KEY ("user_id") REFERENCES "Users" ("user_id");

ALTER TABLE "Files" ADD FOREIGN KEY ("user_id") REFERENCES "Users" ("user_id");

ALTER TABLE "Secrets" ADD FOREIGN KEY ("secret_id") REFERENCES "Access_Logs" ("secret_or_file_id");

ALTER TABLE "Files" ADD FOREIGN KEY ("file_id") REFERENCES "Access_Logs" ("secret_or_file_id");
