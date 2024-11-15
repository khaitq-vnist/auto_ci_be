-- Table for storing project information from repository response
CREATE TABLE projects (
                          id serial PRIMARY KEY,
                          name VARCHAR(255) NOT NULL,
                          full_name VARCHAR(255) NOT NULL,
                          private BOOLEAN DEFAULT FALSE,
                          owner_id INT NOT NULL,
                          html_url VARCHAR(255) NOT NULL,
                          provider_repo_id bigint NOT NULL,
                          description TEXT,
                          language VARCHAR(50),
                          FOREIGN KEY (owner_id) REFERENCES users(id) ON DELETE CASCADE,
                          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                          updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


