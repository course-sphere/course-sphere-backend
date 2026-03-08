CREATE TYPE general.material_kind AS ENUM (
    'text',
    'file',
    'video',
    'quiz',
    'assignment'
);

CREATE TABLE IF NOT EXISTS general.materials(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    course_id uuid NOT NULL REFERENCES general.courses(id) ON DELETE CASCADE,
    position numeric NOT NULL,
    kind general.material_kind NOT NULL,
    lesson text NOT NULL,
    title text NOT NULL,
    content text,
    required_score int,
    required_peers int,
    is_required bool NOT NULL,
    
    UNIQUE (course_id, position)
);
