-- Função para atualizar o campo updated_at
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = NOW();
   RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger para atualizar o campo updated_at antes de inserções ou atualizações
CREATE TRIGGER update_users_modtime
BEFORE INSERT OR UPDATE ON users
FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();