CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    github_id VARCHAR(255) NOT NULL UNIQUE,
    username VARCHAR(255),
    full_name VARCHAR(255),
    avatar_url TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
