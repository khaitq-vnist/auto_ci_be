-- Table for storing pipeline configurations
CREATE TABLE pipelines (
                           id SERIAL PRIMARY KEY,
                           name VARCHAR(255) NOT NULL,
                           description TEXT,
                           owner_id INT NOT NULL, -- Reference to user or project table
                           is_active BOOLEAN DEFAULT TRUE,
                            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table for storing stages within a pipeline
CREATE TABLE pipeline_stages (
                                 id SERIAL PRIMARY KEY,
                                 pipeline_id INT NOT NULL REFERENCES pipelines(id) ON DELETE CASCADE,
                                 name VARCHAR(255) NOT NULL,
                                 stage_order INT NOT NULL, -- Order of the stage in the pipeline
                                 type VARCHAR(50) NOT NULL, -- E.g., "build", "test", "deploy"
                                 created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                 updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table for storing actions within each stage
CREATE TABLE stage_actions (
                               id SERIAL PRIMARY KEY,
                               stage_id INT NOT NULL REFERENCES pipeline_stages(id) ON DELETE CASCADE,
                               name VARCHAR(255) NOT NULL,
                               action_order INT NOT NULL, -- Order of the action in the stage
                               command TEXT NOT NULL, -- Command to be executed (e.g., "mvn clean install")
                               docker_image_name VARCHAR(255), -- Docker image name if needed (e.g., "library/maven")
                               docker_image_tag VARCHAR(50), -- Docker image tag/version (e.g., "3.9.9")
                               working_directory VARCHAR(255), -- Directory where command should run
                               shell_type VARCHAR(50) DEFAULT 'bash', -- Shell type (e.g., "bash", "sh")
                               created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                               updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table for storing environment variables for pipelines
CREATE TABLE pipeline_variables (
                                    id SERIAL PRIMARY KEY,
                                    pipeline_id INT NOT NULL REFERENCES pipelines(id) ON DELETE CASCADE,
                                    key VARCHAR(255) NOT NULL,
                                    value TEXT NOT NULL,
                                    is_secret BOOLEAN DEFAULT FALSE, -- True if variable is a secret
                                    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                    UNIQUE(pipeline_id, key)
);

-- Table for storing services (e.g., databases) required by a pipeline
CREATE TABLE pipeline_services (
                                   id SERIAL PRIMARY KEY,
                                   pipeline_id INT NOT NULL REFERENCES pipelines(id) ON DELETE CASCADE,
                                   service_type VARCHAR(50) NOT NULL, -- Type of service (e.g., "MySQL", "MongoDB")
                                   version VARCHAR(50) NOT NULL, -- Version of the service (e.g., "5.7", "3.2.4")
                                   configuration JSONB, -- Additional configuration in JSON format
                                   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Table for storing pipeline execution history
CREATE TABLE pipeline_history (
                                  id SERIAL PRIMARY KEY,
                                  pipeline_id INT NOT NULL REFERENCES pipelines(id) ON DELETE CASCADE,
                                  started_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                  ended_at TIMESTAMP,
                                  status VARCHAR(50) NOT NULL, -- Status of the pipeline run (e.g., "success", "failed", "in_progress")
                                  triggered_by INT NOT NULL, -- ID of the user or automated process that triggered the run
                                  logs TEXT, -- Logs of the pipeline run
                                  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
