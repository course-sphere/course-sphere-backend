CREATE TABLE IF NOT EXISTS general.attempt_details(
    attempt_id uuid NOT NULL REFERENCES general.attempts(id),
    question_id uuid NOT NULL REFERENCES general.questions(id),
    answer text NOT NULL,

    PRIMARY KEY (attempt_id, question_id)
);
