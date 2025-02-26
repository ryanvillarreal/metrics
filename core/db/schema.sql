-- metrics table: Represents an individual Nmap scan or single xml file
CREATE TABLE IF NOT EXISTS scans (
    id INTEGER PRIMARY KEY AUTOINCREMENT,  -- Unique scan ID
    type TEXT NOT NULL,                    -- type allows for more versatile input later
    start_time TIMESTAMP NOT NULL         -- this is the import time not the scan time
);

-- host table: Hosts discovered during a scan
CREATE TABLE IF NOT EXISTS host (
    id INTEGER PRIMARY KEY AUTOINCREMENT,  -- Unique host ID
    scans_id INTEGER NOT NULL,           -- Foreign key to metrics (scan)

  FOREIGN KEY (scans_id) REFERENCES metrics(id) ON DELETE CASCADE
);


