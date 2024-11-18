

CREATE ROLE fisher_service WITH LOGIN PASSWORD 'ecivreSrehsif';
GRANT CONNECT ON DATABASE fishing_hell TO fisher_service;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA fisher TO fisher_service;
ALTER SCHEMA fisher OWNER TO fisher_service;

CREATE ROLE account_service WITH LOGIN PASSWORD 'ecivreStnuocca';
GRANT CONNECT ON DATABASE fishing_hell TO fisher_service;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA account TO account_service;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA shop TO account_service;
ALTER SCHEMA account OWNER TO account_service;


CREATE ROLE market_service WITH LOGIN PASSWORD 'ecivreStekram';
ALTER SCHEMA market OWNER TO market_service;

CREATE ROLE shop_service WITH LOGIN PASSWORD 'ecivreSpohs';
ALTER SCHEMA shop OWNER TO shop_service;

