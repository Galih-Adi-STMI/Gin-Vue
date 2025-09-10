<template>
  <main class="p-6 bg-gray-100 min-h-screen">
    <div class="max-w-4xl mx-auto bg-white p-6 rounded-xl shadow-lg">
      <!-- Header and Add Button -->
      <div class="flex justify-between items-center mb-6 border-b-2 border-indigo-500 pb-2">
        <h4 class="text-2xl font-bold text-gray-800">Transaksi</h4>
        <RouterLink
          to="/transactions/create"
          class="px-4 py-2 bg-indigo-500 text-white rounded-lg hover:bg-indigo-600 transition-colors"
        >
          Tambah Transaksi
        </RouterLink>
      </div>

      <!-- Table -->
      <div class="overflow-x-auto">
        <table class="min-w-full bg-white border border-gray-200 rounded-lg">
          <thead>
            <tr class="bg-gradient-to-r from-indigo-500 to-purple-600 text-white uppercase text-sm leading-normal">
              <th class="py-3 px-4 border-b border-gray-300 text-left font-medium">ID</th>
              <th class="py-3 px-4 border-b border-gray-300 text-left font-medium">Type</th>
              <th class="py-3 px-4 border-b border-gray-300 text-left font-medium">Ticker</th>
              <th class="py-3 px-4 border-b border-gray-300 text-left font-medium">Volume</th>
              <th class="py-3 px-4 border-b border-gray-300 text-left font-medium">Price</th>
              <th class="py-3 px-4 border-b border-gray-300 text-left font-medium">Date</th>
              <th class="py-3 px-4 border-b border-gray-300 text-left font-medium">Action</th>
            </tr>
          </thead>
          <tbody>
            <tr v-if="loading" class="text-center">
              <td colspan="7" class="py-4 text-gray-500">Memuat data...</td>
            </tr>
            <tr v-for="transaction in transactions" :key="transaction.id" class="hover:bg-gray-50 transition-colors duration-200 border-b border-gray-100" v-else>
              <td class="py-3 px-4 text-gray-700">{{ transaction.id }}</td>
              <td class="py-3 px-4 text-gray-700 capitalize">{{ transaction.type }}</td>
              <td class="py-3 px-4 text-gray-700">{{ transaction.ticker }}</td>
              <td class="py-3 px-4 text-gray-700">{{ transaction.volume }}</td>
              <td class="py-3 px-4 text-gray-700">${{ transaction.price.toFixed(2) }}</td>
              <td class="py-3 px-4 text-gray-700">{{ formatDate(transaction.date) }}</td>
              <td class="py-3 px-4">
                <RouterLink :to="`/transactions/${transaction.id}/edit`" class="text-blue-500 hover:underline mr-2">Edit</RouterLink>
                <button @click="deleteTransaction(transaction.id)" class="text-red-500 hover:underline">Delete</button>
              </td>
            </tr>
            <tr v-if="!loading && transactions.length === 0" class="text-center">
              <td colspan="7" class="py-4 text-gray-500">Tidak ada transaksi atau API bermasalah</td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Error Message -->
      <div v-if="error" class="mt-4 text-red-500 text-center">{{ error }}</div>
    </div>
  </main>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import { RouterLink } from 'vue-router'
import axios from 'axios'

interface Transaction {
  id: number
  type: string
  ticker: string
  volume: number
  price: number
  date: string
}

export default defineComponent({
  name: 'Dashboard',
  components: {
    RouterLink
  },
  setup() {
    const transactions = ref<Transaction[]>([])
    const error = ref<string | null>(null)
    const loading = ref<boolean>(true)

    const getTransactions = async () => {
      loading.value = true
      error.value = null
      try {
        const response = await axios.get('http://localhost/api/txn') // Gunakan Traefik routing
        transactions.value = response.data.Transaksi || response.data.transactions || []
        console.log('API Response:', response.data)
        console.log('Transactions:', transactions.value)
      } catch (err) {
        error.value = 'Gagal terhubung ke API. Periksa log backend atau jaringan.'
        console.error('Error fetching transactions:', err)
      } finally {
        loading.value = false
      }
    }

    const deleteTransaction = async (id: number) => {
      if (confirm('Yakin ingin menghapus transaksi ini?')) {
        try {
          await axios.delete(`http://localhost/api/txn/${id}`)
          transactions.value = transactions.value.filter(t => t.id !== id)
        } catch (err) {
          error.value = 'Gagal menghapus transaksi.'
          console.error('Error deleting transaction:', err)
        }
      }
    }

    const formatDate = (dateString: string) => {
      return new Date(dateString).toLocaleDateString('id-ID', {
        year: 'numeric',
        month: 'long',
        day: 'numeric',
        timeZone: 'Asia/Jakarta'
      })
    }

    // Fetch data on mount
    getTransactions()

    return {
      transactions,
      error,
      loading,
      deleteTransaction,
      formatDate
    }
  }
})
</script>

<style scoped>
table {
  border-collapse: separate;
  border-spacing: 0;
}

th:first-child {
  border-top-left-radius: 0.5rem;
}

th:last-child {
  border-top-right-radius: 0.5rem;
}

td:last-child {
  border-bottom-right-radius: 0.5rem;
}

tr:last-child td {
  border-bottom: none;
}
</style>