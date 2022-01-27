<script setup>
import LinkVue from '../components/Link.vue';
</script>

<template>
    <LinkVue
        title="oauth2"
        to="http://localhost:3000/oauth2?app_id=app-id&redirect_url=http://localhost:3000/test&scope=user&state=random-string"
    />
</template>
