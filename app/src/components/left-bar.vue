<template>
<div class="bg-gray-400 w-60 flex flex-col justify-between h-full">
    <div class="p-4">
        <button class="p-4 rounded-full bg-gradient text-white w-full text-center">New Chat</button>
    </div>
    <Profile />
</div>
</template>

<script setup lang="ts">
import Profile from './profile.vue';

async function createChat() {
    
}

</script>