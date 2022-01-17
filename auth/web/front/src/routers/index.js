import { createRouter, createWebHistory } from "vue-router";
import Test from "../pages/test/App.vue";
import SignIn from "../pages/SignIn.vue";
import SignUp from "../pages/SignUp.vue";
import ForgotPassword from "../pages/ForgotPassword.vue";

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: "/test",
      component: Test,
    },
    {
      path: "/",
      redirect: "/sign_in",
    },
    {
      path: "/sign_in",
      component: SignIn,
    },
    {
      path: "/sign_up",
      component: SignUp,
    },
    {
      path: "/forgot_password",
      component: ForgotPassword,
    },
  ],
});

export default router;
