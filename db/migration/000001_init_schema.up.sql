CREATE TABLE "sbus" (
    "id" bigserial PRIMARY KEY,
    "sbu_name" varchar,
    "sbu_head_user_id" bigint DEFAULT 1,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "users" (
    "id" bigserial PRIMARY KEY,
    "firstname" varchar NOT NULL,
    "lastname" varchar NOT NULL,
    "email" varchar NOT NULL,
    "password" varchar NOT NULL,
    "user_role_id" bigint NULL DEFAULT 1,
    "sbu_id" bigint NULL DEFAULT 1,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "user_roles" (
    "id" bigserial PRIMARY KEY,
    "role_name" varchar NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "clients" (
    "id" bigserial PRIMARY KEY,
    "client_name" varchar NOT NULL,
    "status" varchar NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "projects" (
    "id" bigserial PRIMARY KEY,
    "project_name" varchar NOT NULL,
    "client_id" bigint NOT NULL DEFAULT 1,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "project_teams" (
    "id" bigserial PRIMARY KEY,
    "project_id" bigint NOT NULL,
    "user_id" bigint NOT NULL DEFAULT 1,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "tasks" (
    "id" bigserial PRIMARY KEY,
    "task_title" varchar NOT NULL,
    "progress" bigint NOT NULL,
    "project_id" bigint NOT NULL DEFAULT 1,
    "assigned_user_id" bigint NOT NULL DEFAULT 1,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "sbus" ADD FOREIGN KEY ("sbu_head_user_id") REFERENCES "users" ("id");

ALTER TABLE "projects" ADD FOREIGN KEY ("client_id") REFERENCES "clients" ("id");

ALTER TABLE "project_teams" ADD FOREIGN KEY ("project_id") REFERENCES "projects" ("id");

ALTER TABLE "project_teams" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "tasks" ADD FOREIGN KEY ("project_id") REFERENCES "projects" ("id");

ALTER TABLE "tasks" ADD FOREIGN KEY ("assigned_user_id") REFERENCES "users" ("id");

ALTER TABLE "users" ADD FOREIGN KEY ("user_role_id") REFERENCES "user_roles" ("id");

ALTER TABLE "users" ADD FOREIGN KEY ("sbu_id") REFERENCES "sbus" ("id");
