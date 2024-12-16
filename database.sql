-- tenant table maintaining list of tenants 
SET TIME ZONE 'Asia/Kolkata';

-- Utility function for setting updated timestamp on modification
CREATE OR REPLACE FUNCTION trigger_set_updated_timestamptz()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Tenant table to manage tenants in a multi-tenant environment
DROP TABLE IF EXISTS tenant CASCADE;
CREATE TABLE tenant (
    id SERIAL PRIMARY KEY,
    slug_name VARCHAR(20) NOT NULL UNIQUE,
    description VARCHAR(150) NOT NULL DEFAULT '',
    active BOOLEAN NOT NULL DEFAULT TRUE
);

INSERT INTO tenant (slug_name, description)
VALUES 
    ('try', 'Test tenant for the application'),
    ('main', 'Main tenant for the application'),
    ('app', 'Default tenant for the users in application');

  
-- HTTP methods table to store various HTTP methods
DROP TABLE IF EXISTS http_method CASCADE;
CREATE TABLE http_method (
    id SERIAL PRIMARY KEY,
    name VARCHAR(10) NOT NULL,
    value VARCHAR(10) NOT NULL UNIQUE,
    active BOOLEAN NOT NULL DEFAULT TRUE
);

INSERT INTO http_method (name, value) 
VALUES ('GET', 'GET'), ('POST', 'POST'), ('PUT', 'PUT'), ('PATCH', 'PATCH'), ('DELETE', 'DELETE');

-- Application URLs and associated methods for defining permissions and page hierarchy
DROP TABLE IF EXISTS app_url CASCADE;
CREATE TABLE app_url (
    id SERIAL PRIMARY KEY,
    page_name VARCHAR(30) NOT NULL,
    page_path VARCHAR(50) NOT NULL UNIQUE,
    parent_id INT,
    module_name VARCHAR(30) NOT NULL DEFAULT '',
    active BOOLEAN NOT NULL DEFAULT TRUE,
    FOREIGN KEY (parent_id) REFERENCES app_url (id) ON DELETE CASCADE
);


DROP TABLE IF EXISTS app_url_method CASCADE;
CREATE TABLE app_url_method (
    id SERIAL PRIMARY KEY,
    http_method_id INT NOT NULL,
    app_url_id INT NOT NULL,
    active BOOLEAN NOT NULL DEFAULT TRUE,
    FOREIGN KEY (http_method_id) REFERENCES http_method (id),
    FOREIGN KEY (app_url_id) REFERENCES app_url (id)
);

-- CREATE SCHEMA IF NOT EXISTS app;
-- on landing page when creating tenant, the user will provide email_address, name of the company, first_name and last_name of himself, who is singing up for the company.
-- email_address is mandatory and link will be sent for verification and on that link, the user can update password
-- Person table to manage user records
DROP TABLE IF EXISTS person CASCADE;
CREATE TABLE person (
    id SERIAL PRIMARY KEY,
    email VARCHAR(150) NOT NULL UNIQUE,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(70) NOT NULL,
    password VARCHAR(255),
    email_verified BOOLEAN NOT NULL DEFAULT FALSE,
    active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ DEFAULT current_timestamp(5),
    updated_at TIMESTAMPTZ DEFAULT current_timestamp(5),
    modified_by INT,
    FOREIGN KEY (modified_by) REFERENCES person (id)
);

CREATE TRIGGER set_timestamp BEFORE 
UPDATE ON person FOR EACH ROW EXECUTE PROCEDURE trigger_set_updated_timestamptz();

  
insert into person (email, first_name, last_name) 
values 
  (
    'support@data_abc.com', 'abc', 'xyz'
  );
  
select 
  * 
from 
  person;
  
-- Table for tokens (e.g., access, refresh tokens) for authentication
DROP TABLE IF EXISTS tokens CASCADE;
CREATE TABLE tokens (
    id BIGSERIAL PRIMARY KEY,
    token_type VARCHAR(10) NOT NULL,
    token_value VARCHAR(1000) NOT NULL,
    active BOOLEAN NOT NULL DEFAULT TRUE,
    start_time TIMESTAMPTZ NOT NULL,
    end_time TIMESTAMPTZ NOT NULL CHECK (end_time > start_time),
    created_at TIMESTAMPTZ DEFAULT current_timestamp(5),
    updated_at TIMESTAMPTZ DEFAULT current_timestamp(5)
);

