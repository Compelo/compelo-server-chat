#Documentazione Server

##Specifica generale
Il server della api in Go lavora sulla porta 5051, il server chat lavora sulal porta 5052

##Specifica protocollo CCP (Compleo Chat Protocol)
Il protocollo si basa, a livello transport, sulla specifica QUIC (Next generation TCP) specificato dalla IETF a questo link: https://quicwg.org/

Per creare una chat il client invia "CHATNEW [idUtenteMittente] [idUtenteDestinatario]"
Una volta creata la chat sul database, il server comunica al destinatario la creazione della nuova chat, comunico al mittente l'id corrispondente della chat con "CREATED [idChat]".

Ogni volta che un utente visualizza la chat, li viene inviato dal server la serializzazzione dell'oggetto corrispondente in JSON, la richiesta è:
"CHATVISUALIZE [idChat]".

Ogni volta che si invia un messaggio si utilizza "NEWMESSAGE [idUtente] [tipoMessaggio] [chars]".

Ogniuno dei due client può chiudere la chat in qualsiasi momento eseguendo il comando "CLOSE [idUtente] [idChat]"
