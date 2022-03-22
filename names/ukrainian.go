package names

var ukrFirstNames = map[Gender][]string{
	GenderFemale: {"Alona", "Anastasia", "Angelika", "Antonina", "Darina", "Diana", "Draga", "Ekaterina", "Elena", "Elina", "Galina", "Ganna", "Gorana", "Inna", "Irina", "Katerina", "Katya", "Larissa", "Ludmila", "Lyubov", "Maria", "Marina", "Melania", "Mila", "Milena", "Milica", "Mira", "Nadezhda", "Nadia", "Nadiia", "Nadiya", "Natalia", "Nataliya", "Natalka", "Natasha", "Nevena", "Nikolina", "Oksana", "Olena", "Olga", "Ruslana", "Snežana", "Stefania", "Svitlana", "Valeria", "Veronica", "Vesna", "Yelena", "Yulia", "Zora"},
	GenderMale:   {"Adam", "Anatoly", "Andrei", "Andrey", "Andriy", "Anton", "Artem", "Artjoms", "Artyom", "Berislav", "Blagoje", "Bogdan", "Boris", "Borys", "Boško", "Branimir", "Davor", "Dmitry", "Dmytro", "Dragan", "Dragoljub", "Dragomir", "Dragoslav", "Duško", "Gleb", "Hennadiy", "Hryhoriy", "Ilarion", "Ivan", "Jadranko", "Jaromir", "Jaroslav", "Kliment", "Kostyantyn", "Kyrylo", "Leonid", "Lubomir", "Maksym", "Maxim", "Milovan", "Miroslav", "Mykhailo", "Mykola", "Mykyta", "Myroslav", "Nectarios", "Nenad", "Neven", "Nikita", "Nikola", "Oleg", "Oleksiy", "Orest", "Ostap", "Panteleimon", "Pavsikakiy", "Petro", "Ratimir", "Roman", "Rostislav", "Ruslan", "Sergiy", "Severyn", "Sobieslaw", "Stanimir", "Stepan", "Sviatoslav", "Symon", "Taras", "Vadim", "Vadym", "Vasyl", "Velimir", "Veselin", "Viacheslav", "Victor", "Vitalii", "Vitomir", "Vladimir", "Vladislav", "Vladyslav", "Volodymyr", "Vsevolod", "Vukašin", "Vukmir", "Yevgen", "Yevhen", "Yukhym", "Yury", "Željko", "Zinovy", "Zoran"},
}

