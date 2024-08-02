CREATE TABLE IF NOT EXISTS repositories (
                                            id SERIAL PRIMARY KEY,
                                            name VARCHAR(255) NOT NULL,
    description TEXT,
    url VARCHAR(255),
    language VARCHAR(50),
    forks_count INT,
    stars_count INT,
    open_issues_count INT,
    watchers_count INT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
    );

CREATE TABLE IF NOT EXISTS commits (
                                       id SERIAL PRIMARY KEY,
                                       repository_id INT REFERENCES repositories(id),
    message TEXT,
    author VARCHAR(255),
    date TIMESTAMP,
    url VARCHAR(255),
    UNIQUE (repository_id, url)
    );
