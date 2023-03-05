<template>
    <header class="w-full h-20 border border-orange-400/50 flex items-center justify-between py-4">
        <img class="h-full aspect-square" src="gifs/almond_pet.gif" />
        <div class="flex gap-6 pr-4 items-center">
            <ul class="flex text-gray-500 gap-4">
                <li v-for="route in routes">
                    <router-link
                        :to="route.path"
                        :class="{ 'text-orange-400': routeIsActive(route.path) }"
                    >
                        {{ route.name }}
                    </router-link>
                </li>
            </ul>
            <button
                class="py-1 px-4 bg-orange-400 text-white rounded-md hover:bg-orange-500 hover:rounded-xl transition duration-500 ease-out"
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
