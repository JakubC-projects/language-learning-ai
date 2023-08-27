
<template>
    <div>
      <div v-for="message in conversation" :key="message.content" :style="message.role === 'user' ? 'background-color: gray;': ''" style="padding:16px">
        <p>{{ message.role }}:</p>
        <p>{{ message.content }}</p>
      </div>
      <input v-model="prompt">
      <button @click="submit">Submit</button>
    </div>
  </template>
  <script setup lang="ts">
  import { ref } from 'vue';
  import {sendPrompt} from '../api/message'
  import { Message } from '../types/message';
  
  const prompt = ref("")
  
  const conversation = ref<Message[]>([])
  
  async function submit() {
    conversation.value.push({role: 'user', content: prompt.value})
    const res = await sendPrompt(conversation.value);
    conversation.value.push(res.value);
    prompt.value = ""
  }
  </script>