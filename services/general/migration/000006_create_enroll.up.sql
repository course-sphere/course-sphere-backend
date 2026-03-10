CREATE TABLE IF NOT EXISTS general.enrolls(
    course_id uuid NOT NULL REFERENCES general.courses(id),
    student_id uuid NOT NULL,

    created_at timestamptz NOT NULL DEFAULT now()
);
