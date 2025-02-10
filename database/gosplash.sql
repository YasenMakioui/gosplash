CREATE TABLE "users" (
  "id" uuid PRIMARY KEY,
  "username" varchar(255) UNIQUE NOT NULL,
  "email" varchar(255) UNIQUE NOT NULL,
  "password_hash" varchar(255) NOT NULL,
  "role" varchar(50) DEFAULT 'user',
  "created_at" timestamp NOT NULL
);

CREATE TABLE "files" (
  "id" uuid PRIMARY KEY,
  "uploader_id" uuid NOT NULL,
  "file_name" varchar(255) NOT NULL,
  "file_size" bigint NOT NULL,
  "storage_path" varchar(512) NOT NULL,
  "expires_at" timestamp NOT NULL,
  "max_downloads" int NOT NULL,
  "downloads" int DEFAULT 0,
  "encryption_key" text NOT NULL,
  "created_at" timestamp NOT NULL
);

CREATE TABLE "file_shares" (
  "id" uuid PRIMARY KEY,
  "file_id" uuid NOT NULL,
  "sender_id" uuid NOT NULL,
  "recipient_id" uuid,
  "public_url" varchar(512),
  "created_at" timestamp NOT NULL
);

CREATE TABLE "secrets" (
  "id" uuid PRIMARY KEY,
  "owner_id" uuid NOT NULL,
  "secret_name" varchar(255) NOT NULL,
  "secret_value" text NOT NULL,
  "expires_at" timestamp,
  "created_at" timestamp NOT NULL
);

CREATE TABLE "secret_shares" (
  "id" uuid PRIMARY KEY,
  "secret_id" uuid NOT NULL,
  "sender_id" uuid NOT NULL,
  "recipient_id" uuid,
  "public_url" varchar(512),
  "expires_at" timestamp,
  "created_at" timestamp
);

ALTER TABLE "files" ADD FOREIGN KEY ("uploader_id") REFERENCES "users" ("id");

ALTER TABLE "file_shares" ADD FOREIGN KEY ("file_id") REFERENCES "files" ("id");

ALTER TABLE "file_shares" ADD FOREIGN KEY ("sender_id") REFERENCES "users" ("id");

ALTER TABLE "file_shares" ADD FOREIGN KEY ("recipient_id") REFERENCES "users" ("id");

ALTER TABLE "secrets" ADD FOREIGN KEY ("owner_id") REFERENCES "users" ("id");

ALTER TABLE "secret_shares" ADD FOREIGN KEY ("secret_id") REFERENCES "secrets" ("id");

ALTER TABLE "secret_shares" ADD FOREIGN KEY ("sender_id") REFERENCES "users" ("id");

ALTER TABLE "secret_shares" ADD FOREIGN KEY ("recipient_id") REFERENCES "users" ("id");
