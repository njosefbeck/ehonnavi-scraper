CREATE TABLE "public"."book" (
    "id" serial,
    "url" varchar(255) NOT NULL,
    "title" varchar(255) NOT NULL,
    "age" varchar(10),
    "created_at" timestamp,
    "updated_at" timestamp,
    "deleted_at" timestamp,
    PRIMARY KEY ("id"),
    UNIQUE ("url"),
    FOREIGN KEY ("age") REFERENCES "public"."age"("id")
);