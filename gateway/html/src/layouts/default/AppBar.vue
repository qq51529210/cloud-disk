<template>
  <!-- 抽屉 -->
  <v-navigation-drawer v-model="drawer">
  </v-navigation-drawer>
  <!-- 应用栏 -->
  <v-app-bar>
    <!-- 应用栏 - 抽屉按钮 -->
    <v-app-bar-nav-icon @click="drawer = !drawer"></v-app-bar-nav-icon>
    <!-- 应用栏 - 标题 -->
    <v-toolbar-title>Application</v-toolbar-title>
  </v-app-bar>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
// 
const drawer = ref(false)
</script>
