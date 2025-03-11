-- +goose Up
-- +goose StatementBegin
CREATE TABLE sentences_attachments (
    sentence_id BIGINT NOT NULL REFERENCES sentences(id) ON DELETE CASCADE,
    attachment_id INT NOT NULL REFERENCES attachments(id) ON DELETE CASCADE,
    attached_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (sentence_id, attachment_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE sentences_attachments;
-- +goose StatementEnd