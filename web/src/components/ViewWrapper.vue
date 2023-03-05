<template>
  <div class="pt-4 px-32 min-h-screen">
    <h1 class="text-orange-600 font-semibold tracking-widest text-2xl">{{ name }}</h1>
    <slot />
  </div>
</template>

<script setup lang="ts">
defineProps({
    name: {
        type: String,
        required: true,
    }
});
</script>