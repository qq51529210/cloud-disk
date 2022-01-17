<script setup>
import FlexBoxVue from './FlexBox.vue';
const props = defineProps(['jutify', 'align'])
</script>
<template>
    <FlexBoxVue style="margin-top: 20px;" :jutify="props.jutify" align="center">
        <slot />
    </FlexBoxVue>
</template>
