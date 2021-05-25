DROP TABLE IF EXISTS reaction_role;
DROP TABLE IF EXISTS member_role;
DROP TABLE IF EXISTS mod_mail_entry;
DROP TABLE IF EXISTS member;
DROP TABLE IF EXISTS guild_role;


CREATE TABLE IF NOT EXISTS guild_role (
    id VARCHAR(20) PRIMARY KEY,
    sticky_role BOOLEAN NOT NULL,
    auto_role BOOLEAN NOT NULL,
    mod_role BOOLEAN NOT NULL
);

CREATE TABLE IF NOT EXISTS member (
    id VARCHAR(20) PRIMARY KEY
);

CREATE TABLE IF NOT EXISTS mod_mail_entry (
    id SERIAL PRIMARY KEY,
    offender_id VARCHAR(20) NOT NULL REFERENCES member(id),
    mod_id VARCHAR(20) NOT NULL REFERENCES member(id),
    punishment VARCHAR(5) NOT NULL,
    reason VARCHAR(250) NOT NULL,
    timestamp TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS member_role (
    id SERIAL PRIMARY KEY,
    member_id VARCHAR(20) NOT NULL REFERENCES member(id),
    role_id VARCHAR(20) NOT NULL REFERENCES guild_role(id)
);

CREATE TABLE IF NOT EXISTS reaction_role (
    id SERIAL PRIMARY KEY,
    role_id VARCHAR(20) NOT NULL REFERENCES guild_role(id),
    channel_id VARCHAR(20) NOT NULL,
    message_id VARCHAR(20) NOT NULL,
    emoji VARCHAR NOT NULL
);
