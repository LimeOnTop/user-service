CREATE TABLE IF NOT EXISTS user_preferences (
    user_id INT PRIMARY KEY,
    preference_name VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);