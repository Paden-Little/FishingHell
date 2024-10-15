CREATE ROLE fisher_service WITH LOGIN PASSWORD 'ecivreSrehsif';
ALTER SCHEMA fisher OWNER TO fisher_service;

CREATE ROLE account_service WITH LOGIN PASSWORD 'ecivreStnuocca';
ALTER SCHEMA account OWNER TO account_service;

CREATE ROLE market_service WITH LOGIN PASSWORD 'ecivreStekram';
ALTER SCHEMA market OWNER TO market_service;

CREATE ROLE shop_service WITH LOGIN PASSWORD 'ecivreSpohs';
ALTER SCHEMA shop OWNER TO shop_service;

