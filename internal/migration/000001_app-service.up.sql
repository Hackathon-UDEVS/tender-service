-- Tender jadvali
CREATE TABLE tender (
                        id SERIAL PRIMARY KEY,
                        client_id INT NOT NULL, -- Foydalanuvchi ID
                        title VARCHAR(255) NOT NULL,
                        description TEXT,
                        attachment VARCHAR(250) NOT NULL,
                        deadline TIMESTAMP NOT NULL CHECK (deadline > (CURRENT_TIMESTAMP AT TIME ZONE 'UTC' + INTERVAL '5 hours')), -- Deadline kelajakda bo‘lishi kerak
                        budget NUMERIC(12, 2) NOT NULL, -- Byudjet
                        status VARCHAR(50) NOT NULL CHECK (status IN ('open', 'closed', 'cancelled', 'awarded')) -- Enum sifatida ishlatiladi
);

-- Bid jadvali
CREATE TABLE bid (
                     id SERIAL PRIMARY KEY,
                     tender_id INT NOT NULL REFERENCES tender(id) ON DELETE CASCADE, -- Tender bilan bog‘langan, o‘chirilganda bid ham o‘chiriladi
                     contractor_id INT NOT NULL, -- Kontraktor ID
                     price NUMERIC(12, 2) NOT NULL, -- Narx
                     delivery_time INTERVAL NOT NULL, -- Yetkazib berish muddati
                     comments TEXT, -- Izohlar
                     status VARCHAR(50) NOT NULL CHECK (status IN ('pending', 'accepted', 'rejected')) -- Enum sifatida ishlatiladi
);

-- Notification jadvali
CREATE TABLE notification (
                              id SERIAL PRIMARY KEY,
                              user_id INT NOT NULL, -- Foydalanuvchi ID
                              message TEXT NOT NULL, -- Xabar
                              relation_id INT NOT NULL, -- Bog‘langan ID (tender yoki bid)
                              type VARCHAR(50) NOT NULL CHECK (type IN ('tender', 'bid')), -- Tender yoki Bid turida
                              created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL -- Yaratilgan vaqt
);

-- Indekslar (tezkor qidiruv uchun)
CREATE INDEX idx_tender_client_id ON tender(client_id);
CREATE INDEX idx_bid_tender_id ON bid(tender_id);
CREATE INDEX idx_notification_user_id ON notification(user_id);

-- Deadline o‘tganda statusni avtomatik 'closed' qilish uchun funksiya
CREATE OR REPLACE FUNCTION set_tender_status_to_closed()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.deadline <= CURRENT_TIMESTAMP AT TIME ZONE 'UTC' + INTERVAL '5 hours' THEN
        NEW.status := 'closed';
END IF;
RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Tender jadvali uchun trigger
CREATE TRIGGER check_deadline_before_insert_or_update
    BEFORE INSERT OR UPDATE ON tender
                         FOR EACH ROW
                         EXECUTE FUNCTION set_tender_status_to_closed();

-- Sekundlarni avtomatik olib tashlash uchun funksiya
CREATE OR REPLACE FUNCTION remove_seconds_from_deadline()
RETURNS TRIGGER AS $$
BEGIN
    NEW.deadline := date_trunc('minute', NEW.deadline); -- Sekundlarni olib tashlash
RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Tender uchun deadline formatlash trigger
CREATE TRIGGER format_deadline_before_insert
    BEFORE INSERT OR UPDATE ON tender
                         FOR EACH ROW
                         EXECUTE FUNCTION remove_seconds_from_deadline();
