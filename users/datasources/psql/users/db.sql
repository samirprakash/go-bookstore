CREATE TABLE "public"."users" (
    "id" serial NOT NULL,
    "first_name" varchar NOT NULL,
    "last_name" varchar NOT NULL,
    "email" varchar NOT NULL,
    "created" varchar NOT NULL,
    "status" varchar NOT NULL,
    "password" varchar NOT NULL,
    PRIMARY KEY ("id")
);
CREATE UNIQUE INDEX "users_email" ON "public"."users" USING BTREE ("email");