var ukrLastNames = []string{"Abramenko", "Achtymichuk", "Adamchuk", "Adamenko", "Adamovsky", "Afanasenko", "Akimenko", "Aleksenko", "Alekseyenko", "Andreychuk", "Andriychuk", "Andryushchenko", "Antonenko", "Arkhypenko", "Artemenko", "Avramenko", "Babchenko", "Babenko", "Bakay", "Baranovskyi", "Barilko", "Bazylevych", "Belanov", "Berezhna", "Berezovsky", "Berkovic", "Berkovich", "Berkovits", "Bezugly", "Bilodid", "Bilous", "Bily", "Bilyi", "Bilyk", "Bliznyuk", "Blonsky", "Bobyr", "Bohdanov", "Boki", "Bondar", "Bondarchuk", "Bondarenko", "Borodai", "Bortnik", "Borysenko", "Borysyuk", "Boychenko", "Boychuk", "Boyko", "Brezhnev", "Brodsky", "Brody", "Brusilovsky", "Bublik", "Burachek", "Burliuk", "Butenko", "Buteyko", "Butko", "Bykovsky", "Chalupa", "Chepiga", "Chepiha", "Chernenko", "Chernetsky", "Cherniavsky", "Chernyak", "Chervonenko", "Chorny", "Chumachenko", "Companeez", "Danchenko", "Danylchenko", "Danylenko", "Danyluk", "Danylyuk", "Demchak", "Demchenko", "Demchuk", "Demyanenko", "Denysenko", "Derevyanko", "Derkach", "Didukh", "Dikhtiar", "Dmytrenko", "Donchenko", "Doroshenko", "Dotsenko", "Dovhal", "Dovhan", "Dovzhenko", "Drozd", "Dudchenko", "Dudka", "Dushenko", "Dutko", "Dyachenko", "Dziuba", "Dzyadko", "Eremenko", "Fedorchuk", "Fedorenko", "Fedoruk", "Fesenko", "Filippenko", "Fomenko", "Franchuk", "Gamarnik", "Garmash", "Gerasimenko", "Glavan", "Gnidenko", "Gogotsi", "Goloborodko", "Golovko", "Golubitsky", "Grigorenko", "Grishchuk", "Gritsenko", "Gulko", "Gunko", "Hapon", "Harkusha", "Henyk", "Hirnyk", "Hlavan", "Hlovatskyi", "Hlushchenko", "Hnatyuk", "Hodlevskyi", "Hohol", "Holovatskyi", "Holovchenko", "Holowaty", "Honchar", "Honcharenko", "Hordiyenko", "Horishnyi", "Horobets", "Hrabovskyi", "Hrozný", "Hrynko", "Hrytsak", "Hudyma", "Humenyuk", "Husiev", "Huzenko", "Ianchuk", "Ilchenko", "Ilnytskyi", "Ishchenko", "Ivakhnenko", "Ivanchuk", "Ivanenko", "Ivanik", "Ivanyuk", "Ivashchenko", "Ivashko", "Ivasyuk", "Jakovina", "Kachur", "Kalachnik", "Kalachnyk", "Kalačnik", "Kalashnik", "Kalashnyk", "Kalasnik", "Kalašnik", "Kalynovskyi", "Kalyuzhny", "Kaminsky", "Kapinos", "Karpenko", "Kashchenko", "Kharchenko", "Khmara", "Khmelnytskyi", "Khomenko", "Khristenko", "Khvostenko", "Kichenok", "Kirichenko", "Kirilenko", "Kleshchenko", "Klochko", "Klymenko", "Klympush", "Knizhnik", "Kobets", "Kochubey", "Kohut", "Kolachnik", "Kolachnyk", "Kolashnik", "Kolashnyk", "Kolasnik", "Kolessa", "Kolyada", "Komisaruk", "Kondracki", "Kondratenko", "Kondratyuk", "Kononenko", "Konoval", "Konovalchuk", "Konovalenko", "Konovalyuk", "Korchmar", "Korniyenko", "Korolenko", "Kosenko", "Kostenko", "Kostiuk", "Kostyuk", "Kotlyarevsky", "Kotyk", "Koval", "Kovalchuk", "Kovalenko", "Kovalevich", "Kovalska", "Kovalyuk", "Kovpak", "Kovtun", "Kozhanov", "Kozub", "Kramarenko", "Krasko", "Krasovsky", "Kravchenko", "Kravchuk", "Kravets", "Kravtsov", "Kreviazuk", "Krichevsky", "Kruk", "Krushelnytskyi", "Kryvoruchko", "Kucherenko", "Kudryk", "Kukhar", "Kukharenko", "Kulesha", "Kulinich", "Kulish", "Kulyk", "Kurchenko", "Kurkudym", "Kurylo", "Kuryluk", "Kushnir", "Kutsenko", "Kuzmenko", "Kvasha", "Kvitko", "Lavrinenko", "Lazarenko", "Lebedenko", "Lebid", "Leshchenko", "Levchenko", "Levytsky", "Litvinchuk", "Loboda", "Lopata", "Los", "Luchko", "Lukashenko", "Lukyanenko", "Lutsenko", "Lyashko", "Lyko", "Lysenko", "Lytovchenko", "Lytvyn", "Lytvynenko", "Lytwyn", "Maidanik", "Makara", "Makarenko", "Makhno", "Maksimenko", "Maksyuta", "Makuch", "Malkovich", "Malkovych", "Malyi", "Marchenko", "Marchuk", "Markevych", "Martyniuk", "Marusiak", "Matios", "Matviychuk", "Matviyenko", "Mazurenko", "Melnik", "Melnychenko", "Melnychuk", "Michayluk", "Mikhalchenko", "Mikhaylovsky", "Minenko", "Miroshnichenko", "Miroshnyk", "Miroshnykov", "Mishchenko", "Mitnick", "Molchan", "Momot", "Moroz", "Moskalenko", "Movchan", "Mudry", "Murashko", "Muzychuk", "Mykhaylenko", "Mykhaylychenko", "Mykolenko", "Mysak", "Nad", "Nakonechny", "Naumenko", "Nazarenko", "Nedbaylo", "Nedelko", "Nesterenko", "Netrebko", "Niemirowski", "Nikolayenko", "Nosenko", "Nosko", "Novosad", "Nykoluk", "Okopenko", "Olech", "Olekh", "Oliynyk", "Olynyk", "Omelchenko", "Onischuk", "Onishchenko", "Onopko", "Onoprienko", "Osadchuk", "Ostapchuk", "Ostapenko", "Palahniuk", "Palamarchuk", "Palatnik", "Panchenko", "Pankevych", "Parasyuk", "Pashchenko", "Paskevych", "Pasko", "Pavlenko", "Pavlichenko", "Pavliuchenko", "Pavliuk", "Pavlychko", "Perebiynis", "Perepadenko", "Petko", "Petrenko", "Piddubny", "Pinchuk", "Pisarenko", "Piskun", "Piven", "Pluta", "Podolski", "Pogrebinsky", "Polishchuk", "Polonsky", "Polyakov", "Pomazan", "Ponomarenko", "Potapenko", "Povkh", "Prigozhin", "Prikhodko", "Prokhorenko", "Prokopenko", "Prokopovich", "Prokopovych", "Protsenko", "Pryma", "Prystay", "Prysyazhnyuk", "Prytula", "Radchenko", "Radivilov", "Reznichenko", "Rokita", "Romanchuk", "Romanenko", "Romaniuk", "Romaschenko", "Rudenko", "Rudnitsky", "Rybachenko", "Rybachuk", "Rybalchenko", "Rybalka", "Rybalko", "Sakhno", "Saladukha", "Samoilenko", "Sapunov", "Satschko", "Savaryn", "Savchenko", "Savchuk", "Savenko", "Sayenko", "Sembratovych", "Semenenko", "Semenko", "Semerenko", "Senchenko", "Sergiyenko", "Shaposhnik", "Shapoval", "Sharko", "Shcherba", "Shcherbak", "Shcherbina", "Shelest", "Shepilenko", "Sheremeta", "Shevchenko", "Shevchuk", "Shmatko", "Shostak", "Shvets", "Shymko", "Shynkarenko", "Sibiryakov", "Sidorenko", "Simović", "Sirota", "Sklyarenko", "Skok", "Skornyakov", "Skrypnyk", "Slissenko", "Slobodenyuk", "Slyusar", "Sobko", "Sokolenko", "Soroko", "Sosenko", "Sosna", "Sova", "Spivak", "Stadnik", "Stakhovsky", "Stanev", "Stelmakh", "Stepanenko", "Stepanyuk", "Stępniak", "Stetsenko", "Stetsko", "Stocki", "Stolyar", "Stolyarchuk", "Stolyarenko", "Storchak", "Storozhenko", "Stupak", "Stupka", "Suprun", "Suprunenko", "Sviatchenko", "Sydor", "Sydorenko", "Sytnyk", "Sytnykov", "Taranenko", "Tarasenko", "Tereshchenko", "Tereshchuk", "Teslya", "Tishchenko", "Titarenko", "Tiutiunnyk", "Tkach", "Tkachenko", "Tkachuk", "Tlumak", "Tokarczuk", "Tolpygo", "Trautmann", "Tretiak", "Trofimenko", "Trutovsky", "Tsybulenko", "Tsymbal", "Twersky", "Tymchenko", "Tymchuk", "Tymoshenko", "Ulyanenko", "Umansky", "Usenko", "Vakarchuk", "Vakulenko", "Vashchuk", "Vasilchenko", "Vasilenko", "Vasin", "Vdovichenko", "Vengerov", "Venherov", "Virastyuk", "Vishnevetsky", "Vlasenko", "Voitiuk", "Volesky", "Volokh", "Voloshchenko", "Voloshin", "Volynets", "Vorobey", "Vorona", "Vovk", "Voytenko", "Warhola", "Yakhno", "Yakovenko", "Yaremchuk", "Yaremenko", "Yaroshchuk", "Yaroshenko", "Yarovenko", "Yarovyi", "Yashchenko", "Yasko", "Yastremska", "Yatsenko", "Yeroshenko", "Yevtushenko", "Yurchenko", "Yurin", "Yurkov", "Yushchenko", "Zagrebelny", "Zahorulko", "Zakaluzny", "Zakharchenko", "Zakharenko", "Zaliznyak", "Zanevska", "Zawadzki", "Zayets", "Zelenko", "Zerebko", "Zhabotinsky", "Zhovtyak", "Zhuk", "Zhukovsky", "Zoshchenko", "Zotov", "Zozulya", "Zubko", "Zvarych"}
