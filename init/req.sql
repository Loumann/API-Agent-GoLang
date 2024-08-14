
CREATE TABLE agents (
                        id SERIAL PRIMARY KEY,
                        agentname VARCHAR(255) NOT NULL,
                        status VARCHAR(50) NOT NULL
);

CREATE TABLE task (
                      id SERIAL PRIMARY KEY,
                      quest TEXT NOT NULL,
                      agentid INT,
                      FOREIGN KEY (agentid) REFERENCES agents(id) ON DELETE CASCADE
);