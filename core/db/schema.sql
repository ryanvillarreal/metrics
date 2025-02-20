-- metrics table: Represents an individual Nmap scan or single xml file
CREATE TABLE IF NOT EXISTS scans (
    id INTEGER PRIMARY KEY AUTOINCREMENT,  -- Unique scan ID
    type TEXT NOT NULL,                    -- type allows for more versatile input later
    start_time TIMESTAMP NOT NULL,         -- Scan start time
    end_time TIMESTAMP                     -- Scan end time (optional)
);

-- host table: Hosts discovered during a scan
CREATE TABLE IF NOT EXISTS host (
    id INTEGER PRIMARY KEY AUTOINCREMENT,  -- Unique host ID
    scans_id INTEGER NOT NULL,           -- Foreign key to metrics (scan)
    address TEXT NOT NULL,                 -- Host IP or hostname
    status TEXT NOT NULL,                  -- Host status (up/down)
    hostname TEXT,                         -- Resolved hostname (optional)
    os TEXT,                               -- Detected OS (optional)
    ports TEXT,                            -- Open ports (comma-separated or JSON)
    FOREIGN KEY (scans_id) REFERENCES metrics(id) ON DELETE CASCADE
);

