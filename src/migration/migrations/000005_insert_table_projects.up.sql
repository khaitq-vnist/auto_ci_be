-- Table for storing project information from repository response
CREATE TABLE projects (
                          id BIGINT PRIMARY KEY, -- Repository ID from the response, e.g., 870169906
                          name VARCHAR(255) NOT NULL, -- Name of the project, e.g., "demo_ci_cd"
                          full_name VARCHAR(255) NOT NULL, -- Full name including username, e.g., "khaitq-vnist/demo_ci_cd"
                          private BOOLEAN DEFAULT FALSE, -- Indicates if the project is private or public
                          owner_id INT NOT NULL, -- ID of the owner in your system, references user or organization
                          html_url VARCHAR(255) NOT NULL, -- URL to the project on the repository, e.g., "
                          providerCode VARCHAR(50) NOT NULL, -- Code for the repository provider, e.g., "github", "gitlab"
                          providerUsername VARCHAR(255) NOT NULL, -- Username of the repository provider, e.g., "khaitq-vnist"
                          description TEXT, -- Optional project description
                          language VARCHAR(50), -- Optional primary language used in the project
                          FOREIGN KEY (owner_id) REFERENCES users(id) ON DELETE CASCADE, -- Reference to a user or organization table
                          created_at TIMESTAMP NOT NULL, -- Created timestamp from the repository
                          updated_at TIMESTAMP NOT NULL, -- Last updated timestamp from the repository
);


