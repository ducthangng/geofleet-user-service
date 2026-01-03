

CREATE TABLE users (
    -- Use UUID or BIGSERIAL. UUID is safer for distributed systems.
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    fullname VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    address TEXT NOT NULL,
    phone VARCHAR(50) NOT NULL,
    date_created TIMESTAMPTZ NOT NULL DEFAULT now()
);