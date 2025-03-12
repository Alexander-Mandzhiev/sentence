-- +goose Up
-- +goose StatementBegin
CREATE TABLE sentences (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    department VARCHAR(255),
    phone VARCHAR(15),
    problem TEXT,
    solution TEXT,
    encouragement INT,
    is_completed BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    implementor_id INT REFERENCES implementors(id) ON DELETE SET NULL,
    status_id INT NOT NULL REFERENCES statuses(id) ON DELETE SET NULL,
    direction_id INT NOT NULL REFERENCES directions(id) ON DELETE SET NULL,
    priority_id INT NOT NULL REFERENCES priorities(id) ON DELETE SET NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE sentences;
-- +goose StatementEnd