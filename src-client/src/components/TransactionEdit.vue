<template>
  <main class="p-6 bg-gray-100 min-h-screen">
    <div class="max-w-4xl mx-auto bg-white p-6 rounded-xl shadow-lg">
      <div class="flex justify-between items-center mb-6 border-b-2 border-indigo-500 pb-2">
        <h4 class="text-2xl font-bold text-gray-800">Edit Transaksi</h4>
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
            type="date"
            class="mt-1 block w-full p-2 border border-gray-300 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500"
            required
          />
        </div>

        <div>
          <button
            type="submit"
            class="w-full px-4 py-2 bg-indigo-500 text-white rounded-lg hover:bg-indigo-600 transition-colors"
            :disabled="loading"
          >
            <span v-if="loading">Memproses...</span>
            <span v-else>Simpan Perubahan</span>
          </button>
        </div>

        <div v-if="error" class="text-red-500 text-center">{{ error }}</div>
      </form>
    </div>
  </main>
</template>

<script>
import { ref, onMounted } from 'vue'
import { RouterLink, useRouter, useRoute } from 'vue-router'
import axios from 'axios'

export default {
  name: 'TransactionEdit',
  components: {
    RouterLink
  },
  setup() {
    const form = ref({
      id: null,
      type: '',
      ticker: '',
      volume: null,
      price: null,
      date: ''
    })
    const loading = ref(false)
    const error = ref(null)
    const router = useRouter()
    const route = useRoute()

    onMounted(async () => {
      const id = route.params.id
      try {
        const response = await axios.get(`http://localhost:8080/api/txn/${id}`)
        const transaction = response.data
        form.value = {
          id: transaction.id,
          type: transaction.type,
          ticker: transaction.ticker,
          volume: transaction.volume,
          price: transaction.price,
          date: new Date(transaction.date).toISOString().split('T')[0]
        }
      } catch (err) {
        error.value = 'Gagal memuat transaksi.'
        console.error(err)
      }
    })

    const submitForm = async () => {
      loading.value = true
      error.value = null
      try {
        await axios.put(`http://localhost:8080/api/txn/${form.value.id}`, {
          type: form.value.type,
          ticker: form.value.ticker,
          volume: form.value.volume,
          price: form.value.price,
          date: form.value.date
        }, {
          headers: { 'Content-Type': 'application/json' }
        })
        router.push('/')
      } catch (err) {
        error.value = 'Gagal menyimpan perubahan.'
        console.error(err)
      } finally {
        loading.value = false
      }
    }

    return {
      form,
      loading,
      error,
      submitForm
    }
  }
}
</script>