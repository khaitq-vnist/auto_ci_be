ALTER TABLE execution_history
ADD COLUMN third_party_project VARCHAR(255) NULL;
ALTER TABLE execution_history
ADD COLUMN project_id int NULL;
