<template>
    <div v-if="user" class="flex items-center justify-between gap-2 p-2 ">
        <div class="flex items-center gap-2">
            <img :src="user.picture" class="w-10 h-10 rounded-full"/>
            <h4 class="font-bold">
                {{ user.name }}
            </h4>
        </div>
        <div class="relative" id="profile-actions" v-on-click-outside="() => areActionsOpen = false" >
            <!-- three dots -->
            <Dots class="cursor-pointer" @click="areActionsOpen = true" />
            <!-- menu -->
            <div v-if="areActionsOpen" class="absolute right-0 bg-white rounded-md">
                <a @click="logout" class="hover:bg-gray-300">Logout</a>
            </div>
        
        </div>
    </div>
</template>

<script setup lang="ts">

import { vOnClickOutside } from '@vueuse/components'
import Dots from './icons/dots.vue'
import { ref } from "vue";
import {getUser, logout} from "../auth"
import {User} from "@auth0/auth0-spa-js"

const user = ref<User>()
const areActionsOpen = ref(false);

const init = async () => {
    user.value = await getUser()
}

init();

</script>