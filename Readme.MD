ini adalah api penghubung merchant dan bank. Dimana cara menjalankan dengan cara yakni :
    "go run ."
APi Memiliki beberap fitur diantarannya yakni :
1.	register : customer dapat register dengan mengakses localhost:8080/api/v1/register dipostman dan menginput body raw json berupa username dan password seperti digambar :
    ![Logo](./assets/image/register.png)

2.	login : customer dapat melakukan login dengan mengakses localhost:8080/api/v1/login dipostman dan menginput body raw json berupa username dan password seperti digambar :
    ![Logo](./assets/image/login.png)

3.	register merchant: merchant dapat melakukan register dengan mengakses localhost:8080/api/v1/registerMerchant dipostman dengan menginputkan body raw json berupa name dan password seperti digambar:
    ![Logo](./assets/image/registerMerchant.png)

4.	transaction: customer dapat melakukan transaksi dengan mengakses localhost:8080/api/v1/transaction dipostman dengan menginputkan body raw json berupa from sebagai pengirim/customer, to sebagai penerima/merchant, dan amount atau nilai yang akan dikirim seperti digambar :
    ![Logo](./assets/image/transaction.png)

5.	history transaction: fitur untuk mendapatkan history transaksi yang dilakukan, dengan localhost:8080/api/v1/history seperti digambar :
    ![Logo](./assets/image/history.png)

6. logout : customer dapat melakukan logout akun dengan mengakses loclahost:8080/api/v1/logout di postman seperti di gambar :
    ![Logo](./assets/image/logout.png)


    