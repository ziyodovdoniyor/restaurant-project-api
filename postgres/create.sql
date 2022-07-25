CREATE TABLE first_meal (
    id UUID PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE,
    ingredients TEXT NOT NULL,
    price INTEGER NOT NULL,
    cooked_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE second_meal (
    id UUID PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE,
    ingredients TEXT NOT NULL,
    price INTEGER NOT NULL,
    cooked_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE dessert (
    id UUID PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE,
    ingredients TEXT NOT NULL,
    price INTEGER NOT NULL,
    cooked_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE salad (
    id UUID PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE,
    ingredients TEXT NOT NULL,
    price INTEGER NOT NULL,
    cooked_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE beverage (
    id UUID PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE,
    ingredients TEXT NOT NULL,
    price INTEGER NOT NULL,
    cooked_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE table (
    id UUID PRIMARY KEY,
    table_number INTEGER
);

CREATE TABLE purchase (
    table_id UUID REFERENCES client (id),
    first_meal_id UUID REFERENCES first_meal (id),
    second_meal_id UUID REFERENCES second_meal (id),
    dessert_id UUID REFERENCES dessert (id),
    salad_id UUID REFERENCES salad (id),
    beverage_id UUID REFERENCES beverage (id),
    total INTEGER NOT NULL,
    purchased_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);