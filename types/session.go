package types

type Session struct {
	Connessioni []Connection
}

func (s *Session) FindConnection(ip string) (int, string) {
	for pos, c := range s.Connessioni {
		if ip == c.Utente.IPAddress {
			return pos, ""
		}
	}

	return 0, "Non trovato nulla"
}

func (s *Session) NewConenction(ip string, id int64) Connection {
	var conn Connection
	conn.Utente.DatabaseID = id
	conn.Utente.IPAddress = ip

	s.Connessioni = append(s.Connessioni, conn)

	return conn
}

func (s *Session) RemoveConnection(ip string) {
	index := 0
	for _, k := range s.Connessioni {
		if k.Utente.IPAddress != ip {
			s.Connessioni[index] = k
			index++
		}
	}

	s.Connessioni = s.Connessioni[:index]
}
