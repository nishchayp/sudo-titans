-- insert into mcq_details(
-- 	question_id, 
-- 	level, 
-- 	question_1, 
-- 	option_1_a, option_1_b, option_1_c, option_1_d, 
-- 	question_2, 
-- 	option_2_a, option_2_b, option_2_c, option_2_d, 
-- 	question_3, 
-- 	option_3_a, option_3_b, option_3_c, option_3_d, 
-- 	question_4, 
-- 	option_4_a, option_4_b, option_4_c, option_4_d) 
-- values(
-- 	"MCQ_ ",
-- 	1,
-- 	"",
-- 	"", "", "", "",
-- 	"",
-- 	"", "", "", "",
-- 	"",
-- 	"", "", "", "",
-- 	"",
-- 	"", "", "", ""
-- );

insert into mcq_details(
	question_id, 
	level, 
	question_1, 
	option_1_a, option_1_b, option_1_c, option_1_d, 
	question_2, 
	option_2_a, option_2_b, option_2_c, option_2_d, 
	question_3, 
	option_3_a, option_3_b, option_3_c, option_3_d, 
	question_4, 
	option_4_a, option_4_b, option_4_c, option_4_d) 
values(
	"MCQ_1",
	1,
	
	"Which of the following pieces of information can be found in the IP header?",
	"Source address of the IP packet", "Destination address for the IP packet", "Sequence number of the IP packet", "Both (A) and (B)",
	
	"In asymmetric key cryptography, the private key is kept by",
	"Sender", "Receiver", "Sender and receiver", "All the connected devices",
	
	"________ is the science and art of transforming messages to make them secure and immune to attacks.",
	"Cryptography", "Cryptanalysis", "Both (a) and (b)", "None",
	
	"A combination of an encryption algorithm and a decryption algorithm is called a ________.",
	"Cipher", "Secret", "Key", "None"
);

insert into mcq_details(
	question_id, 
	level, 
	question_1, 
	option_1_a, option_1_b, option_1_c, option_1_d, 
	question_2, 
	option_2_a, option_2_b, option_2_c, option_2_d, 
	question_3, 
	option_3_a, option_3_b, option_3_c, option_3_d, 
	question_4, 
	option_4_a, option_4_b, option_4_c, option_4_d) 
values(
	"MCQ_2",
	1,
	
	"The _______ is a number or a set of numbers on which the cipher operates.",
	"Cipher", "Secret", "Key", "None",
	
	"The Caesar cipher is a _______ cipher that has a key of 3.",
	"Transposition", "Additive", "Shift", "None",
	
	" In Caesar's Cipher, if the ciphertext is “oyrehehvm” and the shift(key) is 5, the plaintext is:",
	"jtmzczcqh", "kunadadri", "islybybpg", "tdwjmjmar",
	
	"Cipher which uses exclusive-or operation as defined in computer science is called",
	"Caesar cipher", "XOR cipher", "Cipher", "Cipher text"
);

insert into mcq_details(
	question_id, 
	level, 
	question_1, 
	option_1_a, option_1_b, option_1_c, option_1_d, 
	question_2, 
	option_2_a, option_2_b, option_2_c, option_2_d, 
	question_3, 
	option_3_a, option_3_b, option_3_c, option_3_d, 
	question_4, 
	option_4_a, option_4_b, option_4_c, option_4_d) 
values(
	"MCQ_3",
	1,
	
	"Services running on a system are determined by _____________.",
	"The system’s IP address.", "The Active Directory", "The system’s network name", "The port assigned",
	
	"What are hybrid attacks?",
	"An attempt to crack passwords using words that can be found in dictionary.", "An attempt to crack passwords by replacing characters of a dictionary word with numbers and symbols.", "An attempt to crack passwords using a combination of characters, numbers, and symbols.", "An attempt to crack passwords by replacing characters with numbers and symbols.",
	
	" Which of the following is not a type of port scan?",
	"Strobe", "Vanilla", "Stealth Scan", "Cookie Scan",
	
	"Anarkali digitally signs a message and sends it to Salim. Verification of the signature by Salim requires",
	"Anarkali’s public key", "Salim’s public key.", "Salim’s private key.", "Anarkali’s private key."
);

insert into mcq_details(
	question_id, 
	level, 
	question_1, 
	option_1_a, option_1_b, option_1_c, option_1_d, 
	question_2, 
	option_2_a, option_2_b, option_2_c, option_2_d, 
	question_3, 
	option_3_a, option_3_b, option_3_c, option_3_d, 
	question_4, 
	option_4_a, option_4_b, option_4_c, option_4_d) 
