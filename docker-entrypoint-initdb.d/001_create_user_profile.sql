CREATE TABLE IF NOT EXISTS user_profile (
    email VARCHAR(255) PRIMARY KEY,
    nama_pt VARCHAR(255),
    nama_bank VARCHAR(255),
    no_rekening VARCHAR(50),
    penanggung_jawab VARCHAR(255),
    alamat TEXT,
    kabupaten VARCHAR(100)
);
