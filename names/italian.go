package names

var itaFirstNames = map[Gender][]string{
	GenderFemale: {"Adelasia", "Adele", "Adriana", "Albina", "Alessandra", "Alessia", "Alina", "Allegra", "Amalia", "Amelia", "Angelica", "Angelina", "Anita", "Annalisa", "Annamaria", "Annetta", "Annunziata", "Antonella", "Antonia", "Antonietta", "Antonina", "Aria", "Aurora", "Barbara", "Beatrice", "Berenice", "Bettina", "Bianca", "Bianca Maria", "Bridgetta", "Camilla", "Carla", "Carlina", "Carlotta", "Carolina", "Cassandra", "Caterina", "Catiuscia", "Cecilia", "Chiara", "Claudia", "Clelia", "Concetta", "Costanza", "Cristina", "Danila", "Daria", "Diana", "Diletta", "Dina", "Donatella", "Donna", "Dora", "Eleanora", "Elena", "Eliana", "Elisa", "Elisabetta", "Elvira", "Emilia", "Emma", "Enrica", "Erika", "Etta", "Eugenia", "Eva", "Evangelista", "Evelina", "Fabia", "Fabrizia", "Federica", "Fernanda", "Fiamma", "Filippa", "Fiorella", "Flavia", "Flora", "Fortunata", "Franca", "Francesca", "Gabriella", "Gelsomina", "Gemma", "Germana", "Gia", "Giada", "Gianna", "Gigliola", "Gina", "Giorgia", "Giovanna", "Gisela", "Giulia", "Giuliana", "Giulietta", "Giuseppina", "Gloria", "Graziella", "Greca", "Griselda", "Ida", "Ilaria", "Imelda", "Iolanda", "Ippolita", "Irene", "Irma", "Isa", "Isabella", "Jolanda", "Karina", "Lara", "Laura", "Lauretta", "Letizia", "Liana", "Licia", "Lidia", "Liliana", "Lilla", "Lina", "Lisa", "Livia", "Lolita", "Loredana", "Lorena", "Loretta", "Lucia", "Luciana", "Lucrezia", "Ludovica", "Luigina", "Luisa", "Marcella", "Margherita", "Maria", "María Antonietta", "Maria Sole", "Mariana", "Mariella", "Marina", "Marisa", "Marta", "Martina", "Matilda", "Maura", "Melania", "Melina", "Melissa", "Mercedes", "Michela", "Milena", "Monica", "Morena", "Nadia", "Natalia", "Nedda", "Nicoletta", "Nina", "Ninetta", "Olga", "Ornella", "Paloma", "Paola", "Paoletta", "Patrizia", "Paulina", "Pearl", "Piera", "Pierina", "Pina", "Raffaella", "Ramona", "Regina", "Renata", "Rita", "Roberta", "Romana", "Romina", "Rosa", "Rosalia", "Rosaria", "Rosina", "Rossana", "Rossella", "Serafina", "Serena", "Sibilla", "Silvia", "Simonetta", "Sonia", "Sophia", "Stefani", "Stefania", "Stella", "Susanna", "Sylvia", "Tatiana", "Teresa", "Tiziana", "Tonia", "Tonina", "Valentina", "Valeria", "Vanessa", "Vanna", "Verona", "Veronica", "Vincenza", "Virginia", "Viridiana", "Vittoria", "Zita"},
	GenderMale:   {"Achille", "Adamo", "Adelmo", "Adriano", "Agnolo", "Agostino", "Alberico", "Alberto", "Alderano", "Aldo", "Alessandro", "Alessio", "Alfio", "Alfredo", "Alphons", "Amadeo", "Amedeo", "Amleto", "Angelo", "Annibale", "Ansaldo", "Antonello", "Antonino", "Antonio", "Armando", "Arnaldo", "Arnulfo", "Arsenio", "Arturo", "Atenolfo", "Atenulfo", "Augusto", "Azeglio", "Baccio", "Baldassare", "Bartolomeo", "Benedetto", "Benito", "Benvenuto", "Beppe", "Bernardo", "Biagio", "Bruno", "Calcedonio", "Calogero", "Camillo", "Carlo", "Carlo", "Carmelo", "Carmine", "Cesare", "Cipriano", "Cirillo", "Ciro", "Claudio", "Coluccio", "Constanzo", "Coriolano", "Corrado", "Costantino", "Costanzo", "Danilo", "Damiano", "Daniele", "Daniello", "Dante", "Dario", "Davide", "Delfino", "Dino", "Dionigi", "Domenico", "Donatello", "Donato", "Durante", "Edgardo", "Edoardo", "Egidio", "Elladio", "Elmo", "Elpidio", "Emanuele", "Emilio", "Ennio", "Enrico", "Enzio", "Enzo", "Eraldo", "Ercole", "Ermanno", "Ermenegildo", "Ermes", "Ernesto", "Ettore", "Eugenio", "Ezio", "Fabio", "Fabrizio", "Fausto", "Fedele", "Federico", "Federigo", "Ferdinando", "Ferruccio", "Filippo", "Fiorenzo", "Fiorino", "Flavio", "Francesco", "Franco", "Fredo", "Fulvio", "Furio", "Gabriele", "Gaetano", "Galasso", "Gaspare", "Gastone", "Gavino", "Gennaro", "Geppetto", "Geronimo", "Giacinto", "Giacobbe", "Giacomo", "Giammaria", "Giampaolo", "Giampiero", "Gian", "Gian Carlo", "Gianantonio", "Giancarlo", "Gianfrancesco", "Gianfranco", "Gianluca", "Gianluigi", "Gianmarco", "Gianmaria", "Giannantonio", "Gianni", "Gianpaolo", "Gianpietro", "Gilberto", "Gino", "Gioacchino", "Gioachino", "Gioele", "Gioffre", "Giona", "Gionata", "Giordano", "Giorgio", "Giosuè", "Giovanni", "Giovanni Battista", "Girolamo", "Giuliano", "Giulio", "Giuseppe", "Giustino", "Goffredo", "Graziano", "Greco", "Gregorio", "Guarino", "Guglielmo", "Guido", "Gustavo", "Hugo", "Iacopo", "Ignazio", "Ippazio", "Ivo", "Jacopo", "Lamberto", "Lando", "Laureano", "Lazzaro", "Leonardo", "Leone", "Leopoldo", "Liberato", "Liberto", "Livio", "Lodovico", "Lorenzo", "Lotario", "Luca", "Luchino", "Luciano", "Lucio", "Ludovico", "Luigi", "Manlio", "Manuel", "Marcantonio", "Marcello", "Marco", "Mariano", "Mario", "Martino", "Martino", "Massimiliano", "Massimo", "Matteo", "Mattia", "Maurilio", "Maurizio", "Mauro", "Michelangelo", "Michele", "Micheletto", "Michelotto", "Milo", "Mirco", "Mirko", "Moreno", "Nanni", "Napoleone", "Niccolò", "Nico", "Nicola", "Nicolò", "Nino", "Noe", "Nunzio", "Omero", "Orazio", "Oreste", "Orlando", "Osvaldo", "Ottavio", "Ottone", "Pandulf", "Panfilo", "Paolo", "Paride", "Pasqual", "Pasquale", "Patrizio", "Pellegrino", "Peppino", "Pier", "Pierangelo", "Piergiorgio", "Piergiuseppe", "Pierluigi", "Piermaria", "Piero", "Pierpaolo", "Piersanti", "Pietro", "Pompeo", "Pomponio", "Puccio", "Quirico", "Quirino", "Raffaele", "Raffaellino", "Raffaello", "Raimondo", "Ranieri", "Renzo", "Riccardo", "Ricciotti", "Rino", "Roberto", "Rocco", "Rodolfo", "Rolando", "Roman", "Romeo", "Romolo", "Ronaldo", "Rosario", "Ruggero", "Ruggiero", "Sabatino", "Salvatore", "Salvi", "Samuele", "Sandro", "Sante", "Santino", "Saverio", "Sebastiano", "Sergius", "Severino", "Silvestro", "Silvio", "Simone", "Stefano", "Tazio", "Telemaco", "Temistocle", "Tiberio", "Tiziano", "Tommaso", "Toni", "Tonino", "Tonio", "Torquato", "Tullio", "Ubaldo", "Uberto", "Ugo", "Ugolino", "Umberto", "Valerio", "Venancio", "Vespasiano", "Vincentio", "Vincenzo", "Virgilio", "Vito", "Vittorio", "Vladimiro", "Wladimiro", "Zanobi"},
}