values(
	"MCQ_4",
	2,
	
	"We don't want our undeliverable packets to hop around forever. What feature/flag limits the life of an IP packet on the network?",
	"Time to Live counter", "Subnet Mask", "Header Checksum", "Wackamole field",
	
	"A transposition cipher reorders (permutes) symbols in a",
	"Block of packets", "Block of slots", "Block of signals", "Block of symbols",
	
	" In Cryptography, when text is treated at bit level, each character is replaced by",
	"4 bits", "6 bits", "8 bits", "10 bits",
	
	"Why would HTTP Tunneling be used?",
	"To identify proxy servers", "Web activity is not scanned", "To bypass a firewall", "HTTP is a easy protocol to work with"
);

insert into mcq_details(
	question_id, 
	level, 
	question_1, 
	option_1_a, option_1_b, option_1_c, option_1_d, 
	question_2, 
	option_2_a, option_2_b, option_2_c, option_2_d, 
	question_3, 
	option_3_a, option_3_b, option_3_c, option_3_d, 
	question_4, 
	option_4_a, option_4_b, option_4_c, option_4_d) 
values(
	"MCQ_5",
	2,
	
	"A packet with no flags set is which type of scan?",
	"TCP", "XMAS", "IDLE", "NULL",
	
	"Sniffing is used to perform ______________ fingerprinting.",
	"Passive stack", "Active stack", "Passive banner grabbing", "Scanned",
	
	"What are the port states determined by Nmap?",
	"Active, inactive, standby", "Open, half-open, closed", "Open, filtered, unfiltered", "Active, closed, unused",
	
	"Which of the following is not a component used in SSL?",
	"SSL Recorded Protocol", "Handshake protocol", "Change Cipher Spec", "Encoding algorithms"
);

insert into mcq_details(
	question_id, 
	level, 
	question_1, 
	option_1_a, option_1_b, option_1_c, option_1_d, 
	question_2, 
	option_2_a, option_2_b, option_2_c, option_2_d, 
	question_3, 
	option_3_a, option_3_b, option_3_c, option_3_d, 
	question_4, 
	option_4_a, option_4_b, option_4_c, option_4_d) 
values(
	"MCQ_6",
	2,
	
	"Using public key cryptography, X adds a digital signature \sigma to message M, encrypts < M, \sigma >, and sends it to Y, where it is decrypted. Which one of the following sequences of keys is used for the operations?",
	"Encryption: X’s private key followed by Y’s private key; Decryption: X’s public key followed by Y’s public key", "Encryption: X’s private key followed by Y’s public key; Decryption: X’s public key followed by Y’s private key", "Encryption: X’s public key followed by Y’s private key; Decryption: Y’s public key followed by X’s private key", "Encryption: X’s private key followed by Y’s public key; Decryption: Y’s private key followed by X’s public key",
	
	"Suppose that everyone in a group of N people wants to communicate secretly with the N–1 others using symmetric key cryptographic system. The communication between any two persons should not be decodable by the others in the group. The number of keys required in the system as a whole to satisfy the confidentiality requirement is",
	"2N", "N(N-1)", "N(N-1)/2", "(N-1)^2",
	
	"A sender is employing public key cryptography to send a secret message to a receiver. Which one of the following statements is TRUE?",
	"Sender encrypts using receiver’s public key", "Sender encrypts using his own public key", "Receiver decrypts using sender’s public key", "Receiver decrypts using his own public key",
	
	"Exponentiation is a heavily used operation in public key cryptography. Which of the following options is the tightest upper bound on the number of multiplications required to compute bn mod m,0<=b,n<=m ?",
	"O(logn)", "O(sqrt(n))", "O(n/logn)", "O(n)"
);

insert into mcq_details(
	question_id, 
	level, 
	question_1, 
	option_1_a, option_1_b, option_1_c, option_1_d, 
	question_2, 
	option_2_a, option_2_b, option_2_c, option_2_d, 
	question_3, 
	option_3_a, option_3_b, option_3_c, option_3_d, 
	question_4, 
	option_4_a, option_4_b, option_4_c, option_4_d) 
