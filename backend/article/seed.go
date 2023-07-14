package article

import "time"

func SeedArticle() []Article {
	return []Article{
		{
			Title:       "Cara membuat bot telegram di rust",
			Content:     "Di Telego Anda dapat mendaftarkan saluran untuk obrolan tertentu dan jenis pembaruan. Jenis pembaruan adalah file yang berisi pembaruan. Ini dapat berisi message, edited_message, inline_query, dll. Anda dapat mendaftarkan saluran yang akan diperbarui saat pembaruan berisi kolom yang ditentukan. Anda juga dapat menentukan obrolan, artinya saluran akan diperbarui ketika dari obrolan itu. Kedua parameter ini dapat digabungkan bersama untuk membuat saluran yang akan dikhususkan untuk jenis pembaruan tertentu dalam obrolan tertentu.Dalam kode di bawah ini kami mendaftarkan saluran hanya untuk menerima pembaruan pesan. Karena kami tidak menentukan id obrolan, semua pembaruan dari semua obrolan yang berisi kolom pesan akan diteruskan ke saluran ini",
			Category:    "programming",
			Status:      "publish",
			CreatedDate: time.Now(),
		},
		{
			Title:       "Cara membuat bot telegram di golang",
			Content:     "Di Telego Anda dapat mendaftarkan saluran untuk obrolan tertentu dan jenis pembaruan. Jenis pembaruan adalah file yang berisi pembaruan. Ini dapat berisi message, edited_message, inline_query, dll. Anda dapat mendaftarkan saluran yang akan diperbarui saat pembaruan berisi kolom yang ditentukan. Anda juga dapat menentukan obrolan, artinya saluran akan diperbarui ketika dari obrolan itu. Kedua parameter ini dapat digabungkan bersama untuk membuat saluran yang akan dikhususkan untuk jenis pembaruan tertentu dalam obrolan tertentu.Dalam kode di bawah ini kami mendaftarkan saluran hanya untuk menerima pembaruan pesan. Karena kami tidak menentukan id obrolan, semua pembaruan dari semua obrolan yang berisi kolom pesan akan diteruskan ke saluran ini",
			Category:    "programming",
			Status:      "publish",
			CreatedDate: time.Now(),
		},
		{
			Title:       "Cara membuat bot telegram di javascript",
			Content:     "Di Telego Anda dapat mendaftarkan saluran untuk obrolan tertentu dan jenis pembaruan. Jenis pembaruan adalah file yang berisi pembaruan. Ini dapat berisi message, edited_message, inline_query, dll. Anda dapat mendaftarkan saluran yang akan diperbarui saat pembaruan berisi kolom yang ditentukan. Anda juga dapat menentukan obrolan, artinya saluran akan diperbarui ketika dari obrolan itu. Kedua parameter ini dapat digabungkan bersama untuk membuat saluran yang akan dikhususkan untuk jenis pembaruan tertentu dalam obrolan tertentu.Dalam kode di bawah ini kami mendaftarkan saluran hanya untuk menerima pembaruan pesan. Karena kami tidak menentukan id obrolan, semua pembaruan dari semua obrolan yang berisi kolom pesan akan diteruskan ke saluran ini",
			Category:    "programming",
			Status:      "publish",
			CreatedDate: time.Now(),
		},
	}
}