var itaLastNames = []string{"Abatantuono", "Abbatucci", "Acanfora", "Addobbati", "Agreiter", "Aiello", "Alderisi", "Alessandrini", "Allori", "Amitrano", "Anfossi", "Angelo", "Angioni", "Archibugi", "Ardisson", "Armenis", "Asaro", "Astorino", "Azzolina", "Balbo", "Barberis", "Barretta", "Belfi", "Bellini", "Bellodi", "Berardi", "Bergamasco", "Bernacchi", "Bernardi", "Bernardini", "Bertacco", "Bertolami", "Bertolin", "Bertoloni", "Besozzi", "Biagetti", "Bilotti", "Bixio", "Boaron", "Boffa", "Bollani", "Bonaparte", "Bongiovanni", "Bontade", "Bonventre", "Borghi", "Borosini", "Borromini", "Boselli", "Bosio", "Bottero", "Botticelli", "Bottici", "Brambilla", "Buffon", "Butera", "Caimi", "Camerota", "Camilleri", "Campioni", "Canaveri", "Cangini", "Cantù", "Capano", "Cardillo", "Carlucci", "Carpena", "Cascarino", "Casciaro", "Casellati", "Castelfranchi", "Castellani", "Castelletti", "Castello", "Cattaneo", "Cavani", "Cavanna", "Ceccarelli", "Cella", "Cellario", "Cenci", "Cercato", "Chiaraviglio", "Chimento", "Ciaccia", "Cocchi", "Colombo", "Colosimo", "Colucci", "Comolli", "Comollo", "Conti", "Coscarelli", "Curione", "Cusanelli", "D'Ambrosi", "Da Milano", "De Carlo", "De Filippo", "Del Colombo", "Delgaudio", "Dell'Orco", "DeLuca", "Destro", "DiCenzo", "Duva", "Este", "Fabbiano", "Fabbri", "Fabretti", "Facchin", "Fantozzi", "Fanucchi", "Farentino", "Farrugia", "Fazzari", "Ferrara", "Ferretti", "Flaim", "Fornasini", "Franceschini", "Franco", "Franzetti", "Fratangelo", "Fresu", "Fumagalli", "Furlan", "Gaipa", "Galizia", "Gallone", "Gallucci", "Garguilo", "Gasparetto", "Gaytán", "Gazzarri", "Gesualdi", "Gherardi", "Ghezzi", "Ghiotto", "Giambalvo", "Giambi", "Giarrusso", "Ginetti", "Giusti", "Gnoli", "Grella", "Gronchi", "Guzzoni", "Iacocca", "Innocenti", "Kostner", "La Malfa", "Lamo", "Lamorgese", "Landucci", "Lanthaler", "Latini", "Lauria", "Ledda", "Lercaro", "Licinio", "Limardo", "Lo Bue", "Locatelli", "Loi", "Lovato", "Lucchi", "Lucchini", "Lupinacci", "Lupino", "Lupo", "Maccioni", "Maestrini", "Maffei", "Maffeis", "Maffeo", "Maggi", "Maggiacomo", "Maggio", "Maglia", "Magliani", "Magnasco", "Magnetto", "Magnoli", "Magrassi", "Maiellaro", "Maisano", "Malagola", "Malanima", "Malnati", "Manardi", "Mancini", "Mancino", "Mancuso", "Mandarino", "Manganello", "Maniscalco", "Mannoia", "Mantia", "Manzato", "Manzoni", "Marafioti", "Marcolin", "Marcone", "Martini", "Martiradonna", "Mascagni", "Maselli", "Masini", "Massucci", "Masullo", "Mazzi", "Menga", "Mengarini", "Mingori", "Minuto", "Modigliani", "Molinelli", "Monfalcone", "Mongin", "Moroder", "Morreale", "Moscon", "Moser", "Murgia", "Nazarro", "Nazzaro", "Neroni", "Nibali", "Nisini", "Nocentini", "Nullo", "Onorati", "Paganelli", "Paletta", "Pellegrini", "Pericoli", "Pestalozzi", "Petrenga", "Petrina", "Picano", "Pierucci", "Pirovano", "Pisacane", "Poggi", "Poggio", "Poletto", "Poli", "Pompili", "Rappa", "Ravaioli", "Ravasi", "Regge", "Rendine", "Renzulli", "Ricciulli", "Ripamonti", "Riva", "Rizzo", "Rodighiero", "Romiti", "Rosati", "Rosso", "Ruotolo", "Sabbadini", "Saccone", "Salvadori", "Sansone", "Santangelo", "Saputo", "Sarnelli", "Sberna", "Schirinzi", "Sciortino", "Seccafien", "Setario", "Severi", "Simari", "Simbari", "Sirico", "Sirigu", "Siviter", "Spannocchi", "Sposato", "Stagnaro", "Stango", "Stefano", "Stramaccioni", "Suppa", "Tantillo", "Tebaldi", "Terry", "Tiso", "Tonetto", "Tonini", "Tortorich", "Tozzi", "Trafeli", "Trentanove", "Triscari", "Trivulzio", "Trombetti", "Turati", "Uberti", "Vagnozzi", "Valeriote", "Vassalli", "Veggio", "Veneziano", "Veneziano", "Vezzani", "Viceconte", "Vitelleschi", "Zago", "Zenetti", "Zola"}
