CREATE TABLE meal(
                       id UUID NOT NULL PRIMARY KEY,
                       name VARCHAR(55) NOT NULL,
                       price INTEGER NOT NULL,
                       type INTEGER NOT NULL
);
CREATE TABLE salad(
                        id UUID NOT NULL PRIMARY KEY,
                        name VARCHAR(55) NOT NULL,
                        price INTEGER NOT NULL
);
CREATE TABLE drinks(
                         id UUID NOT NULL PRIMARY KEY,
                         name VARCHAR(55) NOT NULL,
                         price INTEGER NOT NULL
);
CREATE TABLE products(
                           id UUID NOT NULL PRIMARY KEY,
                           name VARCHAR(55) NOT NULL,
                           time TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
CREATE TABLE recipe(
                         id UUID NOT NULL PRIMARY KEY,
                         meal_id UUID NULL REFERENCES meal(id),
                         product_ids UUID[],
                         salad_id UUID NULL REFERENCES salad(id)
);
CREATE TABLE basket(
                         id UUID NULL PRIMARY KEY,
                         meal_ids UUID[],
                         salad_ids UUID[],
                         drinks_ids UUID[],
                         sum INTEGER NOT NULL
);
