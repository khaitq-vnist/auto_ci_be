-- Create the User table
CREATE TABLE "users" (
                         id serial PRIMARY KEY,
                            name VARCHAR(50) NOT NULL,
                         email VARCHAR(100) UNIQUE NOT NULL,
                         password text NOT NULL,
                         created_at TIMESTAMP DEFAULT NOW(),
                         updated_at TIMESTAMP DEFAULT NOW()
);

-- Create the Provider table
CREATE TABLE providers (
                           id serial PRIMARY KEY,
                           name VARCHAR(50) NOT NULL,  -- Example values: 'GitHub', 'GitLab'
                           code VARCHAR(50) UNIQUE NOT NULL,     -- Provider code for API requests
                           api_url VARCHAR(255) NOT NULL,     -- API base URL for provider
                           auth_url VARCHAR(255) NOT NULL,    -- OAuth authentication URL
                           created_at TIMESTAMP DEFAULT NOW(),
                           updated_at TIMESTAMP DEFAULT NOW()
);

-- Create the Integration table
CREATE TABLE integrations (
                              id serial PRIMARY KEY,
                              name VARCHAR(50) NOT NULL,  -- Example values: 'GitHub', 'GitLab'
                              user_id bigint NOT NULL REFERENCES "users"(id) ON DELETE CASCADE,
                              provider_id bigint NOT NULL REFERENCES providers(id) ON DELETE CASCADE,
                              provider_name VARCHAR(50) NOT NULL,
                              access_token TEXT NOT NULL,        -- Store encrypted token here
                              provider_username VARCHAR(50) NOT NULL,  -- Version control username
                              created_at TIMESTAMP DEFAULT NOW(),
                              updated_at TIMESTAMP DEFAULT NOW(),
                              UNIQUE (user_id, provider_id, access_token)      -- Ensure a user cannot duplicate provider integrations
);