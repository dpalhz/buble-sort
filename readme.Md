# (5) Mengapa Bubble Sort Memiliki Kompleksitas O(n²)?

## Intuisi Umum

Bubble sort melakukan **perbandingan dan pertukaran (swap)** berulang pada **pasangan elemen bersebelahan**. Proses ini diulang dalam beberapa **pass** sampai data terurut. Karena ada **dua loop bersarang** (pass luar dan perbandingan dalam), jumlah operasi tumbuh **kuadratik** terhadap ukuran input `n`.

---

## Perhitungan Jumlah Perbandingan

### Versi Dasar (tanpa pengurangan perbandingan per pass)

- **Loop luar**: sampai `n-1` pass.
- **Loop dalam**: pada tiap pass melakukan ~`n-1` perbandingan.
- Perkiraan total perbandingan:
  \[
  (n-1)\times(n-1) = n^2 - 2n + 1 = O(n^2)
  \]
- Karena operasi dominan adalah perbandingan dan swap pada pasangan elemen, kompleksitas waktu **worst-case** dan **average-case** adalah **O(n²)**.

### Versi dengan Pengurangan Perbandingan per Pass (optimized window)

- Setelah pass ke-1, **1 elemen** sudah berada pada posisi final; setelah pass ke-2, **2 elemen** sudah final; dan seterusnya.
- Jumlah perbandingan menjadi deret:
  \[
  (n-1) + (n-2) + \dots + 1 = \frac{n(n-1)}{2} = O(n^2)
  \]
- Meskipun lebih hemat dibanding versi dasar, **orde** kompleksitas tetap **O(n²)**.

---

## Early Exit (Berhenti Saat Tidak Ada Swap)

- Jika pada satu pass **tidak terjadi swap**, array sudah terurut → algoritma **berhenti**.
- **Best-case**: data sudah terurut → hanya **1 pass**, **O(n)** (hanya satu kali menyapu untuk memastikan tidak ada swap).
- **Average/worst-case**: data acak atau terbalik → tetap mendekati jumlah perbandingan kuadratik, sehingga **O(n²)**.

---

## Ringkasan Kompleksitas

- **Worst-case**: O(n²) — data terbalik, banyak swap dan pass.
- **Average-case**: O(n²) — urutan acak rata-rata membutuhkan ~n²/2 perbandingan.
- **Best-case** (dengan early exit): **O(n)** — satu pass tanpa swap.
- **Space**: **O(1)** — in-place sorting (hanya beberapa variabel bantu).

---

## Mengapa Tetap O(n²)?

Walaupun kita mengurangi perbandingan tiap pass, totalnya membentuk **deret aritmetika** yang skalanya **kuadratik** terhadap `n`. Inilah alasan mendasar kenapa bubble sort, secara sifat algoritmik, memiliki kompleksitas waktu O(n²) pada rata-rata dan kasus terburuk.
