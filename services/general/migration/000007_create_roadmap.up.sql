CREATE TABLE IF NOT EXISTS general.roadmaps(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    author_id uuid NOT NULL,
    position double precision NOT NULL,
    title text NOT NULL,
    description text NOT NULL
);

CREATE TABLE IF NOT EXISTS general.roadmap_courses(
    roadmap_id uuid NOT NULL REFERENCES general.roadmaps(id),
    course_id uuid NOT NULL REFERENCES general.courses(id)
);

CREATE TABLE IF NOT EXISTS general.student_roadmaps(
    student_id uuid NOT NULL,
    roadmap_id uuid NOT NULL REFERENCES general.roadmaps(id)
);
