-- +goose Up
-- +goose StatementBegin
CREATE TABLE "indoor_readings" (
    "ID" INTEGER PRIMARY KEY AUTOINCREMENT,
    "temperature" REAL NOT NULL,
    "humidity" REAL NOT NULL,
    "timestamp" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS indoor_readings;
-- +goose StatementEnd
