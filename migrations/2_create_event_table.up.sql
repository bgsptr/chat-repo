CREATE TABLE IF NOT EXISTS events(
    id VARCHAR(255) NOT NULL,
    event_name VARCHAR(255) NOT NULL,
    from_date VARCHAR(255) NOT NULL,
    to_date VARCHAR(255) NOT NULL,
    event_location VARCHAR(255) NOT NULL,
    descriptions VARCHAR(255),
    PRIMARY KEY (id)
)

CREATE TABLE IF NOT EXISTS event_person_confirmed(
    id VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    confirmed BOOLEAN NOT NULL,
    PRIMARY KEY (event_id, username)
)

-- CREATE TABLE IF NOT EXISTS event_person_invited(
--     id VARCHAR(255) NOT NULL,
--     person_invited VARCHAR(255) NOT NULL
--     PRIMARY KEY (event_id, person_invited)
-- )