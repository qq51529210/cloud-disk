import { createI18n } from "vue-i18n";

const i18n = createI18n({
  locale: navigator.language,
  // locale: 'en',
  fallbackLocale: "en",
  messages: {
    "zh-CN": {
      htmlTitle: "注册",
      tabPhone: "手机注册",
      tabEmail: "邮箱注册",
      placeholderEmail: "邮箱地址",
      placeholderNumber: "手机号码",
      placeholderPassword: "密码",
      placeholderVerificationCode: "验证码",
      sendSMS: "发送短信",
      signUp: "注册",
      signIn: "登录",
      showPassword: "显示密码",
      errPhoneNumber: "错误的手机号码格式",
      errEmail: "错误的邮箱格式",
      errVerificationCode: "错误的验证码格式"
    },
    en: {
      htmlTitle: "Sign Up",
      tabPhone: "Cell-Phone Sign Up",
      tabEmail: "Email Sign Up",
      placeholderEmail: "Email",
      placeholderNumber: "Number",
      placeholderPassword: "Password",
      placeholderVerificationCode: "Verification Code",
      sendSMS: "Send SMS",
      signUp: "Sign Up",
      signIn: "Sign In",
      showPassword: "Show Password",
      errPhoneNumber: "Invalid phone number format",
      errEmail: "Invalid email format",
      errVerificationCode: "Invalid verification code format"
    }
  }
});

document.title = i18n.global.t("htmlTitle");

export default i18n;
