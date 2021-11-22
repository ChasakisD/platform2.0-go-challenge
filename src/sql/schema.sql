CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

DO $outer$
BEGIN
	IF NOT EXISTS (SELECT 1 FROM pg_proc WHERE proname = 'auto_update_on_edit') THEN
		CREATE FUNCTION auto_update_on_edit()   
		RETURNS TRIGGER AS $$
		BEGIN
			NEW.updated_at = CURRENT_TIMESTAMP;
			RETURN NEW;   
		END;
		$$ language 'plpgsql';
	END IF;
END$outer$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'gender') THEN
        CREATE TYPE GENDER AS ENUM ( 'Male', 'Female' );
    END IF;
END$$;

CREATE TABLE IF NOT EXISTS users (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	username VARCHAR NOT NULL UNIQUE,
	password VARCHAR NOT NULL,
	email VARCHAR NOT NULL UNIQUE,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE TRIGGER user_auto_update_on_edit 
BEFORE UPDATE ON users
FOR EACH ROW EXECUTE PROCEDURE auto_update_on_edit();

CREATE TABLE IF NOT EXISTS assets (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	description VARCHAR NOT NULL UNIQUE,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE TRIGGER asset_auto_update_on_edit 
BEFORE UPDATE ON assets
FOR EACH ROW EXECUTE PROCEDURE auto_update_on_edit();

CREATE TABLE IF NOT EXISTS users_assets (
	user_id UUID NOT NULL,
	asset_id UUID NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY(user_id, asset_id),
	CONSTRAINT fk_user
      	FOREIGN KEY(user_id) 
		REFERENCES users(id),
	CONSTRAINT fk_asset
    	FOREIGN KEY(asset_id)
		REFERENCES assets(id)
);

CREATE OR REPLACE TRIGGER user_asset_auto_update_on_edit 
BEFORE UPDATE ON users_assets
FOR EACH ROW EXECUTE PROCEDURE auto_update_on_edit();

CREATE TABLE IF NOT EXISTS charts  (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	asset_id UUID NOT NULL UNIQUE,
	x_axes VARCHAR NOT NULL,
	y_axes VARCHAR NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT fk_asset
		FOREIGN KEY(asset_id)
		REFERENCES assets(id)
);

CREATE OR REPLACE TRIGGER chart_auto_update_on_edit 
BEFORE UPDATE ON charts
FOR EACH ROW EXECUTE PROCEDURE auto_update_on_edit();

CREATE TABLE IF NOT EXISTS chart_points (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	chart_id UUID NOT NULL,
	x_value REAL NOT NULL,
	y_value REAL NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT fk_chart
		FOREIGN KEY(chart_id)
		REFERENCES charts(id)
);

CREATE OR REPLACE TRIGGER chart_point_auto_update_on_edit 
BEFORE UPDATE ON chart_points
FOR EACH ROW EXECUTE PROCEDURE auto_update_on_edit();

CREATE TABLE IF NOT EXISTS insights (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	asset_id UUID NOT NULL UNIQUE,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT fk_asset
		FOREIGN KEY(asset_id)
		REFERENCES assets(id)
);

CREATE OR REPLACE TRIGGER insight_auto_update_on_edit 
BEFORE UPDATE ON insights
FOR EACH ROW EXECUTE PROCEDURE auto_update_on_edit();

CREATE TABLE IF NOT EXISTS audience_stat_types (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	title VARCHAR NOT NULL,
	title_formatted VARCHAR NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE TRIGGER audience_stat_type_auto_update_on_edit 
BEFORE UPDATE ON audience_stat_types
FOR EACH ROW EXECUTE PROCEDURE auto_update_on_edit();

CREATE TABLE IF NOT EXISTS audiences (
	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	asset_id UUID NOT NULL UNIQUE,
	stat_type_id UUID NOT NULL,
	gender GENDER,
	birth_country VARCHAR,
	age_group_min INT,
	age_group_max INT,
	stat_type_value REAL DEFAULT 0.0,
	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT fk_asset
		FOREIGN KEY(asset_id)
		REFERENCES assets(id),
	CONSTRAINT fk_stat_type
		FOREIGN KEY(stat_type_id)
		REFERENCES audience_stat_types(id)
);

CREATE OR REPLACE TRIGGER audience_auto_update_on_edit 
BEFORE UPDATE ON audiences
FOR EACH ROW EXECUTE PROCEDURE auto_update_on_edit();