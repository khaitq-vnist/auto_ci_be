-- Table: templates
CREATE TABLE pipeline_templates (
                           id SERIAL PRIMARY KEY,
                           name VARCHAR(255) NOT NULL,
                           build_tool VARCHAR(100) NOT NULL, -- e.g., Maven, Gradle, Node.js
                           description TEXT,
                           created_at TIMESTAMP DEFAULT NOW(),
                           updated_at TIMESTAMP DEFAULT NOW()
);

-- Table: stages
CREATE TABLE stage_templates (
                        id SERIAL PRIMARY KEY,
                        name VARCHAR(255) NOT NULL, -- e.g., Build, Test, Quality Gate
                        type VARCHAR(50) NOT NULL, -- e.g., build, test, quality_gate
                        docker_image VARCHAR(255), -- e.g., maven, sonarqube
                        docker_image_tag VARCHAR(50) -- e.g., 3.9.9, latest
                        created_at TIMESTAMP DEFAULT NOW(),
                        updated_at TIMESTAMP DEFAULT NOW()
);

-- Table: stage_templates
CREATE TABLE pipeline_stage_templates (
                                id SERIAL PRIMARY KEY,
                                template_id INT NOT NULL,
                                stage_id INT NOT NULL,
                                FOREIGN KEY (template_id) REFERENCES pipeline_templates (id) ON DELETE CASCADE,
                                FOREIGN KEY (stage_id) REFERENCES stages (id) ON DELETE CASCADE,
                                created_at TIMESTAMP DEFAULT NOW(),
                                updated_at TIMESTAMP DEFAULT NOW()


);

-- Table: commands
CREATE TABLE commands_templates (
                          id SERIAL PRIMARY KEY,
                          command TEXT NOT NULL, -- e.g., mvn clean install
                          created_at TIMESTAMP DEFAULT NOW(),
                          updated_at TIMESTAMP DEFAULT NOW()

);

-- Table: stage_commands
CREATE TABLE stage_command_templates (
                               id SERIAL PRIMARY KEY,
                               stage_id INT NOT NULL,
                               command_id INT NOT NULL,
                               FOREIGN KEY (stage_id) REFERENCES stages (id) ON DELETE CASCADE,
                               FOREIGN KEY (command_id) REFERENCES commands (id) ON DELETE CASCADE,
                               created_at TIMESTAMP DEFAULT NOW(),
                               updated_at TIMESTAMP DEFAULT NOW()
);

-- Table: variable_templates
CREATE TABLE variable_templates (
                           id SERIAL PRIMARY KEY,
                           key VARCHAR(255) NOT NULL,
                           value TEXT NOT NULL,
                           created_at TIMESTAMP DEFAULT NOW(),
                           updated_at TIMESTAMP DEFAULT NOW()
);

-- Table: stage_variables
CREATE TABLE stage_variable_templates (
                                id SERIAL PRIMARY KEY,
                                stage_id INT NOT NULL,
                                variable_id INT NOT NULL,
                                FOREIGN KEY (stage_id) REFERENCES stages (id) ON DELETE CASCADE,
                                FOREIGN KEY (variable_id) REFERENCES variables (id) ON DELETE CASCADE,
                                UNIQUE (stage_id, variable_id),
                                created_at TIMESTAMP DEFAULT NOW(),
                                updated_at TIMESTAMP DEFAULT NOW()
);

-- Table: pipeline_setting_templates
CREATE TABLE pipeline_setting_templates (
                                id SERIAL PRIMARY KEY,
                                pipeline_id INT NOT NULL,
                                status VARCHAR(50) NOT NULL, -- e.g., active, inactive
                                FOREIGN KEY (pipeline_id) REFERENCES pipeline_templates (id) ON DELETE CASCADE,
                                created_at TIMESTAMP DEFAULT NOW(),
                                updated_at TIMESTAMP DEFAULT NOW()
);

-- Table: pipeline_trigger_templates
CREATE TABLE pipeline_trigger_templates (
                                id SERIAL PRIMARY KEY,
                                name VARCHAR(255) NOT NULL,
                                trigger_type VARCHAR(50) NOT NULL, -- e.g., manual, webhook
                                created_at TIMESTAMP DEFAULT NOW(),
                                updated_at TIMESTAMP DEFAULT NOW()
);

-- Table: pipeline_code_scope_templates
CREATE TABLE pipeline_code_scope_templates (
                                id SERIAL PRIMARY KEY,
                                code_scope VARCHAR(50) NOT NULL, -- e.g., branch, tag
                                created_at TIMESTAMP DEFAULT NOW(),
                                updated_at TIMESTAMP DEFAULT NOW()
);