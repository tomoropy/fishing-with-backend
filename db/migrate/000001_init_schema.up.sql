CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar NOT NULL,
  "profile_text" varchar NOT NULL,
  "profile_image" varchar,
  "header_image" varchar,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "photos" (
  "id" SERIAL PRIMARY KEY,
  "user_id" int,
  "image" varchar,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "invitations" (
  "id" SERIAL PRIMARY KEY,
  "user_id" int,
  "title" varchar NOT NULL,
  "place" varchar NOT NULL,
  "people" int NOT NULL,
  "start_time" timestamp NOT NULL,
  "end_time" timestamp NOT NULL,
  "transpost" varchar NOT NULL,
  "comment" varchar,
  "applicant" int,
  "image" varchar,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "rooms" (
  "id" SERIAL PRIMARY KEY,
  "user_id" int
);

CREATE TABLE "entries" (
  "id" SERIAL PRIMARY KEY,
  "user_id" int,
  "room_id" int
);

CREATE TABLE "messages" (
  "id" SERIAL PRIMARY KEY,
  "user_id" int,
  "room_id" int,
  "body" varchar NOT NULL,
  "is_read" boolean NOT NULL,
  "created_at" timestamptz DEFAULT (now())
);

CREATE INDEX ON "users" ("name");

CREATE INDEX ON "invitations" ("user_id");

CREATE INDEX ON "invitations" ("place");

CREATE INDEX ON "messages" ("user_id");

ALTER TABLE "photos" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "invitations" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "rooms" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "entries" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "entries" ADD FOREIGN KEY ("room_id") REFERENCES "rooms" ("id");

ALTER TABLE "messages" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "messages" ADD FOREIGN KEY ("room_id") REFERENCES "rooms" ("id");
