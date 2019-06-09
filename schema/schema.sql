CREATE TABLE IF NOT EXISTS Chat (
	ID 				SERIAL NOT NULL PRIMARY KEY,
	Nome			VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS Usuario (
	ID				SERIAL NOT NULL PRIMARY KEY,
	Nome			VARCHAR(100) NOT NULL,
	Foto			BYTEA NOT NULL
);

CREATE TABLE IF NOT EXISTS Chat_tem_Usuario (
	IDChat			SERIAL NOT NULL REFERENCES Chat(ID),
	IDUsuario		SERIAL NOT NULL REFERENCES Usuario(ID),
	PRIMARY KEY(IDChat, IDUsuario)
);

CREATE TABLE IF NOT EXISTS Mensagem (
	IDMsg			SERIAL NOT NULL PRIMARY KEY,
	Hr_env			TIMESTAMP WITHOUT TIME ZONE DEFAULT (NOW() AT TIME ZONE 'UTC+3') NOT NULL,
	Msg				VARCHAR(500) NOT NULL,
	IDChat			SERIAL NOT NULL REFERENCES Chat(ID),
	IDUsuario		SERIAL NOT NULL REFERENCES Usuario(ID)
);