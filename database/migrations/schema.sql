create table web3game.games
(
    id             int auto_increment primary key,
    gameId         varchar(256) null,
    players        varchar(1024) null,
    deposit_amount int null,
    status         int null,
    owner          int null,
    created_at     timestamp null,
    winner         int null,
    constraint games_gameId_uindex unique (gameId)
);


create table web3game.players
(
    id           int auto_increment primary key,
    wallet       varchar(256) null,
    display_name varchar(30) null,
    email        varchar(256) null,
    pass         varchar(30) null,
    constraint players_email_uindex unique (email),
    constraint players_wallet_uindex unique (wallet)
);


create table web3game.payments
(
    id           int auto_increment primary key,
    tx_hash       varchar(256) null,
    player_address varchar(256) null,
    value varchar(256) null
);


create table web3game.nonce
(
    id           int auto_increment primary key,
    wallet_address       varchar(256) null,
    nonce varchar(256) null
);



