-- +goose Up
CREATE TABLE habits (
    habit_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(user_id),
    name VARCHAR(20),
    description VARCHAR(50),
    type VARCHAR(5),
    created_at TIMESTAMP
);

CREATE TABLE habits_completed (
    id UUID PRIMARY KEY,
    habit_id UUID NOT NULL REFERENCES habits(habit_id),
    completed_at TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS habits_completed;
DROP TABLE IF EXISTS habits;