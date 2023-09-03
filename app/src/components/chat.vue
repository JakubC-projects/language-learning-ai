
<template>
  <div class="w-[1024px] mx-auto relative">
    <div v-for="message in conversation" :key="message.content" :style="message.role === 'user' ? 'background-color: gray;': ''" style="padding:16px">
      <p>{{ message.role }}:</p>
      <p>{{ message.content }}</p>
    </div>
    <div class="fixed bottom-0 w-[1024px]">
      <Prompt/>
    </div>
  </div>
</template>
<script setup lang="ts">
import { ref } from 'vue';
import { Message } from '../types/message';
import Prompt from './prompt.vue';


const conversation = ref<Message[]>([])

</script>