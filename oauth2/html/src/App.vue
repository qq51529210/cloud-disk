<template>
  <div>
    <v-snackbar v-model="snackbarShow"
                vertical
                :timeout="2000"
                :color="snackbarColor"
                location="top"
                variant="tonal">
      <div class="text-subtitle-1 pb-2">{{ snackbarTitle }}</div>
      <p>{{ snackbarMessage }}</p>
      <template v-slot:actions>
        <v-btn variant="text"
               :color="snackbarColor"
               @click="snackbarShow = false">关闭</v-btn>
      </template>
    </v-snackbar>
  </div>
  <router-view />
</template>

<script lang="ts" setup>
import { ref, provide } from 'vue'
//
const snackbarShow = ref(false)
const snackbarTitle = ref('')
const snackbarMessage = ref('')
const snackbarColor = 'deep-orange-darken-1'
const showTip = (title: string, message: string) => {
  snackbarTitle.value = title
  snackbarMessage.value = message
  snackbarShow.value = true
}
//
provide('showTip', showTip)
</script>
