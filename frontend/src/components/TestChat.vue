<template>
  <div class="chat-app bg-gray-100 p-4 h-screen">
    <div class="messages-container overflow-y-auto h-5/6 bg-white p-4 mb-4">
      <div
        v-for="message in messages"
        :key="message.id"
        class="message p-2 my-2 bg-blue-100 rounded"
      >
        {{ message.text }}
      </div>
    </div>
    <div class="message-input flex">
      <input
        type="text"
        v-model="newMessage"
        class="flex-grow p-2 border rounded-l"
        placeholder="Type a message..."
      />
      <button @click="sendMessage" class="bg-blue-500 text-white px-4 rounded-r">Send</button>
    </div>
  </div>
</template>

<script>
import { ref } from 'vue'

export default {
  setup() {
    const newMessage = ref('')
    const messages = ref([])

    const sendMessage = () => {
      if (newMessage.value.trim()) {
        messages.value.push({ text: newMessage.value, id: Date.now() })
        newMessage.value = ''
      }
    }

    return { newMessage, messages, sendMessage }
  }
}
</script>

<style scoped>
.messages-container {
  max-height: 80%;
}
</style>
