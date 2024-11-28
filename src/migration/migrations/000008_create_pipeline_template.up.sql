
CREATE TABLE pipeline_templates (
                           id SERIAL PRIMARY KEY,
                           name VARCHAR(255) NOT NULL,
                           build_tool VARCHAR(100) NOT NULL,
                           description TEXT,
                           created_at TIMESTAMP DEFAULT NOW(),
                           updated_at TIMESTAMP DEFAULT NOW()
);


CREATE TABLE stage_templates (
                        id SERIAL PRIMARY KEY,
                        name VARCHAR(255) NOT NULL,
                        type VARCHAR(50) NOT NULL,
                        docker_image VARCHAR(255),
                        docker_image_tag VARCHAR(50),
                        created_at TIMESTAMP DEFAULT NOW(),
                        updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE pipeline_stage_templates (
                                id SERIAL PRIMARY KEY,
                                template_id INT NOT NULL,
                                stage_id INT NOT NULL,
                                FOREIGN KEY (template_id) REFERENCES pipeline_templates (id) ON DELETE CASCADE,
                                FOREIGN KEY (stage_id) REFERENCES stage_templates (id) ON DELETE CASCADE,
                                created_at TIMESTAMP DEFAULT NOW(),
                                updated_at TIMESTAMP DEFAULT NOW()


);


CREATE TABLE commands_templates (
                          id SERIAL PRIMARY KEY,
                          command TEXT NOT NULL,
                            stage_id INT NOT NULL,
                            FOREIGN KEY (stage_id) REFERENCES stage_templates (id) ON DELETE CASCADE,
                          created_at TIMESTAMP DEFAULT NOW(),
                          updated_at TIMESTAMP DEFAULT NOW()

);



CREATE TABLE variable_templates (
                           id SERIAL PRIMARY KEY,
                            stage_id INT NOT NULL,
                           key VARCHAR(255) NOT NULL,
                           value TEXT NOT NULL,
                            FOREIGN KEY (stage_id) REFERENCES stage_templates (id) ON DELETE CASCADE,
                           created_at TIMESTAMP DEFAULT NOW(),
                           updated_at TIMESTAMP DEFAULT NOW()
);



CREATE TABLE pipeline_setting_templates (
                                id SERIAL PRIMARY KEY,
                                pipeline_id INT NOT NULL,
                                status VARCHAR(50) NOT NULL, -- e.g., active, inactive
                                FOREIGN KEY (pipeline_id) REFERENCES pipeline_templates (id) ON DELETE CASCADE,
                                created_at TIMESTAMP DEFAULT NOW(),
                                updated_at TIMESTAMP DEFAULT NOW()
);


CREATE TABLE pipeline_trigger_templates (
                                id SERIAL PRIMARY KEY,
                                name VARCHAR(255) NOT NULL,
                                trigger_type VARCHAR(50) NOT NULL,
                                created_at TIMESTAMP DEFAULT NOW(),
                                updated_at TIMESTAMP DEFAULT NOW()
);

-- Table: pipeline_code_scope_templates
CREATE TABLE pipeline_code_scope_templates (
                                id SERIAL PRIMARY KEY,
                                code_scope VARCHAR(50) NOT NULL,
                                created_at TIMESTAMP DEFAULT NOW(),
                                updated_at TIMESTAMP DEFAULT NOW()
);