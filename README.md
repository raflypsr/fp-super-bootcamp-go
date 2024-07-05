# fp-super-bootcamp-go

1. wajib menggunakan cors (asumsikan semua frontend bisa akses)
2. database boleh local jika tidak deploy ke server, tapi jika deploy ke server wajib database dari server juga
3. wajib menggunakan GORM
4. wajib ada tabel user, profile, komentar
5. terdapat register, login, ganti password dan wajib ada JWT Auth Middleware (bukan basic auth) pada method yang dibutuhkan
6. minimal terdiri dari 6 tabel (harus ada keterkaitan antara tabel satu sama lain)
7. wajib ada dokumentasi API (minimal postman collection yang di export ke JSON)
8. diperbolehkan jika ingin dideploy ke server (optional tidak wajib)


* gambaran program
- diawal register
- user yang teregistrasi lalu login bisa crud profile
- user yang teregistrasi lalu login sebagai admin bisa crud books 
- user yang teregistrasi lalu login sebagai admin bisa crud categories
- user yang teregistrasi lalu login bisa crud review
- user yang teregistrasi lalu login bisa crd likes
- user yang teregistrasi lalu login bisa crud comments

* Plan database 
- semuanya memiliki created_at dan updated_at
- users
    - id, email, nama, password
- profile refrence ke user
    - id, umur, namalengkap, jeniskelamin, user_id
- books
    - id, judul, penulis, penerbit, tahunterbit,
- categories
    - id, nama
- reviews refrence ke books dan users
    - id, user_id, book_id
- likes refrence ke reviews dan user
    - id, user_id, review_id
- comments refrence ke reviews
    - id, deskripsi, , user_id, review_id

* Plan web service
- buat register ke database
- buat login dan menghasilan token dan set ke header authorization
- buat middleware authentication untuk mengecek authorization tiap   mengakses fitur
- buat logout untuk set ke header authorization = kosong dan navigasi ke root
- kalo udh login user dan admin bisa ganti password, gmail, username
- kalo udh login user dan admin bisa liat profil (default kosong semua)
- kalo udh login user dan admin bisa buat profil
- kalo udh login user dan admin bisa update profil
- hanya user dengan role admin yang bisa crud review
- user dengan role user hanya bisa read
- user biasa dan admin bisa comment dan like review
- hanya admin yang bisa crud buku
- user hanya bisa read buku
- hanya admin yang bisa crud categories
- user bisa read categories,  read buku berdasarkan categories

* fitur tambahan kalo keburu
- user bisa cari review berdasarkan like/comment terbanyak/terdikit

* deployment tidak wajib tapi dapet nilai tambah
- deploy bisa pake railway ataupun koyeb
- deploy database bisa di supabase

- dokumetasi gunakan swaggo





