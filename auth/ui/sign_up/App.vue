<template>
  <div style="width: 450px; padding: 30px 50px 30px 50px" class="var-elevation--3">
    <var-style-provider :style-vars="{ '--tab-font-size': '20px' }">
      <var-tabs v-model:active="activeTab" style="margin-bottom: 20px">
        <var-tab name="phone">{{ t("tabPhone") }}</var-tab>
        <var-tab name="email">{{ t("tabEmail") }}</var-tab>
      </var-tabs>
    </var-style-provider>
    <var-tabs-items v-model:active="activeTab" style="margin-bottom: 30px">
      <var-tab-item name="phone">
        <Phone />
      </var-tab-item>
      <var-tab-item name="email"></var-tab-item>
    </var-tabs-items>
  </div>
</template>

<script setup>
import { ref, reactive } from '@vue/reactivity'
// i18n
import { useI18n } from 'vue-i18n'
const { t } = useI18n()
// Active tab name
const activeTab = ref("phone")
// Phone form
import Phone from './PhoneForm.vue'
// // validator
// import isMobilePhone from "validator/es/lib/isMobilePhone";
// import isEmail from "validator/es/lib/isEmail";
// import isNumeric from "validator/es/lib/isNumeric";
// // disabled all components
// const disabled = ref(false);
// const showPassword = ref(false);
// // Rules.
// const rulePhoneNumber = [(v) => isMobilePhone(v) || t("errPhoneNumber")];
// const ruleEmail = [(v) => isEmail(v) || t("errEmail")];
// const ruleVerificationCode = [
//   (v) => (isNumeric(v) && v.length === 6) || t("errVerificationCode"),
// ];
// // Phone form
// const phoneForm = ref(null);
// const phoneFormData = reactive({
//   number: "",
//   password: "",
//   code: "",
// });
// // Send SMS button
// const smsButtonCounter = ref(0);
// const smsButtonDisabled = computed({
//   get: () => smsButtonCounter.value > 0,
//   set: (value) => {
//     if (!value) {
//       smsButtonCounter.value = 0;
//     }
//   },
// });
// // Email form
// const emailForm = ref(null);
// const emailFormData = reactive({
//   email: "",
//   password: "",
//   code: "",
// });
// // Handle sendSMS
// const onSendSMS = () => {
//   smsButtonCounter.value = 10000;
// };
// // Handle sign up
// const onSignUp = () => {
//   disabled.value = true;
//   phoneForm.value
//     .validate()
//     .then((v) => {
//       disabled.value = true;
//       console.log(v);
//     })
//     .catch((e) => {
//       disabled.value = true;
//       console.log(v);
//     });
// };
// // Handle sign in
// const onSignIn = () => {
//   window.location.pathname = "/sign_in";
// };
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