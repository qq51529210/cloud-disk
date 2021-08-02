<template>
  <div style="width:350px; height: 350px;">
    <Tabs :tabBarStyle="{ textAlign: 'center' }" size="large">
      <TabPane key="1" :tab="$t('tabPhone')"></TabPane>
      <TabPane key="2" :tab="$t('tabEmail')"></TabPane>
    </Tabs>
  </div>
</template>

<script setup>
// ui
import Tabs from "ant-design-vue/lib/tabs";
import "ant-design-vue/lib/tabs/style/css";
// i18n
import { useI18n } from "vue-i18n";
const { t } = useI18n();
document.title = t("htmlTitle");
</script>

<style>
html,
body {
  margin: 0;
}
#app {
  width: 100%;
  height: 100%;
  position: absolute;
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>