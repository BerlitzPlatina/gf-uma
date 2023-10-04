create table articles (
	id bigint unsigned not null PRIMARY KEY AUTO_INCREMENT,
	title varchar(255),
	content text(255),
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
)