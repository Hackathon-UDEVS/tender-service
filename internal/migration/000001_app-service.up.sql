CREATE TABLE tender (
    id SERIAL PRIMARY KEY,
    client_id INT NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    attachment VARCHAR(250) NOT NULL,
    deadline TIMESTAMP NOT NULL CHECK (deadline > (CURRENT_TIMESTAMP AT TIME ZONE 'UTC' + INTERVAL '5 hours')),
    budget NUMERIC(12, 2) NOT NULL,
    status VARCHAR(50) NOT NULL CHECK (status IN ('open', 'closed', 'cancelled', 'awarded'))
);

CREATE TABLE bid (
    id SERIAL PRIMARY KEY,
    tender_id INT NOT NULL REFERENCES tender(id) ON DELETE CASCADE,
    contractor_id INT NOT NULL,
    price NUMERIC(12, 2) NOT NULL,
    delivery_time INTERVAL NOT NULL,
    comments TEXT,
    status VARCHAR(50) NOT NULL CHECK (status IN ('pending', 'accepted', 'rejected'))
);

CREATE TABLE notification (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    message TEXT NOT NULL,
    relation_id INT NOT NULL,
    type VARCHAR(50) NOT NULL CHECK (type IN ('tender', 'bid')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);
