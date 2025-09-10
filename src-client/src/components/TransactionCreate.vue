<template>
  <main class="p-6 bg-gray-100 min-h-screen">
    <div class="max-w-4xl mx-auto bg-white p-6 rounded-xl shadow-lg">
      <div class="flex justify-between items-center mb-6 border-b-2 border-indigo-500 pb-2">
        <h4 class="text-2xl font-bold text-gray-800">Tambah Transaksi Baru</h4>
        <RouterLink
          to="/"
          class="px-4 py-2 bg-gray-500 text-white rounded-lg hover:bg-gray-600 transition-colors"
        >
          Kembali
        </RouterLink>
      </div>

      <form @submit.prevent="submitForm" class="space-y-6">
        <div>
          <label for="type" class="block text-sm font-medium text-gray-700">Tipe</label>
          <select
            v-model="form.type"
            id="type"
            class="mt-1 block w-full p-2 border border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500"
            required
          >
            <option value="" disabled>Pilih Tipe</option>
            <option value="buy">Beli</option>
            <option value="sell">Jual</option>
          </select>
        </div>

        <div>
          <label for="ticker" class="block text-sm font-medium text-gray-700">Ticker</label>
          <input
            v-model="form.ticker"
            id="ticker"
            type="text"
            class="mt-1 block w-full p-2 border border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500"
            placeholder="Contoh: AAPL"
            required
          />
        </div>

        <div>
          <label for="volume" class="block text-sm font-medium text-gray-700">Volume</label>
          <input
            v-model.number="form.volume"
            id="volume"
            type="number"
            min="0"
            step="0.1"
            class="mt-1 block w-full p-2 border border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500"
            placeholder="Contoh: 10.5"
            required
          />
        </div>

        <div>
          <label for="price" class="block text-sm font-medium text-gray-700">Harga</label>
          <input
            v-model.number="form.price"
            id="price"
            type="number"
            min="0"
            step="0.01"
            class="mt-1 block w-full p-2 border border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500"
            placeholder="Contoh: 150.75"
            required
          />
        </div>

        <div>
          <label for="date" class="block text-sm font-medium text-gray-700">Tanggal</label>
          <input
            v-model="form.date"
            id="date"
            type="datetime-local"
            class="mt-1 block w-full p-2 border border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500"
            required
          />
        </div>

        <div>
          <button
            type="submit"
            class="w-full px-4 py-2 bg-indigo-500 text-white rounded-lg hover:bg-indigo-600 transition-colors disabled:opacity-50"
            :disabled="loading"
          >
            <span v-if="loading">Memproses...</span>
            <span v-else>Tambah Transaksi</span>
          </button>
        </div>

        <div v-if="error" class="p-3 bg-red-100 border border-red-400 text-red-700 rounded-md">
          {{ error }}
        </div>
        
        <div v-if="success" class="p-3 bg-green-100 border border-green-400 text-green-700 rounded-md">
          Transaksi berhasil ditambahkan! Mengalihkan...
        </div>
      </form>
    </div>
  </main>
</template>

<script>
import { ref } from 'vue'
import { RouterLink, useRouter } from 'vue-router'
import axios from 'axios'

export default {
  name: 'TransactionCreate',
  components: {
    RouterLink
  },
  setup() {
    const form = ref({
      type: '',
      ticker: '',
      volume: null,
      price: null,
      date: ''
    })
    const loading = ref(false)
    const error = ref(null)
    const success = ref(false)
    const router = useRouter()

    const submitForm = async () => {
      loading.value = true
      error.value = null
      success.value = false
      
      try {
        // Validasi form
        if (!form.value.type || !form.value.ticker || !form.value.volume || !form.value.price || !form.value.date) {
          throw new Error('Semua field harus diisi')
        }

        if (form.value.volume <= 0 || form.value.price <= 0) {
          throw new Error('Volume dan harga harus lebih dari 0')
        }

        // Format data untuk API
        const payload = {
          type: form.value.type,
          ticker: form.value.ticker.toUpperCase(), // Convert to uppercase
          volume: parseFloat(form.value.volume),
          price: parseFloat(form.value.price),
          date: new Date(form.value.date).toISOString() // Convert to ISO format
        }

        console.log('Mengirim data:', payload)

        // Gunakan endpoint yang benar melalui Traefik
        const response = await axios.post('/api/txn', payload, {
          headers: { 
            'Content-Type': 'application/json'
          }
        })

        console.log('Response dari API:', response.data)
        
        success.value = true
        
        // Redirect setelah 1.5 detik
        setTimeout(() => {
          router.push('/')
        }, 1500)
        
      } catch (err) {
        if (err.response?.data) {
          error.value = err.response.data.error || 'Gagal menambah transaksi'
        } else if (err.message) {
          error.value = err.message
        } else {
          error.value = 'Gagal menambah transaksi. Cek koneksi API dan console untuk detail.'
        }
        console.error('Error detail:', err)
      } finally {
        loading.value = false
      }
    }

    return {
      form,
      loading,
      error,
      success,
      submitForm
    }
  }
}
</script>