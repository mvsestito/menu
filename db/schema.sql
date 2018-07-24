CREATE EXTENSION IF NOT EXISTS citext;

-- auto UPDATE ts trigger
CREATE OR REPLACE FUNCTION update_updated_ts_column()
RETURNS TRIGGER AS $$
BEGIN
        NEW.updated_ts= now();
            RETURN NEW;
END;
$$ language 'plpgsql';


CREATE TABLE IF NOT EXISTS restaurants (
    id SERIAL PRIMARY KEY,
    name citext NOT NULL,
    street TEXT NOT NULL,
    city TEXT NOT NULL,
    state TEXT NOT NULL,
    zip TEXT NOT NULL,
    created_ts TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_ts TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    UNIQUE(name, street)
    );

DROP TRIGGER IF EXISTS update_restaurants_updated_ts ON restaurants;
CREATE TRIGGER update_restaurants_updated_ts BEFORE UPDATE ON restaurants FOR EACH ROW EXECUTE PROCEDURE update_updated_ts_column();

CREATE TABLE IF NOT EXISTS items (
    id SERIAL PRIMARY KEY,
    restaurant_id int4 NOT NULL REFERENCES restaurants (id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    kind TEXT NOT NULL,
    description TEXT,
    modifiers TEXT[],
    created_ts TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_ts TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    UNIQUE(restaurant_id, name, kind)
    );

DROP TRIGGER IF EXISTS update_items_updated_ts ON items;
CREATE TRIGGER update_items_updated_ts BEFORE UPDATE ON items FOR EACH ROW EXECUTE PROCEDURE update_updated_ts_column();

