package types

type Session struct {
	Connessioni []Connection
}

func (s *Session) FindConnection(ip string) (Connection, string) {
	for _, c := range s.Connessioni {
		if ip == c.Utente.IPAddress {
			return c, ""
		}
	}

	var conn Connection
	return conn, "Non trovato nulla"
}

func (s *Session) NewConenction(ip string, id int64) Connection {
	var conn Connection
	conn.Utente.DatabaseID = id
	conn.Utente.IPAddress = ip

	s.Connessioni = append(s.Connessioni, conn)

	return conn
}