-- Person session table to handle user sessions with device info and session management
DROP TABLE IF EXISTS person_session CASCADE;
CREATE TABLE person_session (
    session_id UUID PRIMARY KEY DEFAULT gen_random_uuid(), -- Unique session ID,
    person_id INT NOT NULL,
    access_token_id BIGINT NOT NULL,
    refresh_token_id BIGINT NOT NULL,
    start_time TIMESTAMPTZ NOT NULL,
    end_time TIMESTAMPTZ,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    device_info VARCHAR(255) DEFAULT '',
    ip_address VARCHAR(45) DEFAULT '',
    user_agent VARCHAR(255) DEFAULT '',
    temp_url_path VARCHAR(300) DEFAULT '',
    login VARCHAR(40) DEFAULT '',
    created_at TIMESTAMPTZ DEFAULT current_timestamp(5),
    updated_at TIMESTAMPTZ DEFAULT current_timestamp(5),

    CONSTRAINT fk_person_session_user FOREIGN KEY (person_id) REFERENCES person (id) ON DELETE CASCADE,
    CONSTRAINT fk_access_token_tokens FOREIGN KEY (access_token_id) REFERENCES tokens (id),
    CONSTRAINT fk_refresh_token_tokens FOREIGN KEY (refresh_token_id) REFERENCES tokens (id)
);

CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TRIGGER set_timestamp BEFORE 
UPDATE ON person_session FOR EACH ROW EXECUTE PROCEDURE trigger_set_updated_timestamptz();
 
select * from person;
select * from person_session;
select * from tokens;


-- -- Ensure only one active session per person
-- CREATE OR REPLACE FUNCTION before_insert_person_session_trigger_function()
-- RETURNS TRIGGER AS $$
-- BEGIN
--     UPDATE person_session SET is_active = false
--     WHERE person_id = NEW.person_id AND is_active = true;
--     RETURN NEW;
-- END;
-- $$ LANGUAGE plpgsql;

-- CREATE TRIGGER before_insert_person_session_trigger
-- BEFORE INSERT ON person_session FOR EACH ROW
-- EXECUTE FUNCTION before_insert_person_session_trigger_function();

ALTER TABLE person_session ADD COLUMN last_activity TIMESTAMPTZ DEFAULT current_timestamp(5);

-- Indexes for faster retrieval
CREATE INDEX idx_person_session_person_lid ON person_session (person_lid);
CREATE INDEX idx_person_session_access_token ON person_session (access_token);
CREATE INDEX idx_person_session_refresh_token ON person_session (refresh_token);
CREATE INDEX idx_person_session_start_time ON person_session (start_time);
CREATE INDEX idx_person_session_end_time ON person_session (end_time);
 
-- delete from app.user_role where user_lid = 3;
-- delete from app.user where id = 3;

-- Role management
DROP TABLE IF EXISTS role CASCADE;
CREATE TABLE role (
    id SERIAL PRIMARY KEY,
    name VARCHAR(20) NOT NULL UNIQUE,
    description VARCHAR(100) DEFAULT '',
    active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ DEFAULT current_timestamp(5),
    updated_at TIMESTAMPTZ DEFAULT current_timestamp(5),
    modified_by INT,
    FOREIGN KEY (modified_by) REFERENCES person (id)
);

CREATE TRIGGER set_timestamp BEFORE 
UPDATE ON role FOR EACH ROW EXECUTE PROCEDURE trigger_set_updated_timestamptz();

insert into role (name, description, modified_by) 
values 
  (
    'super', 'all rights for read and write for all endpoints', 
    1
  );
insert into role (name, description, modified_by) 
values 
  (
    'Not Allocated', 'No rights for read and write for all endpoints, as new user and no role alloted', 
    1
  );
  
select * from role;
-- role table will save only role name and description, to add permitted endpoints with methods for the roles will be maintained in role_permission table and that it will be helpful to get users role and role_permissions. Each user can be assigned multiple roles which is mapped in user_role table
-- Role permissions and person roles
DROP TABLE IF EXISTS role_permission CASCADE;
CREATE TABLE role_permission (
    role_id INT NOT NULL,
    app_url_id INT NOT NULL,
    http_method_id INT NOT NULL,
    active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ DEFAULT current_timestamp(5),
    updated_at TIMESTAMPTZ DEFAULT current_timestamp(5),
    modified_by INT,
    PRIMARY KEY (role_id, app_url_id, http_method_id),
    FOREIGN KEY (role_id) REFERENCES role (id),
    FOREIGN KEY (app_url_id) REFERENCES app_url (id),
    FOREIGN KEY (http_method_id) REFERENCES http_method (id)
);

