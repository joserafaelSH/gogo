CREATE TABLE IF NOT EXISTS energy_read(
  id VARCHAR(36) PRIMARY KEY,
  date_ref VARCHAR(255) NOT NULL,
  date_payment TIMESTAMPTZ,
  date_due TIMESTAMPTZ NOT NULL,
  kwh_compensado DOUBLE PRECISION,
  kwh_billed DOUBLE PRECISION,
  kwh_consumed DOUBLE PRECISION,
  kwh_acumulated_credit_p DOUBLE PRECISION DEFAULT 0, 
  kwh_acumulated_credit_fp DOUBLE PRECISION DEFAULT 0,
  dealership_bill_cost DECIMAL(9, 2),
  dealership_extra_fees DECIMAL(9, 2),
  dealership_bill_status VARCHAR(255) DEFAULT 'PENDING',
  is_micro_min_generator BOOLEAN DEFAULT FALSE,
  file_key VARCHAR(255) NOT NULL,
  created_at TIMESTAMPTZ DEFAULT NOW(),
  updated_at TIMESTAMPTZ,
  client_id VARCHAR(255) NOT NULL,
  scraper_status VARCHAR(255) DEFAULT 'PENDING',
  valor_compensado DECIMAL(9, 2),
  valor_energia DECIMAL(9, 2),
  dealership_bill_id VARCHAR(255) UNIQUE,
  consumo_minimo DOUBLE PRECISION DEFAULT 0
);

DO $$ 
BEGIN
  FOR i IN 1..5000 LOOP
    INSERT INTO energy_read (
      date_ref,
      date_due,
      kwh_compensado,
      kwh_billed,
      kwh_consumed,
      kwh_acumulated_credit_p,
      kwh_acumulated_credit_fp,
      dealership_bill_cost,
      dealership_extra_fees,
      dealership_bill_status,
      is_micro_min_generator,
      file_key,
      client_id,
      scraper_status,
      valor_compensado,
      valor_energia,
      dealership_bill_id,
      consumo_minimo
    ) VALUES (
      '2023-01-01', -- Replace with your desired date
      NOW() + (i || ' days')::INTERVAL,
      RANDOM() * 100,
      RANDOM() * 100,
      RANDOM() * 100,
      RANDOM() * 100,
      RANDOM() * 100,
      RANDOM() * 100,
      RANDOM() * 100,
      CASE WHEN i % 2 = 0 THEN 'PAID' ELSE 'PENDING' END,
      i % 2 = 0,
      'file_key_' || i,
      'client_id_' || i,
      CASE WHEN i % 2 = 0 THEN 'COMPLETED' ELSE 'PENDING' END,
      RANDOM() * 100,
      RANDOM() * 100,
      'bill_id_' || i,
      RANDOM() * 100
    );
  END LOOP;
END $$;






