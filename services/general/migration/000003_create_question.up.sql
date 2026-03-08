CREATE TABLE IF NOT EXISTS general.questions(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    material_id uuid NOT NULL REFERENCES general.materials(id) ON DELETE CASCADE,
    position numeric NOT NULL,
    question text NOT NULL,
    
    UNIQUE (material_id, position)
);

CREATE TABLE IF NOT EXISTS general.question_possible_answers(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    question_id uuid NOT NULL REFERENCES general.questions(id) ON DELETE CASCADE,
    answer text NOT NULL
);

CREATE TABLE IF NOT EXISTS general.question_criteria(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    question_id uuid NOT NULL REFERENCES general.questions(id) ON DELETE CASCADE,
    criterion text NOT NULL,
    score int NOT NULL
);
