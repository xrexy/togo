<template>
    <ViewWrapper
        :name="
            username
                ? `Sup, ${username.slice(0, MAX_USERNAME_LENGTH)}`
                : 'Never seen you before, you new here?'
        "
        name-additional-classes="ml-4 pl-4 border-l-2 border-orange-600"
    >
        <div class="input-group">
            <label class="input-label" for="username">Yup, my name is...</label>
            <input
                placeholder="what was my name again?"
                v-model="username"
                :maxlength="MAX_USERNAME_LENGTH"
            />
        </div>
        <div class="input-group">
            <label class="input-label" for="username">U can find me on</label>
            <input placeholder="myprofessionalemail@anime.xyz" type="email" v-model="email" />
        </div>
        <div class="input-group">
            <label class="input-label" for="username">password pls</label>
            <input placeholder="please dont share kthx" type="password" v-model="password" />
        </div>
        <div class="pt-8 ml-4 pl-4 pb-2 w-full text-left border-b-2 border-l-2 border-orange-600">
            <button @click="handleLoginClick()" :disabled="errors.length > 0" class="disabled:text-gray-500 disabled:hover:cursor-not-allowed">
                Can I go now?
                <div ref="errorProgressBar" class="w-0 border-b-2 border-red-400"></div>
            </button>
        </div>

        <div class="pt-2 pl-6 text-red-400">
            <p v-for="(error, k) in errors" :key="k">{{ error }}</p>
        </div>
    </ViewWrapper>
</template>

<script setup lang="ts">
import { $ref } from 'vue/macros'
import ViewWrapper from './ViewWrapper.vue'

const MAX_USERNAME_LENGTH = 50
const ERRORS_TTL = 3 * 1000

let username = $ref<string>()
let email = $ref<string>()
let password = $ref<string>()

let errors = $ref<string[]>([])
let errorCleanupTimeoutId = $ref<number>()
let errorTTLStart = $ref<number>()
let errorProgressBar = $ref<HTMLElement>()

function cleanupErrors() {
    if (errorCleanupTimeoutId) clearTimeout(errorCleanupTimeoutId)
    errors = []
}

function handleLoginClick() {
    errors = []
    if (!username) {
        errors.push('Username is required')
    }
    if (!email) {
        errors.push('Email is required')
    }
    if (!password) {
        errors.push('Password is required')
    }

    if (errors.length) {
        errorTTLStart = Date.now()
        errorCleanupTimeoutId = setTimeout(() => {
            cleanupErrors()
        }, ERRORS_TTL)
    }

    if (errors.length && errorTTLStart && errorProgressBar) {
        requestAnimationFrame(updateErrorProgressBar)
    }
}

function updateErrorProgressBar() {
    if (!errorTTLStart || !errorProgressBar) return
    
    const errorTTL = Date.now() - errorTTLStart
    const errorTTLPercentage = errorTTL / ERRORS_TTL
    errorProgressBar.style.width = `${errorTTLPercentage * 100}%`
    if (errorTTLPercentage < 1) {
        requestAnimationFrame(updateErrorProgressBar)
    } else {
        errorProgressBar!.style.width = '0%'
    }
}
</script>

<style scoped>
.input-group {
    @apply pb-4 sm:pb-0 ml-4 pl-4 pt-1 border-l-2 border-orange-600 w-full flex flex-col lg:flex-row gap-2;
}

.input-group:last-child {
    @apply pb-1;
}

.input-group input {
    @apply w-[50ch] bg-transparent outline-none border-b border-gray-400/25 text-gray-400/75;
}

.input-group input::placeholder {
    @apply text-gray-100/25;
}

.input-label {
    @apply text-gray-400 w-48;
}
</style>
