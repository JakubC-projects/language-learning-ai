<template>
    <div class="flex items-end gap-2 w-full m-4 bg-white border-2 pl-6 p-1 border-black rounded-3xl">
          <textarea ref="promptTextarea" class="grow outline-none self-center max-h-64 resize-none h-6" placeholder="Type your message here" v-model="prompt" @input="resize"></textarea>
          <button @click="submit" class="bg-gradient rounded-full py-2 px-4 text-white font-bold">Submit</button>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';

const prompt = ref("")
const promptTextarea = ref<HTMLTextAreaElement | null>(null)

const emit = defineEmits<{
  (event: 'submit', prompt: string): void
}>()

function submit() {
    emit('submit', prompt.value)
    prompt.value = ""
}

function resize() {
    if(!promptTextarea.value) return;

    promptTextarea.value.style.height = "1.5rem"
    promptTextarea.value.style.height = promptTextarea.value.scrollHeight + "px";
}
</script>