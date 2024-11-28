-- Insert into pipeline_templates
INSERT INTO pipeline_templates (name, build_tool, description)
VALUES ('Maven Pipeline', 'maven', 'Pipeline for Maven projects with build, test, and quality scan stages.');

-- Retrieve the ID of the inserted pipeline template
SELECT currval(pg_get_serial_sequence('pipeline_templates', 'id')) AS pipeline_id;

-- Assuming the pipeline_id from the above query is 1
-- Insert into stage_templates
INSERT INTO stage_templates (name, type, docker_image, docker_image_tag)
VALUES
    ('Build (Skip Test)', 'build', 'maven', '3.9.9'),
    ('Test', 'test', 'maven', '3.9.9'),
    ('Scan Quality (Sonar)', 'quality_scan', 'sonarsource/sonar-scanner-cli', 'latest');

-- Retrieve the IDs of the inserted stage templates
SELECT currval(pg_get_serial_sequence('stage_templates', 'id')) AS stage_id;

-- Assuming the stage IDs are 1, 2, and 3 respectively
-- Insert into pipeline_stage_templates
INSERT INTO pipeline_stage_templates (template_id, stage_id)
VALUES
    (1, 1),
    (1, 2),
    (1, 3);

-- Insert into commands_templates for each stage
-- Build stage: Skip tests
INSERT INTO commands_templates (command, stage_id)
VALUES ('mvn clean install -DskipTests', 1);

-- Test stage: Run tests
INSERT INTO commands_templates (command, stage_id)
VALUES ('mvn test', 2);

-- Quality scan stage: Run SonarQube scanner
INSERT INTO commands_templates (command, stage_id)
VALUES ('sonar-scanner', 3);

-- Insert into variable_templates (optional, if stages have variables)
-- Example variables for the Build stage
INSERT INTO variable_templates (stage_id, key, value)
VALUES
    (1, 'SKIP_TESTS', 'true'),
    (1, 'MAVEN_OPTS', '-Xmx1024m');

-- Example variables for the Quality Scan stage
INSERT INTO variable_templates (stage_id, key, value)
VALUES
    (3, 'SONAR_HOST_URL', 'http://sonarqube:9000'),
    (3, 'SONAR_PROJECT_KEY', 'maven_project');