values(
	"MCQ_7",
	2,
	
	"A firewall is to be configured to allow hosts in a private network to freely open TCP connections and send packets on open connections. However, it will only allow external hosts to send packets on existing open TCP connections or connections that are being opened (by internal hosts) but not allow them to open TCP connections to hosts in the private network. To achieve this the minimum capability of the firewall should be that of",
	"A combinational circuit", "A finite automaton", "A pushdown automaton with one stack", "A pushdown automaton with two stacks",
	
	"S1- MD5 is vulnerable to the Birthday attack.
S2- traceroute uses the 'Destination port unreachable' ICMP error message.
S3- 3DES is a type of Public Key Encryption Algorithm.
S4- IPv6 has a 40-byte fixed header size(without Options fields).
How many of the above statements are true?
",
	"2", "3", "4", "1",
	
	"MD5 is a widely used hash function for producing hash value of",
	"64 bits", "128 bits", "512 bits", "1024 bits",
	
	"Encrypt the Message “HELLO MY DEARZ” using Transposition Cipher with key:
	Plain Text 2413
	Cipher Text 1234",
	"HLLEO YM AEDRZ", "EHOLL ZYM RAED", "ELHL MDOY AZER", "ELHL DOMY ZAER"
);

insert into mcq_details(
	question_id, 
	level, 
	question_1, 
	option_1_a, option_1_b, option_1_c, option_1_d, 
	question_2, 
	option_2_a, option_2_b, option_2_c, option_2_d, 
	question_3, 
	option_3_a, option_3_b, option_3_c, option_3_d, 
	question_4, 
	option_4_a, option_4_b, option_4_c, option_4_d) 
values(
	"MCQ_8",
	3,
	
	"DES uses a key generator to generate sixteen _______ round keys.",
	"32-bit", "48-bit", "54-bit", "42-bit",
	
	"DES has an initial and final permutation block and _________ rounds",
	"14", "15", "16", "None",
	
	"ECB and CBC are ________ ciphers.",
	"Block", "Stream", "Feild", "None",
	
	"I bank online. Which of the following are application-level encryption protocols that I would most likely use to securely bank online?",
	"SSL and SET", "Verisign and SHA1", "READY, SET, and GO", "PGP, PEM, and SSL"
);

insert into mcq_details(
	question_id, 
	level, 
	question_1, 
	option_1_a, option_1_b, option_1_c, option_1_d, 
	question_2, 
	option_2_a, option_2_b, option_2_c, option_2_d, 
	question_3, 
	option_3_a, option_3_b, option_3_c, option_3_d, 
	question_4, 
	option_4_a, option_4_b, option_4_c, option_4_d) 
values(
	"MCQ_9",
	3,
	
	"Which ports should be blocked to prevent null session enumeration?",
	"Ports 120 and 445", "Ports 1335 and 136", "Ports 110 and 137", "Ports 135 and 139",
	
	"What tool can be used to perform SNMP enumeration?",
	"DNS lookup", "Whois", "Nslookup", "IP Network Browser",
	
	"In the RSA public key cryptosystem, the private and public keys are (e, n) and (d, n) respectively, where n = p*q and p and q are large primes. Besides, n is public and p and q are private. Let M be an integer such that 0 < M < n and f(n) = (p- 1)(q-1). Now consider the following equations. 

I.  M’= Me mod n
    M = (M’)d mod n 

II.  ed is equivalent to 1 mod n 

III. ed is equivalent to 1 mod f(n)

IV. M’= Me mod f(n)
    M = (M’)d mod f(n) ",
	"I and II", "I and III", "II and IV", "III and IV",
	
	"Which database is queried by Whois?",
	"ICANN", "ARIN", "APNIC", "DNS"
);

insert into mcq_details(
	question_id, 
	level, 
	question_1, 
	option_1_a, option_1_b, option_1_c, option_1_d, 
	question_2, 
	option_2_a, option_2_b, option_2_c, option_2_d, 
	question_3, 
	option_3_a, option_3_b, option_3_c, option_3_d, 
	question_4, 
	option_4_a, option_4_b, option_4_c, option_4_d) 
values(
	"MCQ_10",
	3,
	
	"What port does Telnet use?",
	"22", "80", "20", "23",
	
	"What is the sequence of a TCP connection?",
	"SYN-ACK-FIN", "SYN-SYN ACK-ACK", "SYN-ACK", "SYN-SYN-ACK",
	
	"Which form of encryption does WPA use?",
	"Shared key", "LEAP", "TKIP", "AES",
	
	"What protocol is the Active Directory database based on?",
	"LDAP", "TCP", "SQL", "HTTP"
);