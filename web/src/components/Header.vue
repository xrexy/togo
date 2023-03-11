<template>
    <header class="w-full h-20 border-b border-orange-400/25 flex items-center justify-between py-4">
        <img class="h-full aspect-square" src="gifs/almond_pet.gif" />
        <div class="flex gap-6 pr-4 items-center">
            <ul class="flex gap-4">
                <li v-for="route in routes">
                    <router-link
                        class="hover:text-gray-400"
                        :to="route.path"
                        :class="{ 'text-orange-600': routeIsActive(route.path) }"
                    >
                        {{ route.name }}
                    </router-link>
                </li>
            </ul>
            <button
                class="py-1 px-4 bg-orange-600 text-white/80 rounded-md hover:bg-orange-700 hover:text-white transition duration-500 ease-out"
                @click="$router.push({
                    name: 'login'
                })"
            >
                Login
            </button>
        </div>
    </header>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useRoute } from 'vue-router'

type Route = {
    path: string
    name: string
}

const routes: Route[] = [
    {
        path: '/',
        name: 'Home'
    },
    {
        path: '/dashboard',
        name: 'Dashboard'
    }
]

const $route = useRoute()

const routeIsActive = computed(() => (route: string) => {
    return $route.path === route
})
</script>
