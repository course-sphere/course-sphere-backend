CREATE TYPE "public"."role" AS ENUM('student', 'instructor', 'admin');--> statement-breakpoint
ALTER TABLE "user" ADD COLUMN "role" "role" DEFAULT 'student' NOT NULL;