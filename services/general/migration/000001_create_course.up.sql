CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE SCHEMA IF NOT EXISTS course;

CREATE TABLE IF NOT EXISTS course.categories(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name text UNIQUE NOT NULL
);

CREATE TYPE course.level AS ENUM (
    'beginner',
    'intermediate',
    'advanced'
);

CREATE TABLE IF NOT EXISTS course.courses(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    title text UNIQUE NOT NULL,
    subtitle text,
    description text NOT NULL,
    level course.level NOT NULL,
    thumbnail_url text,
    promo_video_url text,
    price real NOT NULL,
    requirements text NOT NULL,
    learning_objectives text NOT NULL,
    target_audiences text NOT NULL
);

CREATE TABLE IF NOT EXISTS course.course_categories(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    course_id uuid NOT NULL REFERENCES course.courses(id),
    category_id uuid NOT NULL REFERENCES course.categories(id)
);

CREATE TABLE IF NOT EXISTS course.course_prerequisites(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    course_id uuid NOT NULL REFERENCES course.courses(id),
    other_id uuid NOT NULL REFERENCES course.courses(id)
);
