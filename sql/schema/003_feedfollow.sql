-- +goose Up
CREATE TABLE feed_follows (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  user_id UUID,
  feed_id UUID,
  FOREIGN KEY (feed_id) REFERENCES feeds(id),
  FOREIGN KEY (user_id) REFERENCES feeds(user_id),
  CONSTRAINT user_feed UNIQUE (user_id, feed_id)
);

-- +goose Down
DROP TABLE feed_follows;
