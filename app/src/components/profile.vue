<template>
    <div v-if="user" class="flex items-center gap-2 border border-black p-2 ">
        <img :src="user.picture" class="w-10 h-10 rounded-full"/>
        <h4 class="font-bold">
            {{ user.name }}
        </h4>
    </div>
</template>

<script setup lang="ts">

import { ref } from "vue";
import {getUser} from "../auth"
import {User} from "@auth0/auth0-spa-js"

const user = ref<User>()

const init = async () => {
    user.value = await getUser()
}

init();

</script>