CREATE OR REPLACE TRIGGER set_timestamp BEFORE 
UPDATE 
  ON role_permission FOR EACH ROW EXECUTE PROCEDURE trigger_set_updated_timestamptz ();

DROP TABLE IF EXISTS person_role CASCADE;
CREATE TABLE person_role (
    person_id INT NOT NULL,
    role_id INT NOT NULL,
    active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ DEFAULT current_timestamp(5),
    updated_at TIMESTAMPTZ DEFAULT current_timestamp(5),
    modified_by INT,
    PRIMARY KEY (person_id, role_id),
    FOREIGN KEY (person_id) REFERENCES person (id),
    FOREIGN KEY (role_id) REFERENCES role (id)
);

CREATE OR REPLACE TRIGGER set_timestamp BEFORE 
UPDATE 
  ON person_role FOR EACH ROW EXECUTE PROCEDURE trigger_set_updated_timestamptz();
  
-- login
-- loading page
-- dashboard page - 
-- connect to remote database
-- tenants
-- users & accounts 

-- organization table - it contains the details related to the organization for which the installation is being made now. It will hold only single record. And mandatory task after first time installation
DROP TABLE IF EXISTS organization CASCADE;
CREATE TABLE organization (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(800) DEFAULT '',
    employee_count INT,
    industry_type VARCHAR(100) DEFAULT '',
    active BOOLEAN NOT NULL DEFAULT TRUE
);

insert into organization(name, description, employee_count, industry_type) values ('abc corp', 'ABC World Corporation', 800, 'Advertising')

---
CREATE OR REPLACE PROCEDURE your_procedure_name(input_json JSONB, OUT response_json JSONB)
LANGUAGE plpgsql
AS $$
DECLARE
    temp_table_name TEXT := 'temp_data'; -- Customize the temporary table name
BEGIN
    -- Step 1: Create a temporary table and parse input JSON
    CREATE TEMP TABLE IF NOT EXISTS temp_data (
        -- Define your temporary table columns based on the JSON structure
        id SERIAL PRIMARY KEY,
        name VARCHAR(255),
        age INT
    );

    -- Insert data from input JSON into the temporary table
    INSERT INTO temp_data (name, age)
    SELECT (json_data->>'name')::VARCHAR(255), (json_data->>'age')::INT
    FROM jsonb_array_elements(input_json) AS json_data;

    -- Step 2: Perform operations on the temporary table
    -- Customize this section based on your specific requirements
    -- Example: Select data from the temporary table
    SELECT * INTO response_json.data
    FROM temp_data;

    -- Step 3: Commit the transaction
    COMMIT;

    -- Step 4: Set status and exception_message
    response_json.status := 'success';
    response_json.exception_message := NULL; -- No exception

EXCEPTION
    WHEN OTHERS THEN
        -- Step 5: Handle exceptions
        -- Rollback the transaction
        ROLLBACK;

        -- Get the exception message
        GET STACKED DIAGNOSTICS response_json.exception_message = MESSAGE_TEXT;

        -- Set status
        response_json.status := 'error';
        response_json.data := NULL; -- No data on error
END;
$$;

---

-- drop table if exists person_session;
-- CREATE TABLE person_session (
--     id bigserial PRIMARY KEY,
--     person_lid INT NOT NULL,
--     session_uuid varchar(100) not null,
--     active BOOLEAN NOT NULL DEFAULT true,
--     device_info VARCHAR(255) not null default '',
--     ip_address VARCHAR(45) not null default '',
--     user_agent VARCHAR(255) not null default '',
-- 	temp_url_path varchar(300) not null default '',
-- 	login varchar(40) not null default '',
-- 	created_at timestamptz not null default current_timestamp (5), 
--   	updated_at timestamptz not null default current_timestamp (5),
	
--     CONSTRAINT fk_person_session_user
--         FOREIGN KEY (person_lid)
--         REFERENCES person(id)
--         ON DELETE CASCADE
-- );


