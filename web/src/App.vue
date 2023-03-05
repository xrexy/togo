<template>
    <Header />
    <div class="w-full h-[calc(100%-80px)] flex gap-4 flex-col justify-center items-center" v-if="isInitializing">
        <img src="gifs/almond_spin.gif" class="h-32" />
        <h4>{{ messages[currentMessageIndex] }}</h4>
    </div>
    <RouterView v-else />
</template>
<script setup lang="ts">
import Header from '@components/Header.vue';
import { onBeforeMount } from 'vue';
import { $ref } from 'vue/macros';
import { useAuth } from './composables';

let isInitializing = $ref(false)
let currentMessageIndex = $ref(0)

const messages: string[] = [
    'Loading',
    'Maybe today...',
    'Maybe tomorrow...',
    'Maybe next week...',
    'Maybe next month...',
    "Okay... what's going on?",
    "I'm getting tired of waiting...",
    "I'm getting really tired of waiting...",
    'Are you guys serious?',
    "I'm going to bed...",
    "Good morni- oh, it's still loading...",
    'Wow... this is depressing...',
    'I give up'
]

onBeforeMount(async () => {
    isInitializing = true
    const { init } = useAuth() // we don't need to store the auth object, we just need to call it and it will initialize itself
    if (init.value) {
        init.value.then(() => {
            isInitializing = false
        })
    }

    // Generate random interval between 500 and 1500 ms and use that one
    let i = interval()
    const id = setInterval(() => {
      // If we're not initializing anymore, stop the interval
        if(!isInitializing) clearInterval(id)

        // If we're still initializing, increment message index and generate a new interval
        i = interval()
        if (currentMessageIndex < messages.length - 1) {
            currentMessageIndex++
        }
    }, i)
})

function interval(max = 2000, min = 650) {
    return Math.floor(Math.random() * (max - min + 1) + min)
}
</script>
