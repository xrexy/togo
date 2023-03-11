<template>
    <ViewWrapper :name="username ? `Sup, ${username}` : 'Never seen you before, you new here?'">
        <input v-model="username" />
    </ViewWrapper>
</template>

<script setup lang="ts">
import { $ref } from 'vue/macros'
import ViewWrapper from './ViewWrapper.vue'

let username = $ref<string>()
let password = $ref<string>()
</script>
