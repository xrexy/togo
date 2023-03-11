<template>
    <div
        :class="`pt-4 px-4 sm:px-8 lg:px-32 w-full min-h-[calc(100%-80px)] overflow-x-hidden ${outterAdditionalClasses || ''}`"
    >
        <h1
            :class="`${nameAdditionalClasses || ''} text-orange-600 font-semibold tracking-widest text-2xl w-full overflow-hidden text-ellipsis`"
        >
            {{ name }}
        </h1>
        <slot />
    </div>
</template>

<script setup lang="ts">
defineProps({
    name: {
        type: String,
        required: true
    },
    outterAdditionalClasses: {
        type: String,
        required: false
    },
    nameAdditionalClasses: {
        type: String,
        required: false
    }
})
</script>
