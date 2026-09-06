CREATE TABLE IF NOT EXISTS public.users (
    id          SERIAL PRIMARY KEY,
    email       VARCHAR(255) UNIQUE NOT NULL,
    first_name  VARCHAR(255),
    last_name   VARCHAR(255),
    password    VARCHAR(60),
    user_active INTEGER DEFAULT 0,
    created_at  TIMESTAMP,
    updated_at  TIMESTAMP
);

-- Usuario de prueba. Password = "verysecret" (hash bcrypt).
INSERT INTO public.users (email, first_name, last_name, password, user_active, created_at, updated_at)
VALUES (
    'admin@example.com',
    'Admin',
    'User',
    '$2a$12$1zGLuYDDNvATh4RA4avbKuheAMpb1svexSzrQm7up.bkFdChFa06e',
    1,
    now(),
    now()
)
ON CONFLICT (email) DO NOTHING;
