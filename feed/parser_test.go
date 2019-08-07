package feed

import "testing"

func TestParse(t *testing.T) {

	feed := `<?xml version="1.0" encoding="ISO-8859-1"?>
    <rss version="2.0">
    <channel>
    <title>Pelastustoimen mediapalvelu: ensitiedotteet</title>
    <link>http://www.peto-media.fi</link>
    <description>Ensitiedotteita Suomesta....</description>
    <language>FI-fi</language>
    <item>
    <title>Liperi/Liperi, öljyvah./ymp.onnet. maalla:  pieni</title>
    <description>07.08.2019 17:17:04 Liperi/Liperi öljyvah./ymp.onnet. maalla:  pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 17:17:04 +0200</pubDate>
    </item>
    <item>
    <title>Mustasaari/Korsholm, maastopalo: pieni</title>
    <description>07.08.2019 17:13:04 Mustasaari/Korsholm maastopalo: pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 17:13:04 +0200</pubDate>
    </item>
    <item>
    <title>Espoo/Esbo, tieliikenneonnettomuus: pieni</title>
    <description>07.08.2019 17:12:16 Espoo/Esbo tieliikenneonnettomuus: pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 17:12:16 +0200</pubDate>
    </item>
    <item>
    <title>Salo/Salo, tieliikenneonnettomuus: pieni</title>
    <description>07.08.2019 17:06:26 Salo/Salo tieliikenneonnettomuus: pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 17:06:26 +0200</pubDate>
    </item>
    <item>
    <title>Turku/Åbo, tieliikenneonnettomuus: pieni</title>
    <description>07.08.2019 17:04:43 Turku/Åbo tieliikenneonnettomuus: pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 17:04:43 +0200</pubDate>
    </item>
    <item>
    <title>Helsinki/Helsingfors, palohälytys</title>
    <description>07.08.2019 17:03:39 Helsinki/Helsingfors palohälytys</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 17:03:39 +0200</pubDate>
    </item>
    <item>
    <title>Leppävirta/Leppävirta, savuhavainto</title>
    <description>07.08.2019 17:00:20 Leppävirta/Leppävirta savuhavainto</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 17:00:20 +0200</pubDate>
    </item>
    <item>
    <title>Kemiönsaari/Kimitoön, palohälytys</title>
    <description>07.08.2019 17:00:03 Kemiönsaari/Kimitoön palohälytys</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 17:00:03 +0200</pubDate>
    </item>
    <item>
    <title>Espoo/Esbo, tieliikenneonnettomuus: pieni</title>
    <description>07.08.2019 16:49:54 Espoo/Esbo tieliikenneonnettomuus: pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 16:49:54 +0200</pubDate>
    </item>
    <item>
    <title>Espoo/Esbo, tieliikenneonnettomuus: keskisuuri</title>
    <description>07.08.2019 16:48:52 Espoo/Esbo tieliikenneonnettomuus: keskisuuri</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 16:48:52 +0200</pubDate>
    </item>
    <item>
    <title>Helsinki/Helsingfors, rakennuspalo: pieni</title>
    <description>07.08.2019 16:48:35 Helsinki/Helsingfors rakennuspalo: pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 16:48:35 +0200</pubDate>
    </item>
    <item>
    <title>Harjavalta/Harjavalta, palohälytys</title>
    <description>07.08.2019 16:47:27 Harjavalta/Harjavalta palohälytys</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 16:47:27 +0200</pubDate>
    </item>
    <item>
    <title>Kotka/Kotka, öljyvah./ymp.onnet. maalla:  pieni</title>
    <description>07.08.2019 16:42:52 Kotka/Kotka öljyvah./ymp.onnet. maalla:  pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 16:42:52 +0200</pubDate>
    </item>
    <item>
    <title>Joensuu/Joensuu, tieliikenneonnettomuus: pieni</title>
    <description>07.08.2019 16:41:37 Joensuu/Joensuu tieliikenneonnettomuus: pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 16:41:37 +0200</pubDate>
    </item>
    <item>
    <title>Ranua/Ranua, tieliikenneonnettomuus: pieni</title>
    <description>07.08.2019 16:40:44 Ranua/Ranua tieliikenneonnettomuus: pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 16:40:44 +0200</pubDate>
    </item>
    <item>
    <title>Helsinki/Helsingfors, tieliikenneonnettomuus: pieni</title>
    <description>07.08.2019 16:36:42 Helsinki/Helsingfors tieliikenneonnettomuus: pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 16:36:42 +0200</pubDate>
    </item>
    <item>
    <title>Hollola/Hollola, öljyvah./ymp.onnet. maalla:  pieni</title>
    <description>07.08.2019 16:32:42 Hollola/Hollola öljyvah./ymp.onnet. maalla:  pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 16:32:42 +0200</pubDate>
    </item>
    <item>
    <title>Alavieska/Alavieska, tieliikenneonnettomuus: pieni</title>
    <description>07.08.2019 16:30:05 Alavieska/Alavieska tieliikenneonnettomuus: pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 16:30:05 +0200</pubDate>
    </item>
    <item>
    <title>Kauhajoki/Kauhajoki, rakennuspalo: pieni</title>
    <description>07.08.2019 16:28:00 Kauhajoki/Kauhajoki rakennuspalo: pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 16:28:00 +0200</pubDate>
    </item>
    <item>
    <title>Tampere/Tammerfors, rakennuspalo: pieni</title>
    <description>07.08.2019 16:22:31 Tampere/Tammerfors rakennuspalo: pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 16:22:31 +0200</pubDate>
    </item>
    <item>
    <title>Iisalmi/Idensalmi, tulipalo muu: pieni</title>
    <description>07.08.2019 16:20:50 Iisalmi/Idensalmi tulipalo muu: pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 16:20:50 +0200</pubDate>
    </item>
    <item>
    <title>Uusikaarlepyy/Nykarleby, tulipalo muu: keskisuuri</title>
    <description>07.08.2019 16:19:10 Uusikaarlepyy/Nykarleby tulipalo muu: keskisuuri</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 16:19:10 +0200</pubDate>
    </item>
    <item>
    <title>Kemiönsaari/Kimitoön, rakennuspalo: pieni</title>
    <description>07.08.2019 16:17:50 Kemiönsaari/Kimitoön rakennuspalo: pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 16:17:50 +0200</pubDate>
    </item>
    <item>
    <title>Laihia/Laihela, maastopalo: pieni</title>
    <description>07.08.2019 16:15:52 Laihia/Laihela maastopalo: pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 16:15:52 +0200</pubDate>
    </item>
    <item>
    <title>Eura/Eura, maastopalo: keskisuuri</title>
    <description>07.08.2019 16:15:10 Eura/Eura maastopalo: keskisuuri</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 16:15:10 +0200</pubDate>
    </item>
    <item>
    <title>Sotkamo/Sotkamo, palohälytys</title>
    <description>07.08.2019 16:12:56 Sotkamo/Sotkamo palohälytys</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 16:12:56 +0200</pubDate>
    </item>
    <item>
    <title>Kirkkonummi/Kyrkslätt, palohälytys</title>
    <description>07.08.2019 16:10:31 Kirkkonummi/Kyrkslätt palohälytys</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 16:10:31 +0200</pubDate>
    </item>
    <item>
    <title>Liperi/Liperi, öljyvah./ymp.onnet. maalla:  pieni</title>
    <description>07.08.2019 16:03:38 Liperi/Liperi öljyvah./ymp.onnet. maalla:  pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 16:03:38 +0200</pubDate>
    </item>
    <item>
    <title>Kuopio/Kuopio, savuhavainto</title>
    <description>07.08.2019 15:59:05 Kuopio/Kuopio savuhavainto</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 15:59:05 +0200</pubDate>
    </item>
    <item>
    <title>Heinävesi/Heinävesi, savuhavainto</title>
    <description>07.08.2019 15:54:42 Heinävesi/Heinävesi savuhavainto</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 15:54:42 +0200</pubDate>
    </item>
    <item>
    <title>Mäntyharju/Mäntyharju, maastopalo: pieni</title>
    <description>07.08.2019 15:49:19 Mäntyharju/Mäntyharju maastopalo: pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 15:49:19 +0200</pubDate>
    </item>
    <item>
    <title>Joensuu/Joensuu, savuhavainto</title>
    <description>07.08.2019 15:48:28 Joensuu/Joensuu savuhavainto</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 15:48:28 +0200</pubDate>
    </item>
    <item>
    <title>Laihia/Laihela, tieliikenneonnettomuus: pieni</title>
    <description>07.08.2019 15:42:34 Laihia/Laihela tieliikenneonnettomuus: pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 15:42:34 +0200</pubDate>
    </item>
    <item>
    <title>Heinävesi/Heinävesi, maastopalo: pieni</title>
    <description>07.08.2019 15:40:33 Heinävesi/Heinävesi maastopalo: pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 15:40:33 +0200</pubDate>
    </item>
    <item>
    <title>Alajärvi/Alajärvi, palohälytys</title>
    <description>07.08.2019 15:34:54 Alajärvi/Alajärvi palohälytys</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 15:34:54 +0200</pubDate>
    </item>
    <item>
    <title>Lahti/Lahtis, öljyvah./ymp.onnet. maalla:  pieni</title>
    <description>07.08.2019 15:30:15 Lahti/Lahtis öljyvah./ymp.onnet. maalla:  pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 15:30:15 +0200</pubDate>
    </item>
    <item>
    <title>Asikkala/Asikkala, savuhavainto</title>
    <description>07.08.2019 15:27:54 Asikkala/Asikkala savuhavainto</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 15:27:54 +0200</pubDate>
    </item>
    <item>
    <title>Salo/Salo, tieliikenneonnettomuus: pieni</title>
    <description>07.08.2019 15:21:26 Salo/Salo tieliikenneonnettomuus: pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 15:21:26 +0200</pubDate>
    </item>
    <item>
    <title>Nousiainen/Nousis, maastopalo: pieni</title>
    <description>07.08.2019 15:06:50 Nousiainen/Nousis maastopalo: pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 15:06:50 +0200</pubDate>
    </item>
    <item>
    <title>Nurmes/Nurmes, tieliikenneonnettomuus: pieni</title>
    <description>07.08.2019 15:06:45 Nurmes/Nurmes tieliikenneonnettomuus: pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 15:06:45 +0200</pubDate>
    </item>
    <item>
    <title>Oulu/Uleåborg, maastopalo: pieni</title>
    <description>07.08.2019 15:05:52 Oulu/Uleåborg maastopalo: pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 15:05:52 +0200</pubDate>
    </item>
    <item>
    <title>Nivala/Nivala, rakennuspalo: pieni</title>
    <description>07.08.2019 14:45:32 Nivala/Nivala rakennuspalo: pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 14:45:32 +0200</pubDate>
    </item>
    <item>
    <title>Oulu/Uleåborg, rakennuspalo: pieni</title>
    <description>07.08.2019 14:39:17 Oulu/Uleåborg rakennuspalo: pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 14:39:17 +0200</pubDate>
    </item>
    <item>
    <title>Helsinki/Helsingfors, palohälytys</title>
    <description>07.08.2019 14:35:33 Helsinki/Helsingfors palohälytys</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 14:35:33 +0200</pubDate>
    </item>
    <item>
    <title>Loppi/Loppi, palohälytys</title>
    <description>07.08.2019 14:34:32 Loppi/Loppi palohälytys</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 14:34:32 +0200</pubDate>
    </item>
    <item>
    <title>Turku/Åbo, palohälytys</title>
    <description>07.08.2019 14:29:27 Turku/Åbo palohälytys</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 14:29:27 +0200</pubDate>
    </item>
    <item>
    <title>Kajaani/Kajana, rakennuspalo: pieni</title>
    <description>07.08.2019 14:26:17 Kajaani/Kajana rakennuspalo: pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 14:26:17 +0200</pubDate>
    </item>
    <item>
    <title>Oulu/Uleåborg, palohälytys</title>
    <description>07.08.2019 14:24:39 Oulu/Uleåborg palohälytys</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 14:24:39 +0200</pubDate>
    </item>
    <item>
    <title>Leppävirta/Leppävirta, maastopalo: pieni</title>
    <description>07.08.2019 14:24:30 Leppävirta/Leppävirta maastopalo: pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 14:24:30 +0200</pubDate>
    </item>
    <item>
    <title>Kuopio/Kuopio, maastopalo: pieni</title>
    <description>07.08.2019 14:20:02 Kuopio/Kuopio maastopalo: pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 14:20:02 +0200</pubDate>
    </item>
    <item>
    <title>Hamina/Fredrikshamn, tieliikenneonnettomuus: pieni</title>
    <description>07.08.2019 14:11:49 Hamina/Fredrikshamn tieliikenneonnettomuus: pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 14:11:49 +0200</pubDate>
    </item>
    <item>
    <title>Padasjoki/Padasjoki, maastopalo: keskisuuri</title>
    <description>07.08.2019 14:08:07 Padasjoki/Padasjoki maastopalo: keskisuuri</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 14:08:07 +0200</pubDate>
    </item>
    <item>
    <title>Suonenjoki/Suonenjoki, maastopalo: pieni</title>
    <description>07.08.2019 14:03:46 Suonenjoki/Suonenjoki maastopalo: pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 14:03:46 +0200</pubDate>
    </item>
    <item>
    <title>Espoo/Esbo, palohälytys</title>
    <description>07.08.2019 14:00:53 Espoo/Esbo palohälytys</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 14:00:53 +0200</pubDate>
    </item>
    <item>
    <title>Helsinki/Helsingfors, öljyvah./ymp.onnet. maalla:  pieni</title>
    <description>07.08.2019 13:59:58 Helsinki/Helsingfors öljyvah./ymp.onnet. maalla:  pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 13:59:58 +0200</pubDate>
    </item>
    <item>
    <title>Kerava/Kervo, tieliikenneonnettomuus: pieni</title>
    <description>07.08.2019 13:56:49 Kerava/Kervo tieliikenneonnettomuus: pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 13:56:49 +0200</pubDate>
    </item>
    <item>
    <title>Joensuu/Joensuu, palohälytys</title>
    <description>07.08.2019 13:40:21 Joensuu/Joensuu palohälytys</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 13:40:21 +0200</pubDate>
    </item>
    <item>
    <title>Helsinki/Helsingfors, tieliikenneonnettomuus: pieni</title>
    <description>07.08.2019 13:39:37 Helsinki/Helsingfors tieliikenneonnettomuus: pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 13:39:37 +0200</pubDate>
    </item>
    <item>
    <title>Kajaani/Kajana, tieliikenneonnettomuus: keskisuuri</title>
    <description>07.08.2019 13:37:49 Kajaani/Kajana tieliikenneonnettomuus: keskisuuri</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 13:37:49 +0200</pubDate>
    </item>
    <item>
    <title>Kangasala/Kangasala, tieliikenneonnettomuus: pieni</title>
    <description>07.08.2019 13:31:58 Kangasala/Kangasala tieliikenneonnettomuus: pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 13:31:58 +0200</pubDate>
    </item>
    <item>
    <title>Suonenjoki/Suonenjoki, savuhavainto</title>
    <description>07.08.2019 13:31:26 Suonenjoki/Suonenjoki savuhavainto</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 13:31:26 +0200</pubDate>
    </item>
    <item>
    <title>Helsinki/Helsingfors, palohälytys</title>
    <description>07.08.2019 13:29:01 Helsinki/Helsingfors palohälytys</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 13:29:01 +0200</pubDate>
    </item>
    <item>
    <title>Rauma/Raumo, palohälytys</title>
    <description>07.08.2019 13:25:19 Rauma/Raumo palohälytys</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 13:25:19 +0200</pubDate>
    </item>
    <item>
    <title>Kauhava/Kauhava, tieliikenneonnettomuus: pieni</title>
    <description>07.08.2019 13:23:33 Kauhava/Kauhava tieliikenneonnettomuus: pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 13:23:33 +0200</pubDate>
    </item>
    <item>
    <title>Kangasala/Kangasala, vahingontorjunta: pieni</title>
    <description>07.08.2019 13:22:47 Kangasala/Kangasala vahingontorjunta: pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 13:22:47 +0200</pubDate>
    </item>
    <item>
    <title>Kurikka/Kurikka, turvetuotantoalue palo: keskisuuri</title>
    <description>07.08.2019 13:22:33 Kurikka/Kurikka turvetuotantoalue palo: keskisuuri</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 13:22:33 +0200</pubDate>
    </item>
    <item>
    <title>Virrat/Virdois, tieliikenneonnettomuus: pieni</title>
    <description>07.08.2019 13:22:05 Virrat/Virdois tieliikenneonnettomuus: pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 13:22:05 +0200</pubDate>
    </item>
    <item>
    <title>Kerava/Kervo, palohälytys</title>
    <description>07.08.2019 13:19:08 Kerava/Kervo palohälytys</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 13:19:08 +0200</pubDate>
    </item>
    <item>
    <title>Turku/Åbo, palohälytys</title>
    <description>07.08.2019 13:18:52 Turku/Åbo palohälytys</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 13:18:52 +0200</pubDate>
    </item>
    <item>
    <title>Tampere/Tammerfors, rakennuspalo: pieni</title>
    <description>07.08.2019 13:17:37 Tampere/Tammerfors rakennuspalo: pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 13:17:37 +0200</pubDate>
    </item>
    <item>
    <title>Tampere/Tammerfors, palohälytys</title>
    <description>07.08.2019 13:16:51 Tampere/Tammerfors palohälytys</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 13:16:51 +0200</pubDate>
    </item>
    <item>
    <title>Keuruu/Keuruu, palohälytys</title>
    <description>07.08.2019 13:05:38 Keuruu/Keuruu palohälytys</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 13:05:38 +0200</pubDate>
    </item>
    <item>
    <title>Hämeenlinna/Tavastehus, öljyvah./ymp.onnet. maalla:  pieni</title>
    <description>07.08.2019 13:02:24 Hämeenlinna/Tavastehus öljyvah./ymp.onnet. maalla:  pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 13:02:24 +0200</pubDate>
    </item>
    <item>
    <title>Rantasalmi/Rantasalmi, maastopalo: pieni</title>
    <description>07.08.2019 12:59:44 Rantasalmi/Rantasalmi maastopalo: pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 12:59:44 +0200</pubDate>
    </item>
    <item>
    <title>Tampere/Tammerfors, palohälytys</title>
    <description>07.08.2019 12:59:21 Tampere/Tammerfors palohälytys</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 12:59:21 +0200</pubDate>
    </item>
    <item>
    <title>Liminka/Limingo, maastopalo: pieni</title>
    <description>07.08.2019 12:58:51 Liminka/Limingo maastopalo: pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 12:58:51 +0200</pubDate>
    </item>
    <item>
    <title>Loimaa/Loimaa, palohälytys</title>
    <description>07.08.2019 12:48:55 Loimaa/Loimaa palohälytys</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 12:48:55 +0200</pubDate>
    </item>
    <item>
    <title>Joutsa/Joutsa, liikennevälinepalo: pieni</title>
    <description>07.08.2019 12:48:11 Joutsa/Joutsa liikennevälinepalo: pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 12:48:11 +0200</pubDate>
    </item>
    <item>
    <title>Lahti/Lahtis, tulipalo muu: pieni</title>
    <description>07.08.2019 12:44:45 Lahti/Lahtis tulipalo muu: pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 12:44:45 +0200</pubDate>
    </item>
    <item>
    <title>Tampere/Tammerfors, tieliikenneonnettomuus: pieni</title>
    <description>07.08.2019 12:42:34 Tampere/Tammerfors tieliikenneonnettomuus: pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 12:42:34 +0200</pubDate>
    </item>
    <item>
    <title>Salo/Salo, palohälytys</title>
    <description>07.08.2019 12:38:02 Salo/Salo palohälytys</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 12:38:02 +0200</pubDate>
    </item>
    <item>
    <title>Lahti/Lahtis, maastopalo: pieni</title>
    <description>07.08.2019 12:33:20 Lahti/Lahtis maastopalo: pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 12:33:20 +0200</pubDate>
    </item>
    <item>
    <title>Lahti/Lahtis, rakennuspalo: keskisuuri</title>
    <description>07.08.2019 12:28:56 Lahti/Lahtis rakennuspalo: keskisuuri</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 12:28:56 +0200</pubDate>
    </item>
    <item>
    <title>Lahti/Lahtis, rakennuspalo: keskisuuri</title>
    <description>07.08.2019 12:28:39 Lahti/Lahtis rakennuspalo: keskisuuri</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 12:28:39 +0200</pubDate>
    </item>
    <item>
    <title>Turku/Åbo, palohälytys</title>
    <description>07.08.2019 12:26:20 Turku/Åbo palohälytys</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 12:26:20 +0200</pubDate>
    </item>
    <item>
    <title>Lahti/Lahtis, ihmisen pelastaminen, muu</title>
    <description>07.08.2019 12:25:01 Lahti/Lahtis ihmisen pelastaminen, muu</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 12:25:01 +0200</pubDate>
    </item>
    <item>
    <title>Lapinlahti/Lapinlahti, rakennuspalo: keskisuuri</title>
    <description>07.08.2019 12:24:14 Lapinlahti/Lapinlahti rakennuspalo: keskisuuri</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 12:24:14 +0200</pubDate>
    </item>
    <item>
    <title>Helsinki/Helsingfors, palohälytys</title>
    <description>07.08.2019 12:19:52 Helsinki/Helsingfors palohälytys</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 12:19:52 +0200</pubDate>
    </item>
    <item>
    <title>Turku/Åbo, rakennuspalo: pieni</title>
    <description>07.08.2019 12:19:38 Turku/Åbo rakennuspalo: pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 12:19:38 +0200</pubDate>
    </item>
    <item>
    <title>Kittilä/Kittilä, palohälytys</title>
    <description>07.08.2019 12:18:02 Kittilä/Kittilä palohälytys</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 12:18:02 +0200</pubDate>
    </item>
    <item>
    <title>Akaa/Akaa, tieliikenneonnettomuus: pieni</title>
    <description>07.08.2019 12:04:46 Akaa/Akaa tieliikenneonnettomuus: pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 12:04:46 +0200</pubDate>
    </item>
    <item>
    <title>Tampere/Tammerfors, palohälytys</title>
    <description>07.08.2019 12:00:37 Tampere/Tammerfors palohälytys</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 12:00:37 +0200</pubDate>
    </item>
    <item>
    <title>Espoo/Esbo, eläimen pelastaminen</title>
    <description>07.08.2019 11:56:03 Espoo/Esbo eläimen pelastaminen</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 11:56:03 +0200</pubDate>
    </item>
    <item>
    <title>Helsinki/Helsingfors, rakennuspalo: pieni</title>
    <description>07.08.2019 11:55:53 Helsinki/Helsingfors rakennuspalo: pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 11:55:53 +0200</pubDate>
    </item>
    <item>
    <title>Raisio/Reso, eläimen pelastaminen</title>
    <description>07.08.2019 11:51:32 Raisio/Reso eläimen pelastaminen</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 11:51:32 +0200</pubDate>
    </item>
    <item>
    <title>Vantaa/Vanda, palohälytys</title>
    <description>07.08.2019 11:50:10 Vantaa/Vanda palohälytys</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 11:50:10 +0200</pubDate>
    </item>
    <item>
    <title>Siikalatva/Siikalatva, palohälytys</title>
    <description>07.08.2019 11:42:45 Siikalatva/Siikalatva palohälytys</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 11:42:45 +0200</pubDate>
    </item>
    <item>
    <title>Närpiö/Närpes, tieliikenneonnettomuus: pieni</title>
    <description>07.08.2019 11:40:28 Närpiö/Närpes tieliikenneonnettomuus: pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 11:40:28 +0200</pubDate>
    </item>
    <item>
    <title>Raasepori/Raseborg, tulipalo muu: pieni</title>
    <description>07.08.2019 11:40:21 Raasepori/Raseborg tulipalo muu: pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 11:40:21 +0200</pubDate>
    </item>
    <item>
    <title>Pori/Björneborg, tieliikenneonnettomuus: pieni</title>
    <description>07.08.2019 11:35:23 Pori/Björneborg tieliikenneonnettomuus: pieni</description>
    <link>http://www.peto-media.fi</link>
    <pubDate>Wed, 07 Aug 2019 11:35:23 +0200</pubDate>
    </item>
    </channel>
    </rss>`

	events := Parse(feed)
	if len(events) != 100 {
		t.Errorf("Feed should contain %d items, contained %d", 100, len(events))
	}
	event := events[0]
	if event.Location != "Liperi/Liperi" {
		t.Errorf("Feed item location %s to be %s", event.Location, "Liperi/Liperi")
	}

}
