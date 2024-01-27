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

DEPLOYMENT TO RAILway.APP

1. Login akun raywails, jika belum maka registrasi.
2. Klik Deploy from repo seperti digamabar ini :
   ![Logo](./assets/image/deployr1.png)
3. Klik Configuration GIthub App.
4. Selanjutnya pilih repository yang ingin dideploy seperti digambar :
   ![Logo](./assets/image/deplor3.png)
5. lalu klik initialize.
6. Setelah itu akan kembali ke menu awal trus klik deploy from repo lagi kemudian klik repositori seperti digambar ini :
    ![Logo](./assets/image/deployr5.png)
7. Klik Deploy Now seperti digambar :
  ![Logo](./assets/image/deplorr6.png)
8. Selanjutnya menambahka variabel PORT dengan nilai 8080 seperti digambar :
   ![Logo](./assets/image/menabhport.png)
9. Tunggu proses hingga selesai seprti digambar :
    ![Logo](./assets/image/donerr.png)

TEST DI POSTMAN 
1. Register :
   ![Logo](./assets/image/deploytest.png)
2. Login :
   ![Logo](./assets/image/deploytestlogin.png)
3. Register Merchant :
   ![Logo](./assets/image/deployregistertest.png)
4. Transaction:
   ![Logo](./assets/image/deploytesttransaction.png)
5. History:
   ![Logo](./assets/image/historydeploytest.png)
6. Logout :
   ![Logo](./assets/image/deploytestlogout.png)


    