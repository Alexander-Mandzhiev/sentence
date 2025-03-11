-- +goose Up
-- +goose StatementBegin
CREATE TABLE sentences (
                           id BIGSERIAL PRIMARY KEY,
                           name VARCHAR(255) NOT NULL,
                           phone VARCHAR(15),
                           implementor_id INT REFERENCES implementors(id) ON DELETE SET NULL,
                           department_id INT REFERENCES departments(id) ON DELETE SET NULL,
                           created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                           due_date DATE,
                           encouragement TEXT,
                           problem TEXT,
                           solution TEXT,
                           status_id INT NOT NULL REFERENCES statuses(id) ON DELETE SET NULL,
                           direction_id INT NOT NULL REFERENCES directions(id) ON DELETE SET NULL,
                           priority_id INT NOT NULL REFERENCES priorities(id) ON DELETE SET NULL, -- Внешний ключ на таблицу priorities
                           is_completed BOOLEAN DEFAULT FALSE,
                           created_by INT REFERENCES users(id) ON DELETE SET NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE sentences;
-- +goose StatementEnd