-- +goose Up
-- +goose StatementBegin
CREATE TABLE sentences (
    id BIGSERIAL PRIMARY KEY,
    status_id INT NOT NULL REFERENCES statuses(id) ON DELETE SET NULL,
    name VARCHAR(255) NOT NULL,
    phone VARCHAR(15),
    department_id INT REFERENCES departments(id) ON DELETE SET NULL,
    problem TEXT,
    solution TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    direction_id INT NOT NULL REFERENCES directions(id) ON DELETE SET NULL,
    implementor_id INT REFERENCES implementors(id) ON DELETE SET NULL,
    priority_id INT NOT NULL REFERENCES priorities(id) ON DELETE SET NULL,
    encouragement INT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE sentences;
-- +goose StatementEnd