-- CREATE OR REPLACE FUNCTION before_insert_person_session_trigger_function()
-- RETURNS TRIGGER AS $$
-- BEGIN
--     -- Update existing active records for the same person_lid
--     UPDATE person_session
--     SET active = false
--     WHERE person_lid = NEW.person_lid AND active = true;

--     -- Return the NEW row as it is for insertion
--     RETURN NEW;
-- END;
-- $$ LANGUAGE plpgsql;

-- CREATE TRIGGER before_insert_person_session_trigger
-- BEFORE INSERT ON person_session
-- FOR EACH ROW
-- EXECUTE FUNCTION before_insert_person_session_trigger_function();

-- data source settings
-- Data source configurations
DROP TABLE IF EXISTS data_source_type CASCADE;
CREATE TABLE data_source_type (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description VARCHAR(700) DEFAULT '',
    active BOOLEAN NOT NULL DEFAULT TRUE
);

insert into data_source_type(name, description) values 
('Postgres', 'Connects to PostgreSQL database instances or other instances which supports postgres connection'),
('SQL Server', 'Connects to SQL Server (Microsoft SQL Server or Azure SQL) database instances or other instances which supports SQL Server connection'),
('SQLite3', 'Connects to SQLite3 database instances or other instances which supports SQLite3 connection');

select * from data_source_type;

DROP TABLE IF EXISTS connection_mode CASCADE;
CREATE TABLE connection_mode (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description VARCHAR(700) DEFAULT '',
    active BOOLEAN NOT NULL DEFAULT TRUE
);

insert into connection_mode(name, description, active) values 
('Host & Port', 'Connection made using Host Address and Port Number', true),
('Socket', 'Connection made using Socket', false);

select * from connection_mode;

DROP TABLE IF EXISTS data_source CASCADE;
CREATE TABLE data_source (
    id SERIAL PRIMARY KEY,
    name VARCHAR(300) NOT NULL,
    description VARCHAR(1000),
    data_source_type_id INT NOT NULL,
    connection_mode_id INT NOT NULL,
    host VARCHAR(700) NOT NULL,
    port INT CHECK (port BETWEEN 1 AND 65535),
    username VARCHAR(100) NOT NULL,
    password VARCHAR(80) NOT NULL,
    database_name VARCHAR(100) DEFAULT '',
    saved BOOLEAN NOT NULL DEFAULT TRUE,
    active BOOLEAN NOT NULL DEFAULT TRUE,
    updated_user_id INT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT current_timestamp(5),
    updated_at TIMESTAMPTZ DEFAULT current_timestamp(5),
    FOREIGN KEY (data_source_type_id) REFERENCES data_source_type (id),
    FOREIGN KEY (connection_mode_id) REFERENCES connection_mode (id),
    FOREIGN KEY (updated_user_id) REFERENCES person (id)
);

ALTER TABLE data_source
ADD CONSTRAINT port_check CHECK (port > 0 AND port < 65535);

select * from data_source;

-- this is for adding or editing the types of visualization type, We provide KPI cards, reports as visualization_type
drop table if exists visualization_type;
create table visualization_type(
	id serial primary key,
	type_name varchar(50),
	description varchar(200) not null default '',
	active bool not null default(true),
	created_at timestamptz not null default current_timestamp (5), 
  	updated_at timestamptz not null default current_timestamp (5)
);
insert into visualization_type(type_name, description) values
	('KPI-cards', 'This KPI card widgets will show the aggregated KPI values from the data source which is configured with it'),
	('Report', 'This type denotes that the visualization is of type Report.');
	

drop table if exists visualization;
create table visualization(
	id bigserial primary key,
	name varchar(100) not null unique,
	description varchar(349) not null default '',
	visualization_type_lid int,
	default_data_source_lid int,
	group_name varchar(100) not null default 'other', -- this group_name will help in grouping of widgets inside the canvas for the user.
	is_widget_resizeable bool not null default(false),
	is_widget_movable bool not null default(false),
	active bool not null default(true),
	updated_user_lid int not null,
	created_at timestamptz not null default current_timestamp (5), 
  	updated_at timestamptz not null default current_timestamp (5),
	constraint visualization_type_foreign_key 
		foreign key(visualization_type_lid)
		references visualization_type(id)
);

-- every visualization widget either be KPI's or charts visualizations inside report must always have data-source-connection-string