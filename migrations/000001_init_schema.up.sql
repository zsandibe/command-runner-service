CREATE TABLE IF NOT EXISTS command (
    id SERIAL PRIMARY KEY,
    script VARCHAR(255) NOT NULL,
    description VARCHAR NOT NULL
);

CREATE TABLE IF NOT EXISTS command_output(
    id SERIAL PRIMARY KEY,
    command_id INT NOT NULL,
    output VARCHAR,
    time TIMESTAMP NOT NULL,
    FOREIGN KEY (command_id) REFERENCES command (id)
);

CREATE INDEX idx_command_id ON command(id);
CREATE INDEX idx_command_output_id ON command_output(id);