CREATE TABLE IF NOT EXISTS lo_bulanan (
    no SERIAL PRIMARY KEY,
    tanggal DATE NOT NULL,
    no_so VARCHAR(50),
    no_lo VARCHAR(50),
    jumlah_tabung INT DEFAULT 0,
    jumlah_kg NUMERIC(10,2) DEFAULT 0,
    tarif NUMERIC(10,2) DEFAULT 0,
    biaya_angkut NUMERIC(10,2) DEFAULT 0
);
