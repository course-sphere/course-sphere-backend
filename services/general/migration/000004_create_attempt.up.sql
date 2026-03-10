CREATE TABLE IF NOT EXISTS general.attempts(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    material_id uuid NOT NULL REFERENCES general.materials(id) ON DELETE CASCADE,
    student_id uuid NOT NULL,
    score int,
    created_at timestamptz NOT NULL DEFAULT now()
);
