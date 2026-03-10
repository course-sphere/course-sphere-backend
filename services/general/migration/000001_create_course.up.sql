CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE SCHEMA IF NOT EXISTS general;

CREATE TABLE IF NOT EXISTS general.categories(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    name text UNIQUE NOT NULL
);

CREATE TYPE general.level AS ENUM (
    'beginner',
    'intermediate',
    'advanced'
);

CREATE TYPE general.status AS ENUM (
    'draft',
    'need-review',
    'ai-approved',
    'approved',
    'removed'
);

CREATE TABLE IF NOT EXISTS general.courses(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    instructor_id uuid NOT NULL,
    title text UNIQUE NOT NULL,
    subtitle text,
    description text NOT NULL,
    level general.level NOT NULL,
    thumbnail_url text,
    promo_video_url text,
    price real NOT NULL,
    requirements text,
    learning_objectives text NOT NULL,
    target_audiences text,
    status general.status NOT NULL DEFAULT 'draft'::general.status
);

CREATE TABLE IF NOT EXISTS general.course_categories(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    course_id uuid NOT NULL REFERENCES general.courses(id),
    category_id uuid NOT NULL REFERENCES general.categories(id)
);

CREATE TABLE IF NOT EXISTS general.course_prerequisites(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    course_id uuid NOT NULL REFERENCES general.courses(id),
    other_id uuid NOT NULL REFERENCES general.courses(id)
);
