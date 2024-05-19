DO $$ BEGIN
    IF NOT EXISTS (
        SELECT FROM pg_catalog.pg_roles WHERE rolname = 'novo_usuario'
    ) THEN
        CREATE ROLE novo_usuario WITH LOGIN PASSWORD 'nova_senha';
    END IF;
END $$;