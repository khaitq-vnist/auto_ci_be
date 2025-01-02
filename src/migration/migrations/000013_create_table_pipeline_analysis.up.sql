CREATE TABLE pipeline_analysis (
                                   id SERIAL PRIMARY KEY,
                                    project_id INT NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
                                   pipeline_id INT NOT NULL REFERENCES pipelines(id) ON DELETE CASCADE,
                                   status VARCHAR(50) NOT NULL,
                                   started_at TIMESTAMP,
                                   ended_at TIMESTAMP,
                                   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE TABLE execution_history (
                                   id SERIAL PRIMARY KEY,
                                   third_party_id INT,
                                   pipeline_id INT NOT NULL REFERENCES pipelines(id) ON DELETE CASCADE,
                                   started_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                   ended_at TIMESTAMP,
                                   status VARCHAR(50) NOT NULL,
                                   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);