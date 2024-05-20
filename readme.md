\### Instruções para configuração e uso do Docker e PostgreSQL

#### Parar e limpar containers Docker

Para parar todos os containers e limpar os volumes Docker, use os seguintes comandos:
```bash
docker-compose down
docker system prune -v
```

#### Executar aplicação

Para iniciar a aplicação usando Docker Compose, execute:
```bash
docker-compose up -d
```

#### Criar usuário via API

Para criar um novo usuário utilizando a API, você pode usar o comando `curl`:
```bash
curl -X POST http://localhost:8080/users -H "Content-Type: application/json" -d '{"name":"João Silva","email":"joaosilva@example.com","password":"senha123"}'
```

#### Conectar ao PostgreSQL

Para se conectar ao banco de dados PostgreSQL, use um dos seguintes comandos:

Conexão com o banco `go_game` usando o usuário `admin`:
```bash
psql -U admin -d go_game
```

Conexão com o banco `go-game` usando o usuário `postgres`:
```bash
psql -U postgres -d go-game
```

Conexão ao PostgreSQL em um host específico:
```bash
psql -h localhost -p 5432 -U admin -d go_game -W
```

#### Comandos úteis no psql

Listar todos os usuários de banco de dados:
```sql
\du
```

Listar todos os bancos de dados:
```sql
\l
```

Conectar-se a um banco de dados específico:
```sql
\c go-game
```

Obter ajuda e sair do psql:
```sql
\?  -- Ajuda
\q  -- Sair
```

Listar todas as tabelas no banco de dados atual:
```sql
\dt
```

Descrever uma tabela específica:
```sql
\d todos;
```

#### Limpar Docker

Para limpar todos os recursos Docker (containers, volumes, imagens, etc.), use os comandos:
```bash
docker system prune -v
docker system prune -a -v
```

---

